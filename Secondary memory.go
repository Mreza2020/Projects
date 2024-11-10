package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() //Time start
	var g = [3]int{80, 100, 60}

	var k = [3]int{90, 15, 8}

	n := 1

	for i := 0; i < n; i++ {
		g[2] += k[2]
		g[0] += k[0]
		g[1] += k[1]

	}
	println(g[0], g[1], g[2])
	elapsed := time.Since(start) //Time end
	fmt.Println(elapsed)
}
