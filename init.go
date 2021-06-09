package main

import (
	"fmt"

	"github.com/nulfrost/saber/commands"
)


func init() {
   fmt.Println("Commands loaded.")
   router.SetCommand(&commands.Hello{})
}