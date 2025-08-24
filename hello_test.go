package main

import "testing"

func TestHello(t *testing.T) {
	//example of subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Justin")
		want := "Hello, Justin"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}


//woah actually using asserts
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
