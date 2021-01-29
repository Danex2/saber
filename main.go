package main

import (
	"context"
	"os"

	"github.com/Danex2/saber/commands"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/auttaja/gommand"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("There was an error loading the .env file")
	}

	client := disgord.New(disgord.Config{
		ProjectName: "saber",
		BotToken:    os.Getenv("DISCORD_TOKEN"),
		RejectEvents: []string{
			disgord.EvtTypingStart,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
			disgord.EvtPresenceUpdate,
		},
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: "This is a new bot!",
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
			_, _ = ctx.Reply("Missing required arguments, use !help [command] for more help!")
			return true
		}

		// This was not handled here.
		return false
	})

	defer client.Gateway().StayConnectedUntilInterrupted()

	logFilter, _ := std.NewLogFilter(client)
	filter, _ := std.NewMsgFilter(context.Background(), client)

	client.Gateway().WithMiddleware(logFilter.LogMsg, filter.NotByBot, filter.HasPrefix, filter.StripPrefix)

}
