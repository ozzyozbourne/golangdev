package main

import "fmt"

type s struct {
	a []string
	b int
}

func main() {
	fmt.Println(nakedReturn())
	fmt.Printf("str2: %q\n", "string")

}

func nakedReturn() (m int) {
	return
}
