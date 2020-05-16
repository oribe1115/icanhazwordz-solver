package solver

import (
	"fmt"
	"strings"

	"github.com/oribe1115/icanhazwordz-solver/lib"
)

func FindEqualWords(dictionary lib.WordList) {
	fmt.Println("input word")
	fmt.Print("> ")

	sortedTarget := lib.StringSort(lib.ReadLine())
	word := dictionary.FindEqualWord(sortedTarget, 0, len(dictionary)-1)
	if word == nil {
		fmt.Println("no word")
		return
	}

	fmt.Println(strings.Join(word.Examples, " "))
}

func FindContainAnagramWords(dictionary lib.WordList) {
	fmt.Println("input word")
	fmt.Print("> ")

	sortedTarget := lib.StringSort(lib.ReadLine())

	result := make([]string, 0)
	for _, word := range dictionary {
		if word.IsEnableConstruct(sortedTarget) {
			result = append(result, word.Examples...)
		}
	}

	fmt.Println(strings.Join(result, " "))
}
