package main

import (
	"embed"
	"fast/gen"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app *cli.App

const (
	NAME    = "fast cli"
	USAGE   = "Create proj scaffold from current framework"
	VERSION = "0.0.1"
)

var (
	//go:embed tpl
	res embed.FS
)

func action(c *cli.Context) error {
	return gen.Execute(
		c.String("name"),
		c.String("output"),
		c.String("remote"),
		c.String("tpl"),
		res,
	)
}

var GenProjectCommand = &cli.Command{
	Name:    "gen",
	Aliases: []string{"g"},
	Usage:   "generate project",
	Action:  action,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "project name",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "tpl",
			Usage: "project tamplate tpl/[*]",
			Value: "gin",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "generate out",
		},
		&cli.StringFlag{
			Name:  "remote",
			Usage: "git remote",
		},
	},
}

func main() {
	app := cli.NewApp()
	app.Name = NAME
	app.Usage = USAGE
	app.Version = VERSION
	app.EnableBashCompletion = true

	app.Commands = []*cli.Command{
		GenProjectCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
