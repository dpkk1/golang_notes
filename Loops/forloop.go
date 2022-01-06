package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i + 1)
		fmt.Printf(" ")
	}

	fmt.Println("")

	for _, i := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i)
	}
}

/*
Output:
1 2 3 4 5 6 7 8 9 10
1
2
3
4
5
*/
