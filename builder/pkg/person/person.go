package person

import "fmt"

//
// Person
//
type person struct {
	name string
	age  int
	weight int
	eyeColor string
}

type Person interface {
	Walk() 
	Talk() 
}

func (p *person) Walk() {
	fmt.Println("Normal walk...")
}
func (p *person) Talk() {
	fmt.Println("Normal talk...")
}

//
// Person Builder
//
type personBuilder struct {
	name string
	age  int
}

type PersonBuilder interface {
	WithName(string) PersonBuilder
	WithAge(int) PersonBuilder
	Build() Person
}

func (p *personBuilder) WithAge(age int) PersonBuilder {
	p.age = age
	return p
}

func (p *personBuilder) WithName(name string) PersonBuilder {
	p.name = name
	return p
}

func (p *personBuilder) Build() Person {

	return &person{
		age: p.age,
		name: p.name,
	}

}

func New() PersonBuilder {
	return &personBuilder{}
}