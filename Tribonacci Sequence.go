package main

import "fmt"

// Tribonacci function to calculate the n-th Tribonacci number
func Tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	// Create an array to store Tribonacci numbers
	trib := make([]int, n+1)
	trib[0], trib[1], trib[2] = 0, 1, 1

	for i := 3; i <= n; i++ {
		trib[i] = trib[i-1] + trib[i-2] + trib[i-3]
	}

	return trib[n]
}

func main() {
	var n int // Example: Get the ...th Tribonacci number
	fmt.Println("How many Tribonacci numbers does it represent?")
	fmt.Scan(&n)
	fmt.Printf("The %d-th Tribonacci number is: %d\n", n, Tribonacci(n))
}
