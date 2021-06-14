package commands

import "github.com/auttaja/gommand"


type Ping struct {
	gommand.CommandBasics
}

func (p *Ping) Init() {
	p.Name = "ping"
	p.Description = "Responds with pong!"
	p.Category = basic
}

func (Ping) CommandFunction(ctx *gommand.Context) error {
	_, _ = ctx.Reply("Pong!")
	
	return nil
}