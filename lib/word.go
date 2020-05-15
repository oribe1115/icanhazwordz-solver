package lib

// Word ソートされた単語ごとの構造体
type Word struct {
	Sorted   string   `json:"sorted"`
	Examples []string `json:"examples"`
}

type WordList []*Word

func (wl WordList) Len() int {
	return len(wl)
}

func (wl WordList) Swap(i, j int) {
	wl[i], wl[j] = wl[j], wl[i]
}

func (wl WordList) Less(i, j int) bool {
	return wl[i].Sorted < wl[j].Sorted
}
