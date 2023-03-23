package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"unsafe"
)

func main() {
	var todo []string
	st := []string{"s;dlvj", "sldijv", "lskdjv", "sd;vjs;ldmv;lsdjv"}
	todo = append(todo, st...)
	s.Show("Todo", todo, st)
	sd := st[:1]
	s.Show("slice sd-> ", sd)
	s.Show("slice st-> ", st)
	sd[0] = "Greeta Van Fleet"
	s.Show("slice sd-> ", sd)
	s.Show("slice st-> ", st)

	news := append([]string(nil), st...)
	s.Show("klsdnf", news)
	s.Show("asdefsad", st[1:3])

	qw := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	as := qw[5:]

	var p *[]int = &as
	fmt.Println(p)
	fmt.Printf("%p\n", &qw)
	fmt.Println(unsafe.Sizeof(as))
	fmt.Println(unsafe.Sizeof(st))
	fmt.Println(unsafe.Sizeof(news))
	var asm []int
	fmt.Printf("Backing arr -> %p", &asm)

}
