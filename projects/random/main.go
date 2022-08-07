package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
This simple guess game show how we can create random number,
in Go if we use random functions our result in every complile are the same
we can Seed the random with rand.seed(),
but still we get same number if we seed it with static number.
we can use time to seed it and get random numbers on every run.
*/


func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("You should enter number")
		return 
	}

	for i := 0; i < 5; i++ {
		randomNumber := rand.Intn(10) 
		if(randomNumber == guess) {
			fmt.Printf("\nYOU WIN. Number is %d", guess)
			return
		}
	}
	fmt.Println("\nYOU LOSE |:")

}