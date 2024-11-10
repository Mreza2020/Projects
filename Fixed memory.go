package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() //Time start

	fmt.Println(example1(45))

	elapsed := time.Since(start) //Time end
	fmt.Println(elapsed)
}

func example1(i int) int {
	return i
}
