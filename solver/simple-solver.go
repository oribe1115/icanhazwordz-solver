package solver

import (
	"fmt"

	"github.com/oribe1115/icanhazwordz-solver/lib"
)

func SolverWithHand(dictionary lib.WordList) {
	fmt.Println("input word")
	fmt.Print("> ")

	sortedTarget := lib.StringSort(lib.ReadLine())
	countQ, countU := lib.CountQU(sortedTarget)
	add := ""
	for i := 0; i < countQ; i++ {
		add += "u"
	}
	sortedTarget = lib.StringSort(sortedTarget + add)

	result := lib.WordList{}
	for _, word := range dictionary {
		if word.IsEnableConstruct(sortedTarget) {
			result = append(result, word)
		}
	}

	for _, word := range result {
		fmt.Printf("%s: %d\n", word.Examples[0], word.CalcScore(countQ, countU))
	}
}
