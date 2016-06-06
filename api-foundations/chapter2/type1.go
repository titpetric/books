package main

type Dog struct {
	name string
}
type Cat struct {
	name           string
	hypoallergenic bool
}

func main() {
	dog := Dog{name: "Rex"}
	cat := Cat(dog)
}
