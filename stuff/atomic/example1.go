package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var waittime sync.WaitGroup

var atmvar int32

func hike(s string) {

	for i := 1; i < 18; i++ {
		// time.Sleep(time.Duration(rand.Intn(5)) * time.Microsecond)
		atomic.AddInt32(&atmvar, 1)
		fmt.Println(s, i, "->", atmvar)
	}
	waittime.Done()
}

func Example1() error {
	waittime.Add(2)

	go hike("test1")
	go hike("test2")
	waittime.Wait()
	fmt.Println("result :", atmvar)

	panic("failed")
}
