package main

import (
	"context"
    "log"
    _ "ocean/commands"
	"ocean/listeners"
	commandhandler "ocean/utils/command"
	"ocean/utils/telemetry"
    "ocean/utils/translation"
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

    translation.Init()
    telemetry.Init()
	client, err := disgo.New(os.Getenv("DISCORD_TOKEN"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentsNone,
			),
		),
//        bot.WithLogger(),
	)
	if err != nil {
		panic(err)
	}

	commandhandler.Init(client)
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
