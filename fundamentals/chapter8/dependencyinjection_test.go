package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Martin")

	got := buffer.String()
	want := "Hello, Martin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
