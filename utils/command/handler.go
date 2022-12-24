package command

import (
  "github.com/disgoorg/disgo/bot"
  "github.com/disgoorg/disgo/discord"
  "github.com/disgoorg/disgo/events"
  "github.com/richaardev/Ocean/utils/telemetry"
  "github.com/richaardev/Ocean/utils/translation"
)

var Commands []Command

var InteractionEvent = make(chan *events.ApplicationCommandInteractionCreate)
var Workers = 128
var WorkersArray = make([]bool, Workers)

type Command struct {
	Name                     string
	Description              string
	NameLocalizations        map[discord.Locale]string
	DescriptionLocalizations map[discord.Locale]string
	Options                  []discord.ApplicationCommandOption

	DefaultMemberPermissions discord.Permissions
	AllowDM                  bool

	Runner func(ctx *Context) error
}

func init() {
	for i := 0; i < Workers; i++ {
		go Worker(i)
	}
}

func Worker(id int) {
	for interaction := range InteractionEvent {
		WorkersArray[id] = true
		telemetry.Debugf("Worker %d is handling command %s executed by %s (%s)", id, interaction.SlashCommandInteractionData().CommandName(), interaction.Member().User.Username+"#"+interaction.Member().User.Discriminator, interaction.Member().User.ID)
		HandleCommand(interaction)
		WorkersArray[id] = false
	}
}

func Init(client bot.Client) {
	RemoveUnregisteredCommands(client)
	SendCommandsToDiscord(client)
}

func HandleCommand(e *events.ApplicationCommandInteractionCreate) {
	slash := e.SlashCommandInteractionData()
	command := GetCommand(slash.CommandName())

	if command == nil {
		return
	}

	err := command.Runner(&Context{
		Event: e,
    T: translation.GetFixedT(translation.DefaultLanguage),
	})

	if err != nil {

	}
}

func RegisterCommand(command Command) {
	Commands = append(Commands, command)
}

func GetCommand(name string) *Command {
	for _, command := range Commands {
		if command.Name == name {
			return &command
		}
	}
	return nil
}

func SendCommandsToDiscord(client bot.Client) {
	var commands []discord.ApplicationCommandCreate
	for _, command := range Commands {
		cmd := discord.SlashCommandCreate{
			Name:                     command.Name,
			Description:              command.Description,
			NameLocalizations:        command.NameLocalizations,
			DescriptionLocalizations: command.DescriptionLocalizations,
			Options:                  command.Options,
		}

		commands = append(commands, cmd)
	}

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), commands); err != nil {
		telemetry.Errorf("Error setting commands: %s", err.Error())
	}

	telemetry.Infof("Slash commands registered and updated with %d commands", len(commands))
}

func RemoveUnregisteredCommands(client bot.Client) {
	commands, err := client.Rest().GetGlobalCommands(client.ApplicationID(), true)
	if err != nil {
		return
	}

	for _, command := range commands {
		if GetCommand(command.Name()) == nil {
			telemetry.Debugf("Removing command %s", command.Name())
			err := client.Rest().DeleteGlobalCommand(client.ApplicationID(), command.ID())
			if err != nil {
				telemetry.Errorf("An error occurred while deleting command %s", command.Name())
			}
		}
	}
	telemetry.Info("The old slash commands have been removed")
}
