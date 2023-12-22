package main

import (
	"fmt"
	"time"
)

func channels() {
	ch := make(chan int)
	go ca(ch)
	go cb(ch)
	time.Sleep(time.Second * 2)
}

func ca(ch chan int) {
	fmt.Println("before a")
	ch <- 12
	time.Sleep(time.Second)
	fmt.Println("after a")
}

func cb(ch chan int) {
	fmt.Println("before b")
	val := <-ch
	fmt.Printf("value %d\n", val)
	fmt.Println("after b")
}
