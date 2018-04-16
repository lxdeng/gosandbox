package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	usage         = `Inventory of products`
	formatOptions = `txt or json`
)

var version = "1.0.0"

var listCommand = cli.Command{
	Name:  "list",
	Usage: "lists products",
	ArgsUsage: `

To list with different format
Examples:
$ inventory list -f json
$ inventory list -f txt`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Value: "txt",
			Usage: `select one of: ` + formatOptions,
		},
	},
	Action: func(context *cli.Context) error {
		switch context.String("format") {
		case "txt":
			fmt.Println("table=3")
		case "json":
			fmt.Println("{table: 3}")
		default:
			fmt.Println("unknow format")
		}
		return nil
	},
}

var addCommand = cli.Command{
	Name:  "add",
	Usage: "adds products",
	ArgsUsage: `

To add a product:
Examples:
$ inventory add -n table -c 1`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "",
		},
		cli.IntFlag{
			Name:  "count, c",
			Value: 0,
		},
	},
	Action: doAdd,
}

func main() {
	app := cli.NewApp()
	app.Name = "Inventory"
	app.Usage = usage
	app.Version = version

	// global options
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log, l",
			Value: "ivt.log",
			Usage: "set the log file path",
		},
	}

	//commands
	app.Commands = []cli.Command{
		listCommand,
		addCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var doAdd = func(context *cli.Context) error {
	if context.String("name") == "" || context.Int("count") <= 0 {
		msg := "name cannot be empty, and count must be postive number"
		fmt.Println("failed")
		return errors.New(msg)
	}
	fmt.Println("Added successfully")
	return nil
}
