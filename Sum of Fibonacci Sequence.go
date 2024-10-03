package main

import "fmt"

var a int // Fibonacci input

func main() {
	fmt.Println("How many Fibonacci numbers does it represent?")
	fmt.Scan(&a)
	fmt.Printf("We calculate up to this %d Fibonacci number:\n", Fibonacci1(a))
	correction()

}

func Fibonacci1(a int) int { //Fibonacci calculator function
	if a <= 1 {
		return a
	}
	return Fibonacci1(a-1) + Fibonacci1(a-2)
}
func correction() { //Fibonacci counter function
	//c := Fibonacci1(a)
	for i := 0; i < a; {
		i += 1
		print(Fibonacci1(i))
		if i == a {
			print(" ... ")
		} else {
			print(" , ")
		}

	}
}
