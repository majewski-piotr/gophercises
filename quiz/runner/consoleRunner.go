
package runner

import (
	"encoding/csv"
	"fmt"
	"gophercises/quiz"
	"log"
	"os"
	"time"
)

type ConsoleRunner struct {
	Q quiz.Quiz
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
	for _, q := range cs.Q.Questions {
		var input string
		fmt.Println("Question:", q.Text)
		fmt.Scanln(&input)
		if q.Check(input) {
			cs.Q.CorrectAnswers++
		}
	}
}

func (cs *ConsoleRunner) RunWithTimer(timer *time.Timer) {
	for _, q := range cs.Q.Questions {
		fmt.Println("Question:", q.Text)

		answerChan := make(chan bool)
		go func() {
			var input string
			fmt.Scanln(&input)
			answerChan <- q.Check(input)
		}()

		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			return
		case answer := <-answerChan:
			if answer {
				cs.Q.CorrectAnswers++
			}
		}
	}
}
