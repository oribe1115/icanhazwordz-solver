package solver

import (
	"fmt"
	"time"

	"github.com/oribe1115/icanhazwordz-solver/lib"
	"github.com/sclevine/agouti"
	"github.com/siddontang/go/log"
)

var (
	pageURL   = "https://icanhazwordz.appspot.com/"
	page      *agouti.Page
	sleepTime time.Duration
)

func AutoSolver(dictionary lib.WordList, _sleepTime time.Duration) {
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, err := agoutiDriver.NewPage()
	if err != nil {
		log.Error(err)
		return
	}

	page.Navigate(pageURL)
	sleepTime = _sleepTime

	for i := 0; i < 10; i++ {
		fmt.Printf("turn: %d\n", i+1)
		err := autoSolve(dictionary)
		if err != nil {
			log.Error(err)
			return
		}
	}

	fmt.Printf("quit?\n> ")
	input := lib.ReadLine()
	if len(input) != 0 {
		return
	}
}

func autoSolve(dictionary lib.WordList) error {
	time.Sleep(sleepTime * time.Second)

	letterClass := page.AllByClass("letter")
	letterElement, err := letterClass.Elements()
	if err != nil {
		return err
	}

	target := ""

	for _, l := range letterElement {
		t, err := l.GetText()
		if err != nil {
			return err
		}

		target += t
	}

	fmt.Printf(" %s\n", target)

	answer, score := solver(dictionary, target)
	fmt.Printf("%s: %d\n", answer, score)

	inputField := page.FindByID("MoveField")
	inputField.Fill(answer)
	time.Sleep(sleepTime * time.Second)

	submitButton := page.FindByButton("Submit")
	err = submitButton.Click()
	if err != nil {
		return err
	}

	return nil
}
