package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Printf("Scanned lines -> %s\n", in.Text())
		fmt.Printf("Scanned Bytes -> %v\n", in.Bytes())
		fmt.Printf("Scanned Bytes -> %b\n", in.Bytes())
	}
}
