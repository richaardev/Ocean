package main

import (
	_ "github.com/richaardev/Ocean/commands"
	"github.com/richaardev/Ocean/listeners"
	"github.com/richaardev/Ocean/utils/command"
	"github.com/richaardev/Ocean/utils/telemetry"
	_ "github.com/richaardev/Ocean/utils/translation"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/joho/godotenv"
)

func main() {
	telemetry.Info("Starting bot...")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := disgo.New(os.Getenv("DISCORD_TOKEN"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentsNone,
			),
		),
	)
	if err != nil {
		panic(err)
	}

	command.Init(client)
	client.AddEventListeners(&events.ListenerAdapter{
		OnReady:                         listeners.ReadyListener,
		OnApplicationCommandInteraction: listeners.InteractionListener,
	})

	if err := client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
