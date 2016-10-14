package gohello

import (
	"fmt"
)

type Person struct {
	id   int64
	name string
}

func NewPerson(id int64, name string) Person {
	return Person{id: id, name: name}
}

func (p *Person) Greet() string {
	return fmt.Sprintf("Hi, I am %v", p.name)
}
