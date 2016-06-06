package main

import "fmt"
import "io/ioutil"
import "encoding/json"
import "github.com/davecgh/go-spew/spew"

type Petstore struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Dogs     []*Pet `json:"dogs,omitempty"`
	Cats     []*Pet `json:"cats,omitempty"`
}

type Pet struct {
	Name  string `json:"name"`
	Breed string `json:"breed,omitempty"`
}

type PetStoreList []*Petstore

func main() {
	petstorelist := PetStoreList{}

	jsonBlob, err := ioutil.ReadFile("example2.json")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	err = json.Unmarshal(jsonBlob, &petstorelist)
	if err != nil {
		fmt.Printf("Error decoding json: %s\n", err)
	}

	spew.Dump(petstorelist)
}
