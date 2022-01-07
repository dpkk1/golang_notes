/*
	A Pointer is a special data type which stores the memory address and that memory address points to another location/memory with the data which is actually stored
*/

package main

import "fmt"

func main() {
	s := "My String"
	sptr := &s
	fmt.Println(s)     //Prints My String
	fmt.Println(sptr)  // Prints the address of My String: 0xc000088220
	fmt.Println(*sptr) // Prints My String (dereference the pointer)
}
