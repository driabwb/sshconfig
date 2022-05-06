package main

import (
	"context"
	"io"
)

func addHost(ctx context.Context, writer io.Writer, hostName string) {
	if hostName == "TestHost" {
		writer.Write([]byte("Host TestHost"))
	} else {
		writer.Write([]byte("Host Test2"))
	}
}
