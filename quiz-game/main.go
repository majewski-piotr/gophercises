package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	timeout := flag.Int("timeout", 30, "Timeout in seconds")
	flag.Parse()
	args := flag.Args()

	rc := loadRecords(args[0])
	q := newQuiz(rc)
	cs := consoleRunner{q: q}
	timeoutChannnel := make(chan bool)
	go runAgainstTimer(&cs, timeoutChannnel)
	go timer(*timeout, timeoutChannnel)
	<-timeoutChannnel
	fmt.Println(cs.q.getResult())
}

func timer(timeout int, ch chan bool) {
	fmt.Println(timeout)
	time.Sleep(time.Second * time.Duration(timeout))
	fmt.Println("Time out !")
	ch <- true
}
func runAgainstTimer(cs *consoleRunner, ch chan bool) {
	cs.run()
	ch <- true
}
