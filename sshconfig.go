package main

import (
	"context"
	"fmt"
	"io"
)

func addHost(ctx context.Context, writer io.Writer, hostName string) {
	writer.Write([]byte(fmt.Sprintf("Host %s", hostName)))
}
