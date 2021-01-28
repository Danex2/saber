
package commands

import "github.com/auttaja/gommand"


func init() {
	Bot.SetCommand(&gommand.Command{
		Name: "ban",
		Description: "Ban a user from the discord server",
		Function: func(ctx *gommand.Context) error {
			_, _ = ctx.Reply("ban")
			return nil
		},
	})
}
