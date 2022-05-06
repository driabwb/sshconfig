package main

import (
	"context"
	"errors"
	"fmt"
	"io"
)

func addHost(ctx context.Context, writer io.Writer, hostName string) error {
	writer.Write([]byte(fmt.Sprintf("Host %s", hostName)))
	return errors.New("Blah")
}
