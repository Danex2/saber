package commands

import (
	"strings"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
)

func init() {
	Bot.SetCommand(&gommand.Command{
		Name: "user",
		Description: "Get info about the user",
		Function: func(ctx *gommand.Context) error {

			result := strings.SplitN(ctx.Message.Member.JoinedAt.String(), "T", 2)

			t, _ := time.Parse("2006-01-02", result[0])

			userImg, _ := ctx.Message.Author.AvatarURL(2048, false)


			_, _ = ctx.Reply(&disgord.Embed{
				Title: "User Details",
				Fields: []*disgord.EmbedField{{
					Name: "ID",
					Value: ctx.Message.Author.ID.String(),
				},{
					Name: "Username",
					Value: ctx.Message.Member.User.Tag(),
				}, {
					Name: "Server Join Date",
					Value: t.Format("January 2 2006"),
				}},
				Thumbnail: &disgord.EmbedThumbnail{
					URL: userImg,
				},
			})

			return nil
		},
	})
}