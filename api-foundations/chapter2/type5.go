package main

import "github.com/davecgh/go-spew/spew"

type Dog struct {
	name  string
	breed string
}
type Cat struct {
	name           string
	hypoallergenic bool
}
type Petstore struct {
	name string
	pets []interface{}
}

func main() {
	p := Petstore{}
	p.pets = append(p.pets, &Dog{name: "Winston"})
	p.pets = append(p.pets, &Cat{name: "Whiskers"})
	spew.Dump(p)
}
