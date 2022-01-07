/*
	Slices are variable length arrays for storing a single data type.
	Maps are key value pairs so that single values can be looked up quickly
*/

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)

	m := map[string]string{}
	m["first"] = "Deepak"
	m["last"] = "Kumar"
	fmt.Println(m)
}

/*
	Output:
				[1 2 3 4 5]
				map[first:Deepak last:Kumar]
*/
