package solver

import (
	"time"

	"github.com/sclevine/agouti"
	"github.com/siddontang/go/log"
)

func AutoSolver() {
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()

	page.Navigate("https://icanhazwordz.appspot.com/")
	log.Info(page.Title())
	time.Sleep(3 * time.Second)

	page.FindByLink("Help").Click()
	log.Info(page.Title())
	time.Sleep(3 * time.Second)
}
