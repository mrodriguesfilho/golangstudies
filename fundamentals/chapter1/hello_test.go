package main

import "testing"

func TestYay(t *testing.T) {
	t.Run("saying yay to someone", func(t *testing.T) {
		got := Yay("MARCIOW", "ENG")
		want := "yay from func, MARCIOW"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying yay to a dummy", func(t *testing.T) {
		got := Yay("", "")
		want := "yay from func, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("hablando yay", func(t *testing.T) {
		got := Yay("MARCIOW", "SPA")
		want := "Hola que tal, MARCIOW"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
