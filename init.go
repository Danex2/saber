package main

import (
	"fmt"

	"github.com/nulfrost/saber/commands"
)


func init() {
   router.SetCommand(&commands.Ping{})
   router.SetCommand(&commands.User{})
   fmt.Println("Commands loaded!")
}