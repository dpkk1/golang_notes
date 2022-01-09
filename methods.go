package main

import (
	"fmt"
	"time"
)

//Struct Person
type Person struct {
	First string
	Last  string
	Dob   time.Time
}

// NewPerson creates a new Person
func NewPerson(first string, last string, year, month, day int) *Person {
	return &Person{
		First: first,
		Last:  last,
		Dob:   time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

//Prints Hello + "First name of the person"
func (p Person) SayHello() {
	fmt.Printf("Hello %s\n", p.First)
}

//Prints Hello + "Name of the person"
func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

//Prints the age of person
func (p Person) GetAge() int {
	d := time.Since(p.Dob)
	return (int(d.Hours()) / 24 / 365)
}

func main() {
	sayHello("Deepak!")

	p := NewPerson("Deepak", "Kumar", 2000, 10, 31)
	p.SayHello()
	fmt.Println(p.GetAge())
}

/*
	Output:
		Hello Deepak!
		Hello Deepak
		21
*/
