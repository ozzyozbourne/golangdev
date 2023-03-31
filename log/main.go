package main

import (
	"bufio"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	err := in.Err()
	if err != nil {
		return
	}

}
