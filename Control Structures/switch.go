package main

import "fmt"

func main() {
	a := 10
	switch a {
	case 100: //
		fmt.Println("100")
	case 9:
		fmt.Println("9")
	default:
		fmt.Println("Doesn't matched!")
	}
}

// Prints Doesn't matched!
