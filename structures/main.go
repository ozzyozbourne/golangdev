package main

import "fmt"

func main() {
	type person struct {
		firstName, lastName string
		age                 int
	}
	p := person{
		firstName: "osaid",
		lastName:  "Khan",
		age:       29,
	}

	var q person
	q.firstName = "homer"
	q.lastName = "simson"
	q.age = 45

	var r person

	fmt.Printf("%+v\n", p)
	fmt.Printf("%+v\n", q)
	fmt.Printf("%#v\n", r)
}
