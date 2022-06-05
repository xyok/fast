package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"{{ .AppName }}/lib/log"
	"{{ .AppName }}/conf"
	"{{ .AppName }}/server"
	"{{ .AppName }}/setup"
	"github.com/urfave/cli"
)

var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Start  server",
	Flags:   startFlags,
	Action:  startAction,
}

var startFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config,c",
		Value: "app.ini",
		Usage: "Custom configuration file path",
	},
	cli.StringFlag{
		Name:  "addr",
		Value: ":9000",
		Usage: "http port",
	},
	cli.StringFlag{
		Name:  "mode",
		Value: "prod",
		Usage: "run server mod,[dev|all]",
	},
	cli.StringFlag{
		Name:  "env",
		Value: "pub",
		Usage: "env [dev|pub]",
	},
	cli.StringFlag{
		Name:  "loglevel,l",
		Value: "",
		Usage: "log level [debug|info|warn|error]",
	},
}

func startAction(ctx *cli.Context) error {
	conf.MustInit(ctx.String("config"))

	if err := setup.All(); err != nil {
		return err
	}

	cctx, cancel := context.WithCancel(context.Background())

	// service.Init()
	go server.Start(cctx, ctx.String("addr"))

	reloadConfSign := make(chan os.Signal, 1)
	signal.Notify(reloadConfSign, syscall.SIGUSR1) //30
	go func() {
		for {
			<-reloadConfSign
			//30 mac,10 linux
			conf.MustInit(ctx.String("config"))
			log.Info("----- ReLoad conf -----")
		}
	}()

	// set up proper shutdown of daemon and web server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutting down...")

	cancel()
	setup.Down()

	time.Sleep(2 * time.Second)

	return nil
}