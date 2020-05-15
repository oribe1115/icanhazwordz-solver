package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

func CreateDictionary() error {
	res, err := http.Get("https://icanhazwordz.appspot.com/dictionary.words")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	dataList := strings.Split(string(b), "\n")

	dataMap := map[string][]string{}

	for _, data := range dataList {
		sorted := StringSort(data)
		dataMap[sorted] = append(dataMap[sorted], data)
	}

	wordList := WordList{}

	for key := range dataMap {
		word := &Word{
			Sorted:   key,
			Examples: dataMap[key],
		}

		wordList = append(wordList, word)
	}

	fmt.Println(wordList[0])

	return nil
}

func StringSort(s string) string {
	list := strings.Split(s, "")

	for i := range list {
		list[i] = strings.ToLower(list[i])
	}

	sort.Strings(list)

	return strings.Join(list, "")
}
