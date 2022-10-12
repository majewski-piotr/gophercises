package main

import "fmt"

type quiz struct {
	questions      []question
	correctAnswers int
}

func newQuiz(csvRecords [][]string) quiz {
	q := quiz{}
	for _, r := range csvRecords {
		qst := question{r[0], r[1]}
		q.questions = append(q.questions, qst)
	}
	return q
}

func (q quiz) getResult() string {
	return fmt.Sprintf("You answered correctly %d out of %d questions",
		q.correctAnswers, len(q.questions))
}

func (q quiz) String() string {
	result := ""
	for _, v := range q.questions {
		result += v.String() + "\n"
	}
	return result
}
