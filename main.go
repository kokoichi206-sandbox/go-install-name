package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	version = "0.0.1"
)

func main() {

	app := cli.App{
		Name:    "go-install-name",
		Usage:   "go-install-name cli",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:        "hello",
				Aliases:     []string{"h"},
				Description: "Hello to something",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
				},
				Action: func(c *cli.Context) error {
					fmt.Printf("Hello %s\n", c.String("name"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
