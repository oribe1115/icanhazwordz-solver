package lib

import (
	"strings"

	"github.com/oribe1115/icanhazwordz-solver/config"
)

// Word ソートされた単語ごとの構造体
type Word struct {
	Sorted   string   `json:"sorted"`
	Examples []string `json:"examples"`
}

type WordList []*Word

func (w *Word) IsEnableConstruct(target string) bool {
	if len(target) < len(w.Sorted) {
		return false
	}

	i := 0
	j := 0

	for i < len(target) && j < len(w.Sorted) {
		if target[i] == w.Sorted[j] {
			i++
			j++
		} else if target[i] < w.Sorted[j] {
			i++
		} else {
			return false
		}
	}

	if i == len(target) && j != len(w.Sorted) {
		return false
	}

	return true
}

func (w *Word) CalcScore(firstQ, firstU int) int {
	list := strings.Split(w.Sorted, "")
	score := 0
	for i := 0; i < len(w.Sorted); i++ {
		score += config.AlphabetScore[list[i]]
	}

	countQ, countU := w.CountQU()
	// if countQ != 0 && countQ < countU {
	// 	score -= config.AlphabetScore["u"]
	// }

	score++ // bonus
	if countU > firstU+countQ {
		score = 0
	} else if countU > firstU {
		score -= config.AlphabetScore["u"] * (countU - firstU)
	}

	score *= score

	return score
}

func (wl WordList) FindEqualWord(target string, first int, last int) *Word {
	tmp := (first + last) / 2
	if tmp == first || tmp == last {
		return nil
	}

	if target == wl[tmp].Sorted {
		return wl[tmp]
	} else if target < wl[tmp].Sorted {
		return wl.FindEqualWord(target, first, tmp)
	} else {
		return wl.FindEqualWord(target, tmp, last)
	}
}

func (wl WordList) Len() int {
	return len(wl)
}

func (wl WordList) Swap(i, j int) {
	wl[i], wl[j] = wl[j], wl[i]
}

func (wl WordList) Less(i, j int) bool {
	return wl[i].Sorted < wl[j].Sorted
}

func (w *Word) CountQU() (int, int) {
	q := 0
	u := 0
	for i := 0; i < len(w.Sorted); i++ {
		if w.Sorted[i] > 'u' {
			break
		} else if w.Sorted[i] == 'q' {
			q++
		} else if w.Sorted[i] == 'u' {
			u++
		}
	}

	return q, u
}
