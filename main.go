package main

import (
	"github.com/oribe1115/icanhazwordz-solver/lib"
)

func main() {
	err := lib.CreateDictionary()
	if err != nil {
		panic(err)
	}
}
