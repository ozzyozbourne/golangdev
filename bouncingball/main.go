package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	height int = 10
	width  int = 50
)

var board [height][width]bool

var i int = 1

func main() {
	var start, stop int

	args := os.Args[1:]
	fmt.Printf("%#v\n", os.Args)
	if len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	} else {
		log.Fatalln("Please enter start and stop numbers")
	}

	fmt.Printf("%-9s %-9s %-13s %-10s\n%s\n", "Literal", "Decimal", "Hexideciaml", "encoded",
		strings.Repeat("-", 45))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-5c %-5s %-10[1]d %-10[1]x % -12[3]x\n", n, "=>", string(n))
	}

}
