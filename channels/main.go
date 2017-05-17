package main

import (
	"fmt"
	"math/rand"
	"time"
)

//someLongFunc is a function that might
//take a while to complete, so we want
//to run it on its own go routine
// defines a channel of ints -> chan int
func someLongFunc(ch chan int) {
	r := rand.Intn(2000)
	d := time.Duration(r)
	time.Sleep(time.Millisecond * d)
	// write the value r into the channel ch
	//if the channel is not full, keep running
	ch <- r

}

func main() {
	//TODO:
	//create a channel and call
	// someLongFunc() on a go routine
	//passing the channel so that
	//someLongFunc() can communicate
	//its results
	var n:= 10
	rand.Seed(time.Now().UnixNano())
	fmt.Println("starting long-running func...")
	//creates empty buffer of ints
	ch := make(chan int)
	// By using these concurrent calls to someLongFunc(), we can have many go routines being run at the same time
	for i := 0; i < n; i++ {
		// the go keyword makes it run concurrently
		// if you take the go out, the calls will be called serially, so all results will be reported ONLY when all
		// calls are completed
		// How many routines are being run in a computer depends on how many cores GO can access
		go someLongFunc(ch)

	}
	for i := 0; i < n; i++ {
		//read the result of ch into result
		// you use the weird left facing arrow to indicate where it should read the results to
		result := <-ch
		fmt.Printf("result was %d", result)
	}
	fmt.Printf("took %v", time.Since(start))

}
