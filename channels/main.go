package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("hello", phrase)
	doneChan <- true
}
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello", phrase)
	doneChan <- true
}

// channels dont make go routines run, they only help in making goroutines in sync with data or commun so that mulitple go routines can work with each other
// <-done in main is used to tell that wait until this is executed

func main() {
	done := make(chan bool)
	// currently no output will be displayed
	go greet("nice to meet you", done)
	go greet("how are you", done)
	go slowGreet("how are you (slow)", done)
	go greet("I like this course", done)
	<-done
}
