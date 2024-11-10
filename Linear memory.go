package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() //Time start
	fmt.Println(h1(49))
	elapsed := time.Since(start) //Time end
	fmt.Println(elapsed)

}

func h1(g int) int {
	if 0 == g%2 {
		println("type: yes")
		for i := 0; i < g; i++ {
			print(i)
			print(",")

		}
	}

	if 1 == g%2 {
		println("type: no")
		for i := 0; i < g; i++ {
			print(i)
			print(",")

		}
	}
	println("\nnumber:")
	return g
}
