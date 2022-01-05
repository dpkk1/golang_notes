package main

import "fmt"

func main() {

	//INT DATA TYPE

	// := this operator used for declaring a new variable that doesn't exist, Automatically detect the type.
	i := 100
	//Declaring j as int via using var, indicating int data type
	var j int = 1000
	fmt.Printf("%v + %v = %v\n", i, j, i+j) //print 1100

	//FLOAT DATA TYPE
	f := 1.5
	fmt.Printf("f = %v\n", f) //print 1.5
}

/*
	Output:
		100 + 1000 = 1100
		f = 1.5
*/
