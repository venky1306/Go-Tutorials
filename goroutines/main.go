package main

import (
	"fmt"
	"sync"
)

var sg sync.WaitGroup

func main() {
	// sg.Add(2)

	// go a()
	// go b()

	// sg.Wait()
	// fmt.Println("Main goroutine exit")
	channels()
}

func a() {
	defer sg.Done()
	sg.Add(2)
	go aa()
	go ab()
}

func b() {
	defer sg.Done()
	sg.Add(2)
	go ba()
	go bb()
}

func aa() {
	defer sg.Done()
	fmt.Println("aa")
}

func ab() {
	defer sg.Done()
	fmt.Println("ab")
}

func ba() {
	defer sg.Done()
	fmt.Println("ba")
}

func bb() {
	defer sg.Done()
	fmt.Println("bb")
}
