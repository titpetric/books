package main

import "fmt"
import "encoding/json"

type Petstore struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Dogs     []*Pet `json:"dogs"`
	Cats     []*Pet `json:"cats"`
}

type Pet struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

type PetStoreList []*Petstore

func main() {
	petstorelist := PetStoreList{}
	petstore := &Petstore{
		Name:     "Fuzzy's",
		Location: "New York, 5th and Broadway",
	}
	petstore.Dogs = append(petstore.Dogs,
		&Pet{
			Name:  "Whiskers",
			Breed: "Pomeranian",
		},
	)
	petstore.Dogs = append(petstore.Dogs,
		&Pet{Name: "Trinity"},
	)
	petstorelist = append(petstorelist, petstore)

	jsonString, _ := json.MarshalIndent(petstorelist, "", "\t")
	fmt.Printf("%s", jsonString)
}
