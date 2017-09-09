package main

import "github.com/urfave/cli"

var (
	login = cli.StringFlag{
		Name:  "l",
		Usage: "Specifies the user to log in as on the remote machine.  This also may be specified on a per-host basis in the configuration file.",
	}

	id_file = cli.StringFlag{
		Name:  "i",
		Usage: "",
	}

	port = cli.IntFlag{
		Name:  "p",
		Usage: "",
		Value: 22,
	}
)
