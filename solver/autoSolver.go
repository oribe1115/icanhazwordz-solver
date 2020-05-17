package solver

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/oribe1115/icanhazwordz-solver/lib"
	"github.com/sclevine/agouti"
	"github.com/siddontang/go/log"
)

var (
	pageURL   = "https://icanhazwordz.appspot.com/"
	NICK_NAME = ""
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
	NICK_NAME = os.Getenv("NICK_NAME")

	fmt.Println("targetScore")
	fmt.Print("> ")
	targetScore, err := strconv.Atoi(lib.ReadLine())
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println("limit")
	fmt.Print("> ")
	limit, err := strconv.Atoi(lib.ReadLine())
	if err != nil {
		log.Error(err)
		return
	}

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
		errorFlag := false
		for j := 0; j < 10; j++ {
			_, _, err := autoSolve(dictionary, page, 0)
			if err != nil {
				log.Error(err)
				errorFlag = true
				return
			}
		}

		if errorFlag {
			page.FindByLink("Start a new game").Click()
			continue
		}

		printedScore := page.FindByXPath("/html/body/table[1]/tbody/tr/td[2]/table/tbody/tr[1]/td[2]")
		got, err := printedScore.Text()
		if err != nil {
			log.Error(err)
		}

		score, err := strconv.Atoi(got)
		if err != nil {
			log.Error(err)
		} else {
			fmt.Printf("%3d times: %4d\n", i, score)
			if score > targetScore {
				err = submitHighScore(page)
				if err != nil {
					log.Error(err)
				}
				fmt.Println("submit as high score")
				break
			}
		}

		page.FindByLink("Start a new game").Click()
	}

	fmt.Println("end")
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
	if score == 0 {
		passButton := page.FindByButton("PASS")
		err = passButton.Click()
		if err != nil {
			return "", 0, err
		}

		return fmt.Sprintf(" %s\n%s: PASS\n", target, answer), score, nil
	}

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

func submitHighScore(page *agouti.Page) error {
	nameHolder := page.FindByName("NickName")
	err := nameHolder.Fill(NICK_NAME)
	if err != nil {
		return err
	}

	urlHolder := page.FindByName("URL")
	err = urlHolder.Fill("https://github.com/oribe1115/icanhazwordz-solver")
	if err != nil {
		return err
	}

	checkRobot := page.FindByID("AgentRobot")
	err = checkRobot.Click()
	if err != nil {
		return err
	}

	recordButton := page.FindByButton("Record!")
	err = recordButton.Click()
	if err != nil {
		return err
	}

	return nil
}
