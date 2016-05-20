// A concurrent prime sieve that reports output every few seconds

package main

import (
	"fmt"
	"time"
)

const DELAY = 3 //seconds

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func RunTimer(timer chan bool) {
	for {
		time.Sleep(time.Duration(DELAY) * time.Second)
		timer <- true
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	timer := make(chan bool)
	go RunTimer(timer)
	lastIndex, lastPrime := 0, -1
	for { // run forever
		select {
		case prime := <-ch:
			lastIndex = lastIndex + 1
			lastPrime = prime
			if lastIndex == 1 {
				fmt.Printf("Prime number %d is %v\n", lastIndex, lastPrime)
			}
			ch1 := make(chan int)
			go Filter(ch, ch1, prime)
			ch = ch1
		case <-timer:
			fmt.Printf("Prime number %d is %v\n", lastIndex, lastPrime)
		}
	}
}
