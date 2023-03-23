package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type StringHeader struct {
	pointer uintptr
	length  int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}

func dumpSlice(s []string) {
	ptr := *(*SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}

func main() {
	a := ""
	dump(a)
	h := "Hello"
	dump(h)
	dump("Hello")
	dump(h[:1])
	fmt.Printf("%s\n", strings.Repeat(" -*- ", 20))
	for i, _ := range h {
		dump(h[i : i+1])
	}
	dump(string([]byte(h)))
	dump(string([]byte(h)))
	dump(string([]rune(h)))

	ac := []string{"sdf", "fsdf"}
	dumpSlice(ac)
	dumpSlice(ac[:1])
}
