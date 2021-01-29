package commands

import "github.com/auttaja/gommand"

func init() {
	Bot.SetCommand(&gommand.Command{
		Name:        "ping",
		Description: "Displays the average latency for the bot in milliseconds",
		Function: func(ctx *gommand.Context) error {
			_, _ = ctx.Reply(ctx.Session.AvgHeartbeatLatency())
			return nil
		},
	})
}
