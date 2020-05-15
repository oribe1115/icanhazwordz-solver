package lib

// Word ソートされた単語ごとの構造体
type Word struct {
	Sorted   string
	Examples []string
}

type WordList = []*Word
