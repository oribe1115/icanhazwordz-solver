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
	fmt.Println("2: Find words constructed with input")
	fmt.Println("3: solve with hand")
	fmt.Println("4: auto")
	fmt.Printf("> ")

	num := lib.ReadLine()

	switch num {
	case "1":
		solver.FindEqualWords(dictionary)
		return
	case "2":
		solver.FindContainAnagramWords(dictionary)
		return
	case "3":
		solver.SolverWithHand(dictionary)
		return
	case "4":
		solver.AutoSolver()
		return
	default:
		fmt.Println("No such mode")
	}

}
