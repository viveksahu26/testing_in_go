package main

import "testing"

func Test_Hello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		actual := Hello("vivek")
		desired := "Hello, vivek"

		if actual != desired {
			t.Errorf("got %q want %q", actual, desired)
		}
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("")
		desired := "Hello, World"

		if actual != desired {
			t.Errorf("got %q want %q", actual, desired)
		}
	})
}
