package router

import "github.com/auttaja/gommand"

func main() {
	Bot.SetCommand(&gommand.Command{
		Name: "ping",
		Description: "pong",
		Function: func(ctx *gommand.Context) error {
			_, _ = ctx.Reply("Pong")
			return nil
		},
	})
}
