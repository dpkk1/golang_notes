//How to handle errors in GO

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	for _, a := range os.Args[1:] {
		//Parse into integer
		i, err := strconv.Atoi(a) //store int in a & also return an error, if it's empty, check!
		if err != nil {
			//Handle the error
			panic(fmt.Sprintf("Invalid Value: %v", err))
		}
		sum += i
	}
	fmt.Printf("Sum = %v\n", sum)
}

/*
	Checking for Correct Input:
	Run the command using: go run errors.go	1 2 3 4 5 6
	Output: Sum = 21
*/

/*
	Checking for Wrong input: Run the command using: go run errors.go a b c
	Output: Error Message
	panic: Invalid Value: strconv.Atoi: parsing "a": invalid syntax
	goroutine 1 [running]:
	main.main()
*/
