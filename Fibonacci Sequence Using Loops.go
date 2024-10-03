package main

import "fmt"

var aa int // Fibonacci input

func main() {
	fmt.Println("How many Fibonacci numbers does it represent?")
	fmt.Scan(&aa)
	fmt.Printf("We calculate up to this %d Fibonacci number:\n", Fibonacci1(aa))
	correction()

}

func Fibonacci1(aa int) int { //Fibonacci calculator function
	if aa <= 1 {
		return aa
	}
	return Fibonacci1(aa-1) + Fibonacci1(aa-2)
}
func correction() { //Fibonacci counter function
	//c := Fibonacci1(a)
	for i := 0; i < aa; {
		i += 1
		print(Fibonacci1(i))
		if i == aa {
			print(" ... ")
		} else {
			print(" , ")
		}

	}
}
