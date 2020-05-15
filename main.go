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

	sortedTarget := lib.StringSort("Brock")
	word := dictionary.FindEqualWord(sortedTarget, 0, len(dictionary)-1)
	if word == nil {
		fmt.Println("no answer")
	} else {
		fmt.Println("found")
		fmt.Println(word)
	}
}
