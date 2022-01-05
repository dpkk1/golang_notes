package main

import "fmt"

func main() {
	/*
		:= this operator used for declaring a new variable that doesn't exist, Automatically detects the type.
	*/
	a := true  //Declared 'a' variable as true
	b := false //Declared 'b' variable as false

	//for printing true/false => use %v
	fmt.Printf("a = %v, b = %v, and a==b = %v", a, b, a == b)
}

/*
	The output of the program is:
	a = true, b = false, and a==b = false [ True != False ]

*/
