package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"ocean/utils/command"
)

func init() {
	command.RegisterCommand(command.Command{
		Name:        "ocean",
		Description: "_",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionSubCommand{
				Name:        "info",
				Description: "Get info about the bot",
			},
		},

		Runner: infohandle,
	})
}

func infohandle(ctx *command.Context) error {
	ownerId, _ := snowflake.Parse("646416170123132959")
	helperId, _ := snowflake.Parse("711384230533398528")
	owner, _ := ctx.Client().Rest().GetUser(ownerId)
	helper, _ := ctx.Client().Rest().GetUser(helperId)

	self, er := ctx.GetSelf()
	if !er {
		return nil
	}

	return ctx.Event.CreateMessage(
		discord.NewMessageCreateBuilder().
			SetEmbeds(
				discord.NewEmbedBuilder().
					SetColor(0x3194c6).
					SetThumbnail(self.EffectiveAvatarURL()).
					SetDescriptionf("Olá atualmente eu estou em %d servidores com %d comandos."+
						"\nFui criada em <t:%d> por **`%s`** junto com a **`%s`**, programada em "+
						"**[GoLang](https://go.dev/)** usando **[disgo](https://github.com/disgoorg/disgo)** "+
						"e atualmente eu me mantenho hospedada gratuitamente na **[SquareCloud](https://squarecloud.app/)**",
						len(ctx.Client().Caches().Guilds().All()), len(command.Commands), self.CreatedAt().Unix(),
						owner.Username+"#"+owner.Discriminator, helper.Username+"#"+helper.Discriminator,
					).
					SetAuthorNamef("Olá, eu sou a %s", self.Username).
					SetAuthorIcon(self.EffectiveAvatarURL()).
					SetFooterIcon(owner.EffectiveAvatarURL()).
					SetFooterTextf("%s foi criada e desenvolvida por %s - https://richaar.dev", self.Username, owner.Username).Build(),
			).
			Build(),
	)
}
