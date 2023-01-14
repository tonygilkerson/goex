package main

import (
	"local/builder/pkg/person"
)

func main() {
	p := person.New().
		WithAge(42).
		WithName("tony").Build()

	p.Talk()
	p.Walk()
}
