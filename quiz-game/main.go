package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	timeout := flag.Int("timeout", 30, "Timeout in seconds")
	filename := flag.String("csv", "problems.csv", "Path to CSV file")
	flag.Parse()

	rc := loadRecords(*filename)
	q := newQuiz(rc)
	cs := consoleRunner{q: q}

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	cs.runWithTimer(timer)
	fmt.Println(cs.q.getResult())
}
