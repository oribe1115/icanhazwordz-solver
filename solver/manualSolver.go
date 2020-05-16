package solver

import (
	"fmt"

	"github.com/oribe1115/icanhazwordz-solver/lib"
)

func SolverWithHand(dictionary lib.WordList) {
	for i := 0; i < 10; i++ {
		fmt.Printf("turn: %d\n > ", i+1)
		target := lib.ReadLine()
		answer, score := solver(dictionary, target)
		fmt.Printf("%s: %d\n", answer, score)
	}

	fmt.Println("finish")
}
