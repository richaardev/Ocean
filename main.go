package main

import (
	_ "bot/commands"
	"bot/listeners"
	commandhandler "bot/utils/command"
	"bot/utils/telemetry"
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
	client.AddEventListeners(&events.ListenerAdapter{
		OnReady: listeners.ReadyListener,
		OnApplicationCommandInteraction: listeners.InteractionListener,
	})

	go commandhandler.Init(client)
	
	if err != nil {
		log.Fatal(err)
	}

	if err := client.OpenGateway(context.TODO()); err != nil {
		log.Fatal(err)
	}


	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}