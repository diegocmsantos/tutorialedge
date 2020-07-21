package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutine Tutorial")

	// sequential exceution of our compute function
	go compute(10)
	go compute(10)

	// we scan fmt for input and print that to our console
	var input string
	fmt.Scanln(&input)
}

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
