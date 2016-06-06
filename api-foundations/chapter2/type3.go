package main

import "fmt"

type Dog struct {
	name  string
	Breed string
}

type Cat struct {
	name           string
	hypoallergenic bool
}

type Pet interface {
	getName() string
}

func (d Dog) getName() string {
	return d.name
}
func (c Cat) getName() string {
	return c.name
}

func main() {
	var pet Pet

	pet = Cat{name: "Whiskers"}
	fmt.Printf("Cat name: %s, pet %#v\n", pet.getName(), pet)

	pet = Dog{name: "Winston", Breed: "Samoyed"}
	fmt.Printf("Dog name: %s, pet %#v\n", pet.getName(), pet)
}
