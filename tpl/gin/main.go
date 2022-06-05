package main

import (
	"log"
	"os"
	"{{ .AppName }}/commands"

	"github.com/urfave/cli"
)

var (
	build   string
	version string
	commit  string
)

func main() {
	app := cli.NewApp()
	app.Name = "{{ .AppName }} svr"
	app.Usage = "{{ .AppName }} server"
	app.Version = version + " - " + build + "\ncommit:" + commit
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		commands.StartCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}