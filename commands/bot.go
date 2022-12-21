package commands

import (
	"bot/utils/command"

	"github.com/disgoorg/disgo/discord"
)

func init() {
	command.RegisterCommand(command.Command{
		Name:        "bot",
		Description: "Bot related commands",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionSubCommand{
				Name:        "info",
				Description: "Get info about the bot",
			},
		},

		Runner: infohandle,
	})
}

func infohandle(ctx *command.CommandContext) {

	embed := discord.NewEmbedBuilder().
		SetAuthorNamef("Olá, eu sou a %s", "Madee").Build()
	ctx.Event.CreateMessage(
		discord.NewMessageCreateBuilder().SetEmbed(0, embed).
			Build(),
	)
}
