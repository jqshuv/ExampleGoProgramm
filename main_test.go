package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestExitCode(t *testing.T) {
	want := 0
	code := Run()
	if code != want {
		t.Errorf("got %d want %d", code, want)
	}
}
