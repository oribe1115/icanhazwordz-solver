package solver

import (
	"github.com/oribe1115/icanhazwordz-solver/lib"
	"github.com/siddontang/go/log"
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

	maxIndex := 0
	maxScore := 0
	for i, word := range list {
		score := word.CalcScore(countQ, countU)
		if score > maxScore || (score == maxScore && len(word.Sorted) > len(list[maxIndex].Sorted)) {
			maxScore = score
			maxIndex = i
		}
	}

	log.Errorf("maxIndex: %d", maxIndex)
	if list[maxIndex] == nil {
		log.Errorf("nil index at taraget: %s", target)
	}
	if len(list[maxIndex].Examples) == 0 {
		log.Errorf("no example at %d", maxIndex)
	}

	return list[maxIndex].Examples[0], maxScore
}
