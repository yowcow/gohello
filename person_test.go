package gohello

import (
	"testing"
)

func TestPerson(t *testing.T) {
	p := NewPerson(1, "HOGE")

	if p.id != 1 {
		t.Errorf("Got %v where 1 was expected", p.id)
	}

	if p.name != "HOGE" {
		t.Errorf("Got %v where HOGE was expected", p.name)
	}
}

func TestPersonGreet(t *testing.T) {
	p := NewPerson(2, "FUGA")
	actual := p.Greet()
	expected := "Hi, I am FUGA"

	if actual != expected {
		t.Errorf("Got $v where %v was expected", actual, expected)
	}
}
