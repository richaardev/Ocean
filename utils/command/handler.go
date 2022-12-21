package command

import (
	"bot/utils/telemetry"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var Commands = []Command{}

var InteractionEvent = make(chan *events.ApplicationCommandInteractionCreate)
var Workers = 128
var WorkersArray = make([]bool, Workers)

type Command struct {
	Name        string
	Description string
	Options     []discord.ApplicationCommandOption

	DefaultMemberPermissions discord.Permissions
	AllowDM                  bool

	Runner func(ctx *CommandContext)
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
	SendCommandsToDiscord(client)
	RemoveUnregisteredCommands(client)
}

func HandleCommand(e *events.ApplicationCommandInteractionCreate) {
	slash := e.SlashCommandInteractionData()
	command := GetCommand(slash.CommandName())
	
	if command == nil {
		return
	}

	command.Runner(&CommandContext{
		Event: e,
	})
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
	commands := []discord.ApplicationCommandCreate{}
	for _, command := range Commands {
		cmd := discord.SlashCommandCreate{
			Name:        command.Name,
			Description: command.Description,
			Options:     command.Options,
		}

		commands = append(commands, cmd)
	}

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), commands); err != nil {
		telemetry.Errorf("Error setting commands: %s", err.Error())
	}
}

func RemoveUnregisteredCommands(client bot.Client) {
	commands, err := client.Rest().GetGlobalCommands(client.ApplicationID(), true)
	if err != nil {
		return
	}

	for _, command := range commands {
		if GetCommand(command.Name()) == nil {
			telemetry.Debugf("Removing command %s", command.Name())
			client.Rest().DeleteGlobalCommand(client.ApplicationID(), command.ID())
		}
	}

}
