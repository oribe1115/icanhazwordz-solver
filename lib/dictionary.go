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

	wordList := strings.Split(string(b), "\n")

	fmt.Println(wordList[0])

	po := StringSort(wordList[0])
	fmt.Println(po)

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
