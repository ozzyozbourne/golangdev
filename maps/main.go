package main

import "fmt"

func main() {
	s := map[string][]rune{}
	s["sdf"] = []rune{'2'}
	key, ok := s["sdc"]
	fmt.Printf("%v %v %v, %#v", key, ok, len(s), s)
	st := make(map[string][]rune, 10)
	fmt.Println(len(st))
}
