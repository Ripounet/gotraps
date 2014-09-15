package main

import "fmt"

func main() {
	var x Animal = Dog{"Rosie"}

	if x, ok := x.(Human); ok {
		fmt.Println("%v doesn't want to be treated like dogs and cats.", x.lastName)
	} else {
		fmt.Println(x.Say())
	}
}

type Animal interface {
	Say() string
}

type Dog struct {
	name string
}

func (d Dog) Say() string {
	return fmt.Sprintf("%v barks", d.name)
}

type Cat struct {
	name string
}

func (c Cat) Say() string {
	return fmt.Sprintf("%v meows", c.name)
}

type Human struct {
	firstName string
	lastName  string
}

// Humans are technically animals, and they say things.
func (h Human) Say() string {
	return fmt.Sprintf("%v %v speaks", h.firstName, h.lastName)
}
