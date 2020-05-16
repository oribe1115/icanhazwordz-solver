package solver

import (
	"fmt"
	"time"

	"github.com/oribe1115/icanhazwordz-solver/lib"
	"github.com/sclevine/agouti"
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

	// page.FindByLink("Help").Click()

	fmt.Printf("quit?\n> ")
	input := lib.ReadLine()
	if len(input) != 0 {
		return
	}
}
