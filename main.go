package main

import (
	"fmt"

	"github.com/oribe1115/icanhazwordz-solver/lib"
	"github.com/oribe1115/icanhazwordz-solver/solver"
)

func main() {
	dictionary, err := lib.GetDictionary()
	if err != nil {
		panic(err)
	}

	lib.InitStdin()

	fmt.Println("Choose mode with number")
	fmt.Println("1: Find perfect match anagrams")
	fmt.Printf("> ")

	num := lib.ReadLine()

	switch num {
	case "1":
		solver.FindEqualWords(dictionary)
		return
	default:
		fmt.Println("No such mode")
	}

}
