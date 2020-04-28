package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "micro"
	app.Usage = "micro工具集"
	app.Version = Version
	app.Commands = []*cli.Command{
		{
			Name:            "new",
			Aliases:         []string{"n"},
			Usage:           "创建新项目",
			Action:          runNew,
			SkipFlagParsing: true,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "micro build",
			Action:  buildAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "micro run",
			Action:  runAction,
		},
		{
			Name:            "tool",
			Aliases:         []string{"t"},
			Usage:           "micro tool",
			Action:          toolAction,
			SkipFlagParsing: true,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "micro version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
		{
			Name:   "self-upgrade",
			Usage:  "micro self-upgrade",
			Action: upgradeAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func runNew(ctx *cli.Context) error {
	return installAndRun("genproject", ctx.Args().Slice())
}
