package main

import "fmt"

func main() {
	a := 100 //Declared a integer variable a assiging value 100
	if a < 0 {
		fmt.Println("Value is negative")
	} else if a < 10 {
		fmt.Println("Value is single digit")
	} else if a < 100 {
		fmt.Println("Value is double digit")
	} else {
		fmt.Println("Value has multiple digits")
	}
}

// Output: Value has multiple digits
