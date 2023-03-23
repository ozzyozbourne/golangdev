package main

import "fmt"

func main() {
	pooper := []string{}

	//pooper[0] = "s;dmvsdmv"
	fmt.Println(len(pooper))
	pooper = nil
	pooper[0] = "s;dmvsdmv"
	fmt.Println(len(pooper))
}
