package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I am learning Go lang"
	fmt.Printf("%s\n", s) //prints, I am learning Go lang

	//Using string package
	fmt.Printf("Ends with string lang? %v\n ", strings.HasSuffix(s, "lang")) //prints true
}

/*
	Output:
		I am learning Go lang
		Ends with string lang? true
*/
