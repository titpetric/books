package main

type Dog struct {
	name string
}
type Cat struct {
	name           string
	hypoallergenic bool
}

func main() {
	cat := Cat{name: "Whiskers"}
	dog := Dog(cat)
}
