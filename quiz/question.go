package quiz

type Question struct {
	Text   string
	Answer string
}

func (q Question) Check(answer string) bool {
	return answer == q.Answer
}

func (q Question) String() string {
	return "Question: " + q.Text + ", Answer: " + q.Answer
}
