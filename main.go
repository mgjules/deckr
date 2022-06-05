package main

import (
	"fmt"
	"os"

	"github.com/mgjules/deckr/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "deckr"
	app.Usage = "A REST/gRPC API for playing with a deck of cards"
	app.Description = "Deckr exposes a REST/gRPC API for playing with a deck of cards of your choice."
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
