package main

import (
	"os"

	"github.com/Danex2/saber/commands"
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
	
	commands.Bot.Hook(client)

	commands.Bot.AddErrorHandler(func(ctx *gommand.Context, err error) bool {
		switch err.(type) {
		case *gommand.CommandNotFound, *gommand.CommandBlank:
			// We will ignore. The command was not found in the router or the command was blank.
			return true
		case *gommand.InvalidTransformation:
			_, _ = ctx.Reply("Invalid argument:", err.Error())
			return true
		case *gommand.IncorrectPermissions:
			_, _ = ctx.Reply("Invalid permissions:", err.Error())
			return true
		case *gommand.InvalidArgCount:
			_, _ = ctx.Reply("Invalid argument count.")
			return true
		}

		// This was not handled here.
		return false
	})

	defer client.Gateway().StayConnectedUntilInterrupted()

}