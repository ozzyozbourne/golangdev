package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"userName"`
	Permissions permissions `json:"perm"`
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}
	var u []user
	err := json.Unmarshal(input, &u)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", u)
}
