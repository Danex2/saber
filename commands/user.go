package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
)

func init() {

	// Bot.GetCommand("user").(*gommand.Command).Category = &gommand.Category{
	// 	Name: "Info",
	// 	Description: "Info commands provide information about the user or the server",
	// }

	Bot.SetCommand(&gommand.Command{
		Name: "user",
		Description: "Get info about the user",
		Function: func(ctx *gommand.Context) error {

			result := strings.SplitN(ctx.Message.Member.JoinedAt.String(), "T", 2)

			t, _ := time.Parse("2006-01-02", result[0])

			userImg, _ := ctx.Message.Author.AvatarURL(2048, false)


			// Move this logic into a utility function
			unixTimeStamp := ((ctx.Message.Author.ID / 4194304) + 1420070400000) / 1000

			parsedTimeStamp, err := strconv.ParseInt(fmt.Sprint(unixTimeStamp), 10, 64)

			if err != nil {
				panic(err)
			}

			joinDate := time.Unix(parsedTimeStamp, 0)

			y, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", joinDate.String())
		
			// again more clean up probably

			var userRoles []string

			for i := range ctx.Message.Member.Roles {
				userRoles = append(userRoles, fmt.Sprintf("<@&%s>", ctx.Message.Member.Roles[i]))
			}

			_, _ = ctx.Reply(&disgord.Embed{
				Title: "User Details",
				Fields: []*disgord.EmbedField{{
					Name: "ID",
					Value: ctx.Message.Author.ID.String(),
				},{
					Name: "Username",
					Value: ctx.Message.Member.User.Tag(),
					Inline: true,
				}, {
					Name: "Avatar URL",
					Value: fmt.Sprintf("[Link](%s)", userImg),
				}, {
					Name: "Roles",
					Value: strings.Join(userRoles, ""),
				}, {
					Name: "Server Join Date",
					Value: t.Format("January 2 2006"),
					Inline: true,
				}, {
					Name: "Account Creation Date",
					Value: y.Format("January 2 2006"),
					Inline: true,
				}},
				Thumbnail: &disgord.EmbedThumbnail{
					URL: userImg,
				},
				Timestamp: disgord.Time{
					Time: time.Now().UTC(),
				},
				Color: 15724753,
			})

			return nil
		},
	})
}