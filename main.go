package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum, tree = 1, 2
	fmt.Println(sum, tree)
	summer, tree := 123, 209384
	fmt.Println(summer, tree)
	s, _ := strconv.Atoi("123")
	fmt.Println(s)
	q, _ := strconv.ParseInt("123", 10, 0)
	fmt.Println(q)
	a := strconv.Itoa(234)
	fmt.Println(a)
}
