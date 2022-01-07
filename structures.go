/*
	Structures: Provide a way to combine data with different types into a single data structure.
*/

package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p := Person{
		First: "Deepak",
		Last:  "Kumar",
		Age:   21,
	}
	
	fmt.Println(p) //Prints {Deepak Kumar 21}
}
