package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("ONe")
	want := "Hello, ONe!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
