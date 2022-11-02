package main

import "testing"

func Test_guessInput(t *testing.T) {
	actualInput := 20
	desiredInput := GuessInput()

	if actualInput != int(desiredInput) {
		t.Errorf("want %q got %q", actualInput, int(desiredInput))
	}
}
