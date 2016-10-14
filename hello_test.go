package gohello

import (
	"testing"
)

func TestHelloGetGreeting(t *testing.T) {
	expected := "Hello!"
	actual := Hello()

	if actual != expected {
		t.Errorf("Got %v where %v was expected", actual, expected)
	}
}
