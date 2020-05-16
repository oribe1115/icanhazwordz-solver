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

func solver(dictionary lib.WordList, target string) (string, int) {
	sortedTarget := lib.StringSort(target)
	countQ, countU := lib.CountQU(sortedTarget)
	add := ""
	for i := 0; i < countQ; i++ {
		add += "u"
	}
	sortedTarget = lib.StringSort(sortedTarget + add)

	list := lib.WordList{}
	for _, word := range dictionary {
		if word.IsEnableConstruct(sortedTarget) {
			list = append(list, word)
		}
	}

	maxIndex := 0
	maxScore := 0
	for i, word := range list {
		score := word.CalcScore(countQ, countU)
		if score > maxScore || (score == maxScore && len(word.Sorted) > len(list[maxIndex].Sorted)) {
			maxScore = score
			maxIndex = i
		}
	}

	return list[maxIndex].Examples[0], maxScore
}
