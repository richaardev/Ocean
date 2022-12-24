package commands

import (
	"github.com/richaardev/Ocean/utils/command"
	"github.com/richaardev/Ocean/utils/translation"

	"github.com/disgoorg/disgo/discord"
)

func init() {
	command.RegisterCommand(command.Command{
		Name:        "server",
		Description: "show info",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionSubCommand{
				Name:                     "info",
				Description:              translation.Translate("commands:server.info.slash.description"),
				NameLocalizations:        translation.GetLocalizationsValues("commands:server.info.slash.name"),
				DescriptionLocalizations: translation.GetLocalizationsValues("commands:server.info.slash.description"),
			},
		},

		Runner: handleserver,
	})
}

func handleserver(ctx *command.Context) error {
	switch ctx.GetSubCommandName() {
	case "info":
		return ctx.Event.CreateMessage(
			discord.NewMessageCreateBuilder().
				SetContent("OK").
				Build(),
		)
	default:
		return nil
	}
}
