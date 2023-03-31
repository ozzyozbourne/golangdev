package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())

		if len(word) > 2 {
			words[word] = true
		}
	}

	query := "sun"

	if _, ok := words[query]; ok {
		fmt.Println("yes \t", query)
	}

}
