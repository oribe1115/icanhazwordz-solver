package solver

import (
	"fmt"

	"github.com/oribe1115/icanhazwordz-solver/lib"
)

func SolverWithHand(dictionary lib.WordList) {
	fmt.Println("input word")
	fmt.Print("> ")

	sortedTarget := lib.StringSort(lib.ReadLine())
	isContainU := lib.IsContainU(sortedTarget)
	isContainQ := lib.IsContainQ(sortedTarget)
	if isContainQ {
		sortedTarget = lib.StringSort(sortedTarget + "u")
	}

	result := lib.WordList{}
	for _, word := range dictionary {
		if word.IsEnableConstruct(sortedTarget) {
			result = append(result, word)
		}
	}

	// if isContainQ {
	// 	for _, word := range result {
	// 		fmt.Printf("%s: %d -> %d\n", word.Examples[0], word.CalcScore(false), word.CalcScore(lib.IsQUCase(word.Examples[0])))
	// 	}
	// } else {
	// 	for _, word := range result {
	// 		fmt.Printf("%s: %d\n", word.Examples[0], word.CalcScore(false))
	// 	}
	// }
	for _, word := range result {
		fmt.Printf("%s: %d -> %d\n", word.Examples[0], word.CalcScore(false), word.CalcScore(lib.IsDeductCase(word.Examples[0], isContainU, isContainQ)))
	}
}
