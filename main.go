package main

import (
	"log"
	"os"

	"github.com/hyperupcall/knowledge/aggregators"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "aggregator",
		Usage: "Run an aggregator",
		Action: func(c *cli.Context) error {
			aggregators.AnsiAggregator()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
