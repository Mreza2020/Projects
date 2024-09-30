package main

import "fmt"

var f int

func main() {
	var d string
	var x int
	var y int

	fmt.Println("What math operation do you want to perform?")
	fmt.Println("+ , - , * , / , %")
	fmt.Scan(&d) //Taking mathematical operations
	fmt.Print("The first value\t")
	fmt.Scan(&x) //The first value
	fmt.Printf("\n%v\n\n", d)
	fmt.Print("The second value ")
	fmt.Scan(&y) //The second value
	if d == "+" {
		positive(x, y)
	}
	if d == "-" {
		minus(x, y)
	}
	if d == "*" {
		multiplication(x, y)
	}
	if d == "/" {
		division(x, y)
	}
	if d == "%" {
		remaining(x, y)
	}

	fmt.Printf("%v", f)
}

func positive(x, y int) {
	f = x + y
}
func minus(x, y int) {
	f = x - y
}
func multiplication(x, y int) {
	f = x * y
}
func division(x, y int) {
	f = x / y
}
func remaining(x, y int) {
	f = x % y
}//heloo
