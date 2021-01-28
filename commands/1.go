package commands

import "github.com/auttaja/gommand"

var Bot *gommand.Router

func init() {
	router := gommand.NewRouter(&gommand.RouterConfig{
		PrefixCheck: gommand.MultiplePrefixCheckers(gommand.StaticPrefix("!"), gommand.MentionPrefix),
	})
	
	Bot = router
}
