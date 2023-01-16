package main

import (
	"fmt"
	"local/pkg/person"
)

type NormalPerson interface {
	Walk() 
	Talk() 
}

func main() {

	var p NormalPerson = person.NewPerson()
	
	fmt.Printf("\nNormal person:\n")
	p.Walk()
	p.Talk()

}
