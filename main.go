package main

import (
	"os"

	"github.com/Danex2/saber/router"
	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Out: os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks: make(logrus.LevelHooks),
	Level: logrus.ErrorLevel,
}



func main() {
	
	err := godotenv.Load()

	if err != nil {
		log.Fatal("There was an error loading the .env file")
	}

	client := disgord.New(disgord.Config{
		ProjectName: "saber",
		BotToken: os.Getenv("DISCORD_TOKEN"),
		RejectEvents: []string{
			disgord.EvtTypingStart,
			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
					Name: "This is just a test lol",
			},
		},
		Logger: log,

	})
	
	router.Bot.Hook(client)

	defer client.Gateway().StayConnectedUntilInterrupted()
	

}

func init() {
	router.Bot.SetCommand(&gommand.Command{
		Name: "ping",
		Description: "pong",
		Function: func(ctx *gommand.Context) error {
			_, _ = ctx.Reply("Pong")
			return nil
		},
	})
}
