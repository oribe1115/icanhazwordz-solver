package main

import (
	"fmt"

	"github.com/oribe1115/icanhazwordz-solver/lib"
)

func main() {
	dictionary, err := lib.GetDictionary()
	if err != nil {
		panic(err)
	}

	fmt.Println(dictionary[10])
}
