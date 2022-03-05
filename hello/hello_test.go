package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Neo", "")
		want := "Hello, Neo"
		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello matrix when name parameter empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Matrix"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Turkish", func(t *testing.T) {
		got := Hello("Neo", "Turkish")
		want := "Maraba, Neo"
		assertCorrectMessage(t, got, want)

	})

}
