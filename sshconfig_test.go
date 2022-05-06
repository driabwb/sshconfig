package main

import (
	"bytes"
	"context"
	"fmt"
	"testing"
)

type badWriter struct{}

func (b badWriter) Write(_ []byte) (int, error) {
	return 0, fmt.Errorf("Bad Writer Error")
}

func TestWriteHost(t *testing.T) {
	t.Run("Write TestHost host line", func(t *testing.T) {
		buf := bytes.Buffer{}

		err := addHost(context.Background(), &buf, "TestHost")

		got := buf.String()
		want := "Host TestHost"

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}

		assetNoError(t, err)
	})

	t.Run("Write Test2 host line", func(t *testing.T) {
		buf := bytes.Buffer{}

		err := addHost(context.Background(), &buf, "Test2")

		got := buf.String()
		want := "Host Test2"

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}

		assetNoError(t, err)
	})

	t.Run("Return an error if writing fails", func(t *testing.T) {

		got := addHost(context.Background(), badWriter{}, "BadWriterTest")

		assetHasError(t, got)
	})
}

func assetHasError(t testing.TB, e error) {
	t.Helper()
	if nil == e {
		t.Errorf("Expected an error; got nil")
	}
}

func assetNoError(t testing.TB, e error) {
	t.Helper()
	if nil != e {
		t.Errorf("Expected no error; got %s", e.Error())
	}
}
