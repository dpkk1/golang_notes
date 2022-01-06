package main

import "fmt"

func main() {
	a := 1
	for a <= 10 {
		fmt.Print(a)
		a++
		fmt.Print(" ")
	}
}

//Prints 1 2 3 4 5 6 7 8 9 10
