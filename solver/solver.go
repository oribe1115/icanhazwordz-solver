package solver

import (
	"github.com/oribe1115/icanhazwordz-solver/lib"
)

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

	if len(list) == 0 {
		return "", 0
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
