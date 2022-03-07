package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Neo")

	got := buffer.String()
	want := "Hello, Neo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
