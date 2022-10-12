package main

import (
	"encoding/csv"
	"fmt"
	"os"
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
	fmt.Println(cs.q.getResult())
}
