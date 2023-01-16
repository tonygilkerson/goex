package main

import (
	"fmt"
	"local/pkg/person"
)

//
// Person
//
type NormalPerson interface {
	Walk() 
	Talk() 
}

//
// Person Builder
//

// Holds just the fields needed to build a valid person
type personBuilder struct {
	name string
	age  int
}

type PersonBuilder interface {
	WithName(string) PersonBuilder
	WithAge(int) PersonBuilder
	Build() NormalPerson
}

func (p *personBuilder) WithAge(age int) PersonBuilder {
	p.age = age
	return p
}

func (p *personBuilder) WithName(name string) PersonBuilder {
	p.name = name
	return p
}

func (p *personBuilder) Build() NormalPerson {

	person := person.NewPerson() // implementation is decoupled, could swap in something different here
	person.Age = p.age
	person.Name = p.name

	return person

}

func NewPersonBuilder() PersonBuilder {
	return &personBuilder{}
}

func main() {

	p := NewPersonBuilder().
	WithAge(42).
	WithName("tony").
	Build()

	
	fmt.Printf("\nNormal person:\n")
	p.Walk()
	p.Talk()

}
