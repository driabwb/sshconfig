package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func Add(c *cli.Context) error {
	//if hostExists() {
	//		return errHostExists
	//	}
	fmt.Println(c.Args())
	return nil
}
