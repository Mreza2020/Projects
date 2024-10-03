package main

import "fmt"

var a int // Fibonacci input

func main() {
	fmt.Println("How many Fibonacci numbers do you want?\t")
	fmt.Scan(&a)
	println(Fibonacci(a)) //print Fibonacci

}

func Fibonacci(a int) int {
	if a <= 1 {
		return a
	} //To avoid stack overflow
	return Fibonacci(a-1) + Fibonacci(a-2) //Formula ==   fn = (fn-1) + (fn-2)

}
