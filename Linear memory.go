package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() //Time start
	println(h(5000))
	elapsed := time.Since(start) //Time end
	fmt.Println(elapsed)
}

func h(d int) int {

	f := 0
	for i := 0; i < d; i++ {
		f += 5

	}
	return f

}
