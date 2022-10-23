package main

import (
	"flag"
	"fmt"
	"gophercises/quiz"
	"time"
)

func main() {
	timeout := flag.Int("timeout", 30, "Timeout in seconds")
	filename := flag.String("csv", "problems.csv", "Path to CSV file")
	flag.Parse()

	rc := quiz.LoadRecords(*filename)
	q := quiz.New(rc)
	cs := quiz.ConsoleRunner{Q: q}

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	cs.RunWithTimer(timer)
	fmt.Println(cs.Q.GetResult())
}
