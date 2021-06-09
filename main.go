package main

import (
	"context"
	"log"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/auttaja/gommand"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)


var router = gommand.NewRouter(&gommand.RouterConfig{
	PrefixCheck: gommand.MultiplePrefixCheckers(gommand.StaticPrefix("!!"), gommand.MentionPrefix),
})

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logrus.SetLevel(logrus.DebugLevel)
	client := disgord.New(disgord.Config{
		ProjectName: "saber",
		BotToken: os.Getenv("DISCORD_TOKEN"),
		Logger: logrus.New(),
		RejectEvents: []string{
			disgord.EvtTypingStart,
			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},
	})


	router.Hook(client)
	defer client.Gateway().StayConnectedUntilInterrupted()

	filter, _ := std.NewMsgFilter(context.Background(), client)

	client.Gateway().WithMiddleware(
		filter.NotByBot,
		filter.StripPrefix,
	)
}
