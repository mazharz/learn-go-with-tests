package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "me")

	got := buffer.String()
	want := "hello, me"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
