# Saber (Yet Another Discord Bot)

Saber is a discord bot I am making using Golang to get a better understanding of the language.

## Setup

- Download and install Go from [here](https://golang.org/)

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file. You'll also need to rename the `.env.sample` file to just `.env`.

`DISCORD_TOKEN=your token`

To find out how to get a discord bot token check out this [guide](https://discordjs.guide/preparations/setting-up-a-bot-application.html#creating-your-bot).

## Run Locally

Clone the project

```bash
  git clone https://github.com/nulfrost/saber.git
```

Go to the project directory

```bash
  cd saber
```

Install dependencies

```bash
  go mod download
```

Start the bot

```bash
  go run *.go
```

## Usage / Examples

### Adding more commands

To add more commands to the bot make a new file in the `commands` folder where the name of the command is the name of the file.

```bash
cd commands
touch scream.go
```

Then inside of the file add this code.

```go
package commands

import "github.com/auttaja/gommand"


type Scream struct {
	gommand.CommandBasics
}

func (s *Scream) Init() {
	s.Name = "scream"
	s.Description = "i reply with a yell"
}

func (Scream) CommandFunction(ctx *gommand.Context) error {
	_, _ = ctx.Reply("WHAT'S UP?!")
	return nil
```

Next, inside of `init.go` register the new command.

```go
package main

import (
	"fmt"

	"github.com/nulfrost/saber/commands"
)


func init() {
   fmt.Println("Commands loaded.")
   router.SetCommand(&commands.Hello{})
   router.setCommand(&commands.Scream{})
}
```

Finally, save and restart the bot.

### Changing the prefix

In `main.go` under the import statements look for this code.

```go
var router = gommand.NewRouter(&gommand.RouterConfig{
	PrefixCheck: gommand.MultiplePrefixCheckers(gommand.StaticPrefix("!!"), gommand.MentionPrefix),
})
```

You can change `gommand.StaticPrefix()` to anything you like. For example, instead of `!!` you can change it to `??`.

```go
gommand.MultiplePrefixCheckers(gommand.StaticPrefix("??"), gommand.MentionPrefix)
```
