package main

import (
	"fmt"

	"github.com/venky1306/struct/test"
)

func main() {

	chinnu := test.Venky
	chinnu.ID = 6
	fmt.Println(test.Venky.ID)
	fmt.Println((chinnu).ID)
}
