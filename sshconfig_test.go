package sshconfig

import (
	"bytes"
	"context"
	"testing"
)

func TestWriteHost(t *testing.T) {
	t.Run("Write a host line", func(t *testing.T) {
		buf := bytes.Buffer{}

		addHost(context.Background(), &buf, "TestHost")

		got := buf.String()
		want := "Host TestHost"

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}
	})

}
