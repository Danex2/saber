package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
)

func init() {
	Bot.SetCommand(&gommand.Command{
		Name:        "ban",
		Description: "Ban a user from the discord server",
		ArgTransformers: []gommand.ArgTransformer{
			{
				Function: gommand.MemberTransformer,
			},
			{
				Function: gommand.IntTransformer,
			},
			{
				Function: gommand.StringTransformer,
				Greedy:   true,
			},
		},
		Function: func(ctx *gommand.Context) error {

			user := ctx.Args[0].(*disgord.Member)
			days := ctx.Args[1].(int)
			reasonArray := ctx.Args[2].([]interface{})
			var reason []string

			for _, v := range reasonArray {
				reason = append(reason, v.(string))
			}

			embed := &disgord.Embed{
				Title:  fmt.Sprintf("User %s has been banned!", user.User.Tag()),
				Fields: []*disgord.EmbedField{},
			}

			embed.Fields = append(embed.Fields, &disgord.EmbedField{
				Name:  "Duration",
				Value: fmt.Sprintf("%d days", days),
			})

			embed.Fields = append(embed.Fields, &disgord.EmbedField{
				Name:  "Reason",
				Value: strings.Join(reason[:], " "),
			})

			embed.Color = 15724753

			embed.Timestamp = disgord.Time{
				Time: time.Now().UTC(),
			}

			ctx.Session.Guild(ctx.Message.GuildID).Member(user.User.ID).Ban(&disgord.BanMemberParams{
				DeleteMessageDays: days,
				Reason:            strings.Join(reason[:], " "),
			})

			_, _ = ctx.Reply(embed)

			return nil
		},
	})
}
