package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var i int32
	pool := sync.Pool{
		New: func() interface{} {
			i = i + 1
			fmt.Println("Creating new object...", i)

			time.Sleep(1 * time.Millisecond)
			return i
		},
	}

	numWorkers := 10
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			obj := pool.Get()
			fmt.Println("Obj= ", obj)
			defer pool.Put(obj)
		}()
	}
	wg.Wait()
}
