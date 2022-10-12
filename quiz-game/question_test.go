package main

import (
	"strings"
	"testing"
)

func TestStringQuestion(t *testing.T) {
	q := question{"1+1", "2"}
	result := q.String()
	expected := "Question: 1+1, Answer: 2\nQuestion: 2+2, Answer: 4\nQuestion: 4+2, Answer: 6"
	if strings.EqualFold(result, expected) {
		t.Errorf("Incorrect printed value, should be \n%v, is \n%v", expected, result)
	}
}

func TestCheckQuestion(t *testing.T) {
	q := question{"1+1", "2"}
	result := q.check("2")
	expected := true
	if result != expected {
		t.Errorf("Incorrect printed value, should be \n%v, is \n%v", expected, result)
	}

	result = q.check("3")
	expected = false
	if result != expected {
		t.Errorf("Incorrect printed value, should be \n%v, is \n%v", expected, result)
	}
}
