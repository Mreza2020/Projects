package main

import "fmt"

var n int // lucas input

func main() {
	fmt.Println("How many lucas numbers do you want?\t")
	fmt.Scan(&n)
	result := lucas(n)
	fmt.Printf("The %dth Lucas number is: %d\n", n, result)
}

func lucas(n int) int {
	if n == 0 {
		return 2
	} else if n == 1 {
		return 1
	} else {
		return lucas(n-1) + lucas(n-2)
	}
}
