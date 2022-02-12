package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//callSomethingChannel()
	//callSomethingWithWaitGroup()
	now := time.Now()
	callWorker(1)
	final := time.Since(now)
	fmt.Println("time with one worker %v", final)

	now4 := time.Now()
	callWorker(4)
	final4 := time.Since(now4)
	fmt.Println("time with four workers %v", final4)
}

func callWorker(k int) {
	n := 20
	jobs := make(chan int, n)
	results := make(chan int, n)

	for g := 0; g < k; g++ {
		go worker(jobs, results)
	}

	for i := 0; i < n; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < n; j++ {
		<-results
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func twoChannels() {
	c1 := make(chan string)
	c2 := make(chan string)

	t1 := time.NewTicker(time.Millisecond * 500)
	t2 := time.NewTicker(time.Second * 2)

	go func() {
		for ; ; <-t1.C {
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for ; ; <-t2.C {
			c2 <- "Every two seconds"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}
}

func callSomethingWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		something("some")
		wg.Done()
	}()

	wg.Wait()

}

func something(thing string) {
	t := time.NewTicker(time.Second)

	val := true
	counter := 0
	now := time.Now()
	for ; val; <-t.C {
		counter++
		fmt.Println(counter, val, time.Since(now), thing)
		if counter == 5 {
			val = false
		}
	}
}

func callSomethingChannel() {
	c := make(chan string)
	go somethingChannel("channel", c)

	// if this channel is not closed, it will create a deadlock
	// basically this will keep waiting while the routine is not sending any message
	for {
		msg, open := <-c
		fmt.Println(msg)

		if !open {
			break
		}
	}
}

func somethingChannel(thing string, c chan string) {

	t := time.NewTicker(time.Second)

	val := true
	counter := 0
	now := time.Now()
	for ; val; <-t.C {
		c <- thing
		counter++
		fmt.Println(counter, val, time.Since(now), thing)
		if counter == 5 {
			val = false
		}
	}

	// if this channel is not closed, it will create a deadlock
	// basically this will keep waiting while the routine is not sending any message
	close(c)

}
