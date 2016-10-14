package gohello

import (
	"testing"
)

func TestByeGetGreeting(t *testing.T) {
	expected := "Bye!"
	actual := Bye()

	if actual != expected {
		t.Errorf("Got %v where %v was expected", actual, expected)
	}
}
