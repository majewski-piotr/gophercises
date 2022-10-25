// Package implements running quiz
package quiz

import "fmt"

type Quiz struct {
	Questions      []Question
	CorrectAnswers int
}

func New(csvRecords [][]string) Quiz {
	q := Quiz{}
	for _, r := range csvRecords {
		qst := Question{r[0], r[1]}
		q.Questions = append(q.Questions, qst)
	}
	return q
}

func (q Quiz) GetResult() string {
	return fmt.Sprintf("You answered correctly %d out of %d questions",
		q.CorrectAnswers, len(q.Questions))
}

func (q Quiz) String() string {
	result := ""
	for _, v := range q.Questions {
		result += v.String() + "\n"
	}
	return result
}
