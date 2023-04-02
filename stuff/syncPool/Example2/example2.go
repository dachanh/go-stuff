package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheObj struct {
	C string
}

func main() {
	T := "sdasd"
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new object...")
			return &CacheObj{
				C: T,
			}
		},
	}

	testObject := pool.Get().(*CacheObj)
	fmt.Println(testObject.C) // print 1
	T = "2"
	pool.Put(testObject)
	time.Sleep(1 * time.Millisecond)
	testObject2 := pool.Get().(*CacheObj)
	fmt.Println(testObject2.C)
}
