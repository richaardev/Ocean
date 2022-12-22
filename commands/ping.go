package commands

import (
    "fmt"
    handler "ocean/utils/command"
    "ocean/utils/translation"

    "github.com/disgoorg/disgo/discord"
)

func init() {
	handler.RegisterCommand(handler.Command{
		Name: "ping",
		Description: "Pings the bot",
		Runner: pinghandle,
	})
}

func pinghandle(ctx *handler.Context) (err error) {
    t := translation.GetFixedT("en-us")
    fmt.Println(t("test"))

	return ctx.Event.CreateMessage(
		discord.NewMessageCreateBuilder().SetContent("Pong!").
		Build(),
	)
}