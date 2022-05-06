package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	hostname := "TODO"
	fileName := "todo"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("Failed to open config file: %s", fileName)
		os.Exit(1)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error attempting to close config file %s", err.Error())
			os.Exit(1)
		}
	}()

	addHost(context.Background(), file, hostname)
}
