package commands

import "github.com/auttaja/gommand"


type Hello struct {
	gommand.CommandBasics
}

func (h *Hello) Init() {
	h.Name = "hello"
	h.Description = "says hello"
}

func (Hello) CommandFunction(ctx *gommand.Context) error {
	_, _ = ctx.Reply("hello")
	return nil
}