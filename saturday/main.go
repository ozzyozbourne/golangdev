package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

type

func main() {
	as := "lksdnf"
	pl(reflect.TypeOf(12))
	pl(reflect.TypeOf(as))
	pl(fmt.Sprintf("%.10f", 3.14))
	seedSec := time.Now().Unix()
	rand.Seed(seedSec)
	pl(rand.Int())
	switch as := "amlly"; {
	case as == "bob", as == "sally":
		pl("Ataki Ataki")
	default:
		pl("Bay calm dowm")
	}
	s := [...]string{"123", "12"}
	pl(s)
	rs := [...]int{
		11: 1,
		12,
		13,
	}
	pf("%v\n", rs)

	const (
		SUN = 6 - iota
		SAT
		FRI
		THU
		WED
		TUE
		MON
	)

	wl := [...]string{
		TUE: "Its Tuesday!",
		FRI: "Finally Friday cometh",
		SUN: "Yippe Kaye Its Holiday",
		MON: "Let the Grind Commence",
		THU: "Getting closer patience pal",
		SAT: "Neflix and chill",
		WED: "Stuck in the middle",
	}
	pf("%v\n", wl)
	for i, v := range wl {
		pf("Index:-\t%d\t\t%s\n", i, v)
	}

	pf("%v\n", wl[SUN])
	pf("%v\n", wl[WED])
	pf("%v\n", wl[MON])

}
