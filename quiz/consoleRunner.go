package quiz

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type ConsoleRunner struct {
	Q quiz
}

func LoadRecords(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	return records
}

func (cs *ConsoleRunner) Run() {
	for _, q := range cs.Q.questions {
		var input string
		fmt.Println("Question:", q.text)
		fmt.Scanln(&input)
		if q.check(input) {
			cs.Q.correctAnswers++
		}
	}
}

func (cs *ConsoleRunner) RunWithTimer(timer *time.Timer) {
	for _, q := range cs.Q.questions {
		fmt.Println("Question:", q.text)

		answerChan := make(chan bool)
		go func() {
			var input string
			fmt.Scanln(&input)
			answerChan <- q.check(input)
		}()

		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			return
		case answer := <-answerChan:
			if answer {
				cs.Q.correctAnswers++
			}
		}
	}
}
