package main

import (
	"fmt"

	"github.com/nulfrost/saber/commands"
)


func init() {
   router.SetCommand(&commands.Hello{})
   fmt.Println("Commands loaded!")
}