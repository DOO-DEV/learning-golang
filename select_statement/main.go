package main

import "fmt"

// What a Select does actully is lets go ruotine wait on multiple comminucations

func main() {
	myChannle := make(chan string)
	anotherChannle := make(chan string)

	go func() {
		myChannle <- "data"
	}()

	go func() {
		anotherChannle <- "anotherData"
	}()

	select {
	case msgFromMyChannle := <-myChannle:
		fmt.Println(msgFromMyChannle)
	case msgFromAnotherChannle := <- anotherChannle:
		fmt.Println(msgFromAnotherChannle)
	}

}