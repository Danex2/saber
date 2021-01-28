package commands

import "github.com/auttaja/gommand"


func init() {
	Bot.SetCommand(&gommand.Command{
		Name: "server",
		Description: "Get info about the discord server",
		Function: func(ctx *gommand.Context) error {
			_, _ = ctx.Reply("server")
			return nil
		},
	})
}
