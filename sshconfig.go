package sshconfig

import (
	"bytes"
	"context"
)

func addHost(ctx context.Context, writer *bytes.Buffer, hostName string) {
	writer.WriteString("Host TestHost")
}
