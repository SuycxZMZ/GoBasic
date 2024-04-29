package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "fucker")

	got := buffer.String()

	want := "Hello, fucker"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
