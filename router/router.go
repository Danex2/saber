package router

import "github.com/auttaja/gommand"

var Bot *gommand.Router

func init() {
	router := gommand.NewRouter(&gommand.RouterConfig{
		PrefixCheck: gommand.StaticPrefix("!"),
	})
	
	Bot = router
}
