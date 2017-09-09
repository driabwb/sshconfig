package main

import "github.com/urfave/cli"
import "os"

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "Add a new ssh host",
			Flags: []cli.Flag{login, id_file, port},
			Action: func(c *cli.Context) error {
				return Add(c)
			},
		},
	}

	app.Run(os.Args)
}
