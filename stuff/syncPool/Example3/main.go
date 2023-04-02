package main

import "sync"

type temporaryCache struct {
	T string
}

func main() {

	pool := sync.Pool{
		New: func() interface{} {
			s := temporaryCache{}
			return &s
		},
	}

	data := pool.Get().(*temporaryCache)
	data.Reset()

}
