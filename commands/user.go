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

	Bot.SetCommand(&gommand.Command{
		Name:        "user",
		Description: "Display info about the user.",
		ArgTransformers: []gommand.ArgTransformer{
			{
				Function: gommand.MemberTransformer,
			},
		},
		Usage: "@user",
		Function: func(ctx *gommand.Context) error {

			mentionedUser := ctx.Args[0]
			result := strings.SplitN(mentionedUser.(*disgord.Member).JoinedAt.String(), "T", 2)

			t, _ := time.Parse("2006-01-02", result[0])

			userImg, _ := mentionedUser.(*disgord.Member).User.AvatarURL(2048, false)

			// Move this logic into a utility function
			unixTimeStamp := ((mentionedUser.(*disgord.Member).UserID / 4194304) + 1420070400000) / 1000

			parsedTimeStamp, err := strconv.ParseInt(fmt.Sprint(unixTimeStamp), 10, 64)

			if err != nil {
				panic(err)
			}

			joinDate := time.Unix(parsedTimeStamp, 0)

			y, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", joinDate.String())

			// again more clean up probably

			var userRoles []string

			for i := range mentionedUser.(*disgord.Member).Roles {
				userRoles = append(userRoles, fmt.Sprintf("<@&%s>", mentionedUser.(*disgord.Member).Roles[i]))
			}

			embed := &disgord.Embed{
				Title:  "User Details",
				Fields: []*disgord.EmbedField{},
			}

			embed.Fields = append(embed.Fields, &disgord.EmbedField{
				Name:  "ID",
				Value: mentionedUser.(*disgord.Member).UserID.String(),
			})

			embed.Fields = append(embed.Fields, &disgord.EmbedField{
				Name:   "Username",
				Value:  mentionedUser.(*disgord.Member).User.Tag(),
				Inline: true,
			})

			embed.Fields = append(embed.Fields, &disgord.EmbedField{
				Name:  "Avatar URL",
				Value: fmt.Sprintf("[Link](%s)", userImg),
			})

			if len(mentionedUser.(*disgord.Member).Roles) > 0 {

				embed.Fields = append(embed.Fields, &disgord.EmbedField{
					Name:  "Roles",
					Value: strings.Join(userRoles, ""),
				})
			} else {

				embed.Fields = append(embed.Fields, &disgord.EmbedField{
					Name:  "Roles",
					Value: "0",
				})
			}

			embed.Fields = append(embed.Fields, &disgord.EmbedField{
				Name:   "Server Join Date",
				Value:  t.Format("January 2 2006"),
				Inline: true,
			})

			embed.Fields = append(embed.Fields, &disgord.EmbedField{
				Name:   "Account Creation Date",
				Value:  y.Format("January 2 2006"),
				Inline: true,
			})

			embed.Thumbnail = &disgord.EmbedThumbnail{
				URL: userImg,
			}

			embed.Timestamp = disgord.Time{
				Time: time.Now().UTC(),
			}

			embed.Color = 15724753

			_, _ = ctx.Reply(embed)

			return nil
		},
	})
}
