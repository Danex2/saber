package commands

import (
	"fmt"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
)

type User struct {
	gommand.CommandBasics
}

func (u *User) Init() {
	u.Name = "user"
	u.Description = "Get basic info about the user"
	u.Usage = "@username"
	u.Category = basic
}

func (User) CommandFunction(ctx *gommand.Context) error {

	avatar, err := ctx.Message.Author.AvatarURL(2048, false)
	if err != nil {
		fmt.Println(err)
	}
	
	
	embed := disgord.Embed{
		Author: &disgord.EmbedAuthor{
			Name: ctx.Message.Author.Tag(),
			IconURL: avatar,
		},
		Thumbnail: &disgord.EmbedThumbnail{
			URL: avatar,
		},
	}

	embed.Fields = append(embed.Fields, &disgord.EmbedField{
		Name: "Join Date",
		Value: ctx.Message.Member.JoinedAt.Format(time.RFC1123Z),
	}, &disgord.EmbedField{
		Name: "Account creation date",
		Value: "test", 
	})


	_, _ = ctx.Reply(embed)
	
	return nil
}