package main

import "fmt"

var a1 int // Fibonacci input
var f1 int

func main() {
	fmt.Println("How many Fibonacci numbers does it represent?")
	fmt.Scan(&a1)
	fmt.Printf("We calculate up to this %d Fibonacci number:\n", Fibonacci2(a1))
	correction1()
	//f -= a
	println(f1)

}

func Fibonacci2(a1 int) int { //Fibonacci calculator function
	if a1 <= 1 {
		return a1
	}
	return Fibonacci2(a1-1) + Fibonacci2(a1-2)
}
func correction1() { //Fibonacci counter function
	//c := Fibonacci1(a)

	for i := 0; i < a1; {
		i += 1

		print(Fibonacci2(i))
		if i == a1 {
			print(" ... ")
			f1 += Fibonacci2(i)
		} else {
			print(" , ")
			f1 += Fibonacci2(i)
		}

	}
}
