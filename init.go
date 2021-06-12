package main

import (
	"fmt"

	"github.com/nulfrost/saber/commands"
)


func init() {
   router.SetCommand(&commands.Ping{})
   fmt.Println("Commands loaded!")
}