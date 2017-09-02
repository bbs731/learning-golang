package main

import (
	"fmt"
)

type User struct{
	name string
	salary float64
}

func main() {
	//ch := make(chan int, 1)
	ch := make(chan *User)
	close(ch)
	// 如果从一个已经 closed channel 来读，  ok 会是 false
	// v 将会只是一个初始化的零值， 和 closed channel read 没啥关系
	// nil 和 零值是不同的， nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.

	v, ok := <-ch
	if ok {
		fmt.Printf("reading value %v from closed channel\n", v)
	}else{
		fmt.Printf("not reading ok from closed channel %v\n", v)

		if v == nil {
			println("v is nil")
		}
	}
}
