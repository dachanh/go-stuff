package main

import (
	"fmt"
)

type example struct {
	open int
}

func foo(m map[int]*example) {
	m[10] = &example{
		open: 2,
	}
}

func main() {
	var mapExample map[int]*example

	if mapExample == nil {
		fmt.Println("nil")
	}

	mapExample = make(map[int]*example)

	for i := 0; i < 5; i++ {
		mapExample[i] = &example{
			open: i,
		}
	}
	foo(mapExample)
	for idx, v := range mapExample {
		fmt.Println(idx, v)
	}

	testMap := make(map[string]string)
	testMap["test"] = "sds"

	for _, v := range testMap {
		fmt.Println(v)
	}

	// var arr [3]int
	arr := make([]int, 0, 3)

	if arr == nil {
		fmt.Println("sds")
	}

	// for i := 0; i < 10; i++ {
	// 	arr = append(arr, i)
	// }

	for _, v := range arr {
		fmt.Println(v)
	}
}
