package main

type question struct {
	text   string
	answer string
}

func (q question) check(answer string) bool {
	return answer == q.answer
}

func (q question) String() string {
	return "Question: " + q.text + ", Answer: " + q.answer
}
