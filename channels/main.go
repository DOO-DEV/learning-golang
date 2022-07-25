package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Channels are a safe way of transferring data between goroutines without
// using a mutex. You can send data to a channel in one goroutine, then
// consume data from the same channel in another goroutine. By default, a
//channel does not have space to store data, so you must simultaneously send
// and receive data from the channel
// to avoid a deadlock. A buffered
//channel allows you to allocate space to temporarily store data in the channel.

func main() {
	/*
	dataChan := make(chan int)
	// we can send and recive data from channel

	dataChan <- 10 
	// send

	n := <- dataChan 
	// recive

	fmt.Printf("n = %d\n", n) 
	// when run the program it is a deadlock
	// because channle are like portal, there is no space in unbuffred channel
	// for run this program we need to send and recive data simuntensly
	// it is possible with help of go routines
	// below program use go routines for send and recive through a channel
	*/


	/*

	unBufferdDataChan := make(chan int)
	//second go routine
	go func(){
		unBufferdDataChan <- 100
	}()
	// main go routine
	n := <- unBufferdDataChan

	fmt.Printf("n = %d\n", n)

	// we can solve this problem without go routine with Bufferd channel
	// it means channel has some space that we defined and we can send things to
	// that space and read it from that space

	bufferDataChan := make(chan int, 1) 
	// Bufferd channle with size 1

	bufferDataChan <- -100

	n = <- bufferDataChan

	fmt.Printf("n = %d\n", n)

	// if we try add new item to channel it will be deadlock 
	// because it has no space for new item for solving this
	// we should increase number of size channle
	// make(chan int, 2)

	*/

	// bufferDataChan <- -1000 => deadlock!
	
	/*
	someChannle := make(chan int)

	
	go func ()  {
		for i:= 0; i < 100; i++ {
			someChannle <- i
		}
		// add close for preventing deadlock on the main go routine
		close(someChannle)
	}()
	// this is also cause deadlock because its sends numbers to the channel
	// and when send 99 to channle it exit from go routin meanwhile the
	// main go routine try to get data from channle that doesnt recive any thing
	// and this is the time program go to deadlock!
	// for solving this we add a close() 
	// this function will close the channle and tell to recivers that 
	// there is no more thing to recive

	for i := range someChannle {
		fmt.Printf("i = %d\n", i)
	}

	*/

	// some little real example of how use channle
	dataChan := make(chan int)
	/*
	// this implemention will send result to the channle every one sec and it almost take 1000 sec
	go func(){
		for i := 0; i < 1000; i++ {
			result := DoWork()
			dataChan <- result
		}
		close(dataChan)
	}()

	*/

	// with this approach it take almose one sec to do all the work not 1000 sec (;
	go func ()  {
		// because we create go routine in every iteration we need to track these go routine
		// whene done. because of this we use Wait group
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1) // adding go routines to wait group
			go func ()  {
				defer wg.Done() // invoke whene go routine is done
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait() // wait until all iteration is done and then close channle
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}

}

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

// although channle are simple and fun to use but if we dont use theme well
// it can be out of hand
// and for most of problem some Mutex will solve our problem