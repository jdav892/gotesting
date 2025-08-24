package main

import "testing"

func TestHello(t *testing.T) {
//example of subtests
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "", "")
		want := "Hello, Citizen of Earth"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := Hello("Justin", "English", "USA")
		want := "Hello, Justin from America"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish", "Spain")
		want := "Hola, Elodie from Europe"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", "French", "France")
		want := "Bonjour, Pierre from Europe"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Akira", "Japanese", "Japan")
		want := "Konichiwa, Akira from Asia"
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
