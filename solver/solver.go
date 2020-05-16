package solver

import (
	"fmt"
	"strconv"
	"time"

	"github.com/oribe1115/icanhazwordz-solver/lib"
	"github.com/sclevine/agouti"
	"github.com/siddontang/go/log"
)

var (
	pageURL = "https://icanhazwordz.appspot.com/"
)

func AutoSolver(dictionary lib.WordList, sleepTime time.Duration, logMode bool) {
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, err := agoutiDriver.NewPage()
	if err != nil {
		log.Error(err)
		return
	}

	page.Navigate(pageURL)
	logBuffer := ""
	totalScore := 0

	for i := 0; i < 10; i++ {
		logs, score, err := autoSolve(dictionary, page, sleepTime)
		if err != nil {
			log.Error(err)
			return
		}
		totalScore += score
		if logMode {
			fmt.Printf("turn: %d\n%s", i+1, logs)
		} else {
			logBuffer += fmt.Sprintf("turn: %d\n%s", i+1, logs)
		}
	}

	if !logMode {
		fmt.Printf(logBuffer)
	}

	fmt.Printf("totalScore: %d\n", totalScore)

	fmt.Printf("quit?\n> ")
	input := lib.ReadLine()
	if len(input) != 0 {
		return
	}
}

func AutoSolverToBeFirst(dictionary lib.WordList) {
	fmt.Println("targetScore")
	fmt.Print("> ")
	targetScore, err := strconv.Atoi(lib.ReadLine())
	if err != nil {
		log.Error(err)
		return
	}
	limit := 10

	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, err := agoutiDriver.NewPage()
	if err != nil {
		log.Error(err)
		return
	}

	page.Navigate(pageURL)

	for i := 0; i < limit; i++ {
		totalScore := 0
		for j := 0; j < 10; j++ {
			_, score, err := autoSolve(dictionary, page, 0)
			if err != nil {
				log.Error(err)
				return
			}
			totalScore += score
		}

		fmt.Printf("%3dtimes: %4d\n", i, totalScore)
		if totalScore > targetScore {
			break
		}

		page.FindByLink("Start a new game").Click()
	}

	fmt.Printf("quit?\n> ")
	input := lib.ReadLine()
	if len(input) != 0 {
		return
	}
}

func autoSolve(dictionary lib.WordList, page *agouti.Page, sleepTime time.Duration) (string, int, error) {
	err := page.Refresh()
	if err != nil {
		return "", 0, err
	}
	time.Sleep(sleepTime * time.Second)

	letterClass := page.AllByClass("letter")
	letterElement, err := letterClass.Elements()
	if err != nil {
		return "", 0, err
	}

	target := ""

	for _, l := range letterElement {
		t, err := l.GetText()
		if err != nil {
			return "", 0, err
		}

		target += string(t[0])
	}

	answer, score := solver(dictionary, target)

	inputField := page.FindByID("MoveField")
	inputField.Fill(answer)
	time.Sleep(sleepTime * time.Second)

	submitButton := page.FindByButton("Submit")
	err = submitButton.Click()
	if err != nil {
		return "", 0, err
	}

	return fmt.Sprintf(" %s\n%s: %d\n", target, answer, score), score, nil
}
