package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

func InitCli(args []string) {
	app := cli.App{
		Usage: "goredis",
		Version: "0.0.0.1",
		Action: func(ctx *cli.Context) error{
			fmt.Println(ctx.Args())
			return nil
		},
	}
	app.Run(args)
}
