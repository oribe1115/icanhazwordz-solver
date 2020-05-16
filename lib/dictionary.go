package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

var (
	dataURL        = "https://icanhazwordz.appspot.com/dictionary.words"
	dictionaryFile = "dictionary.json"
)

func GetDictionary() (WordList, error) {
	if !fileExists(dictionaryFile) {
		return createDictionary()
	}

	data, err := ioutil.ReadFile(dictionaryFile)
	if err != nil {
		return nil, err
	}

	var wordList WordList
	err = json.Unmarshal(data, &wordList)
	if err != nil {
		return nil, err
	}

	return wordList, nil
}

func createDictionary() (WordList, error) {
	res, err := http.Get(dataURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	dataList := strings.Split(string(b), "\n")

	dataMap := map[string][]string{}

	for _, data := range dataList {
		if data == "" {
			continue
		}
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

	sort.Sort(wordList)

	j, _ := json.Marshal(wordList)

	err = ioutil.WriteFile(dictionaryFile, ([]byte)(j), os.ModePerm)
	if err != nil {
		return nil, err
	}

	return wordList, nil
}

func StringSort(s string) string {
	list := strings.Split(strings.ToLower(s), "")

	sort.Strings(list)

	return strings.Join(list, "")
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
