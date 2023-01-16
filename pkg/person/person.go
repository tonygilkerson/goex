package person

import "fmt"

//
// Person
//
type Person struct {
	Name string
	Age  int
	Weight int
}

func (p Person) Walk() {
	fmt.Printf("Normal walk... %+v\n", p)
}
func (p Person) Talk() {
	fmt.Printf("Normal talk... %+v\n", p)
}

func NewPerson() Person {
	return Person{}
}