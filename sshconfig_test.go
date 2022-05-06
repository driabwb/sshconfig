package main

import (
	"bytes"
	"context"
	"testing"
)

func TestWriteHost(t *testing.T) {
	t.Run("Write TestHost host line", func(t *testing.T) {
		buf := bytes.Buffer{}

		addHost(context.Background(), &buf, "TestHost")

		got := buf.String()
		want := "Host TestHost"

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}
	})

	t.Run("Write Test2 host line", func(t *testing.T) {
		buf := bytes.Buffer{}
		addHost(context.Background(), &buf, "Test2")

		got := buf.String()
		want := "Host Test2"

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}
	})
}
