package main

import (
	"fmt"
	"os"

	"github.com/mgjules/deckr/cmd"
	"github.com/urfave/cli/v2"
)

// @title        Deckr
// @version      v0.1.2
// @description  A REST API for playing with a deck of cards.

// @contact.name   Michaël Giovanni Jules
// @contact.url    https://mgjules.dev
// @contact.email  julesmichaelgiovanni@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	app := cli.NewApp()
	app.Name = "deckr"
	app.Usage = "A REST API for playing with a deck of cards"
	app.Description = "Deckr exposes a REST API for playing with a deck of cards of your choice."
	app.Authors = []*cli.Author{
		{
			Name:  "Michaël Giovanni Jules",
			Email: "julesmichaelgiovanni@gmail.com",
		},
	}
	app.Copyright = "(c) 2022 Michaël Giovanni Jules"
	app.Commands = cmd.Commands

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("run app: %v\n", err)
		os.Exit(1)
	}
}
