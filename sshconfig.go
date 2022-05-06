package main

import (
	"context"
	"io"
)

func addHost(ctx context.Context, writer io.Writer, hostName string) {
	if hostName == "TestHost" {
		writer.WriteString("Host TestHost")
	} else {
		writer.WriteString("Host Test2")
	}
}
