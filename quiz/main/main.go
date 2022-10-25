package main

import (
	"flag"
	"fmt"
	"gophercises/quiz"
	"gophercises/quiz/runner"
	"time"
)

func main() {
	timeout := flag.Int("timeout", 30, "Timeout in seconds")
	filename := flag.String("csv", "problems.csv", "Path to CSV file")
	flag.Parse()

	rc := runner.LoadRecords(*filename)
	q := quiz.New(rc)
	cs := runner.ConsoleRunner{Q: q}

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	cs.RunWithTimer(timer)
	fmt.Println(cs.Q.GetResult())
}
