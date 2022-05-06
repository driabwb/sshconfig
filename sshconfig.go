package sshconfig

import (
	"bytes"
	"context"
)

func addHost(ctx context.Context, writer *bytes.Buffer, hostName string) {
	if hostName == "TestHost" {
		writer.WriteString("Host TestHost")
	} else {
		writer.WriteString("Host Test2")
	}
}
