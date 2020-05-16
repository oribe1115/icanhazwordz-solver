package solver

import (
	"fmt"
	"time"

	"github.com/oribe1115/icanhazwordz-solver/lib"
	"github.com/sclevine/agouti"
	"github.com/siddontang/go/log"
)

var (
	pageURL = "https://icanhazwordz.appspot.com/"
)

func AutoSolver(sleepTime time.Duration) {
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()

	page.Navigate(pageURL)
	time.Sleep(sleepTime * time.Second)

	letterClass := page.AllByClass("letter")
	letterElement, err := letterClass.Elements()
	if err != nil {
		log.Error(err)
		return
	}

	target := ""

	for _, l := range letterElement {
		t, err := l.GetText()
		if err != nil {
			log.Error(err)
			return
		}

		target += t
	}

	fmt.Println(target)

	fmt.Printf("quit?\n> ")
	input := lib.ReadLine()
	if len(input) != 0 {
		return
	}
}
