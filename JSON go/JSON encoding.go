package main

//Importing packages
import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	PetName string
	Color   string
	Mice    int
}

//Declare and implement a main package
func main() {

	//An Instance of Cat
	Cattus := Cat{"Lenerina", "brownish", 23}

	cat_enc, err := json.Marshal(Cattus)

	if err != nil {
		fmt.Println("Error marshalling Leni")
	}

	fmt.Println(string(cat_enc))

	//An Array of Cats
	Catti := []Cat{

		{"Lenerina", "brownish", 23},

		{"Buffi", "black and white", 144},

		{"Vincino", "brownish", 42},

		{"Leni", "orange and white", 5},
	}

	cat_enc2, err := json.Marshal(Catti)

	if err != nil {
		fmt.Println("Error marshalling Cats")
	}

	fmt.Println(string(cat_enc2))

}
