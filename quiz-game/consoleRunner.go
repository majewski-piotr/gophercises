package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type consoleRunner struct {
	q quiz
}

func loadRecords(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	return records
}

func (cs *consoleRunner) run() {
	for _, q := range cs.q.questions {
		var input string
		fmt.Println("Question:", q.text)
		fmt.Scanln(&input)
		if q.check(input) {
			cs.q.correctAnswers++
		}
	}
}

func (cs *consoleRunner) runWithTimer(timer *time.Timer) {
	for _, q := range cs.q.questions {
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
				cs.q.correctAnswers++
			}
		}
	}
}
