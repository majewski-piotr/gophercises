package quiz

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewQuiz(t *testing.T) {
	q := New(getTestRecords())

	if q.questions[0].text != "1+1" {
		t.Errorf("Incorrect first question, should be 1+1, is %v", q.questions[0].text)
	}
	if q.questions[0].answer != "2" {
		t.Errorf("Incorrect first question, should be 2, is %v", q.questions[0].answer)
	}
	if len(q.questions) != 3 {
		t.Errorf("Incorrect number of questions, should be 3, is %v", len(q.questions))
	}
}

func TestStringQuiz(t *testing.T) {
	q := New(getTestRecords())
	result := q.String()
	expected := "Question: 1+1, Answer: 2\nQuestion: 2+2, Answer: 4\nQuestion: 4+2, Answer: 6"
	if strings.EqualFold(result, expected) {
		t.Errorf("Incorrect printed value, should be \n%v, is \n%v", expected, result)
	}
}

func TestGetResult(t *testing.T) {
	q := New(getTestRecords())
	q.correctAnswers = 2

	result := q.GetResult()
	expected := fmt.Sprintf("You answered correctly 2 out of %d", len(q.questions))

	if strings.EqualFold(result, expected) {
		t.Errorf("Incorrect returned value, should be \n%v, is \n%v", expected, result)
	}
}

func getTestRecords() [][]string {
	return [][]string{
		{"1+1", "2"},
		{"2+2", "4"},
		{"4+2", "6"},
	}
}
