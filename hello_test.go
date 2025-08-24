package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Justin")
	want := "Hello, Justin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
