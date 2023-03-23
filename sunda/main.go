package main

import (
	"fmt"
	st "github.com/inancgumus/prettyslice"
	"reflect"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	as := s[:4:4]
	fmt.Println(cap(s), "	", cap(as))
	fmt.Printf("%p\t%p\n", &s, &as)
	st.Show("One", s)
	st.Show("Two", as)
	v := reflect.ValueOf(as)
	fmt.Println(v.Pointer())
	vt := reflect.ValueOf(s)
	fmt.Println(vt.Pointer())

}

//
//func fetch() [][]int {
//
//
//
//
//}
