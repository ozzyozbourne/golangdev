package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

const (
	vedio   string  = ""
	current int     = 10
	puller  float64 = 10.5
)

func main() {
	s := "awerwer/sserts/agreement.pdf"
	a, b := path.Split(s)
	fmt.Println(a, "  ", b)
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf(`skdfksdf ksndvlksn\ef\ef\ef\\n\n\n\n\n\n`)
	p := os.Args[1]
	p = p + strings.Repeat("!", utf8.RuneCountInString(p))
	p = strings.ToUpper(p)
	qw := `sdlvmsdkvskldv          `
	fmt.Println(strings.TrimRight(qw, "v "))
	fmt.Printf("%q\n", p)
	const (
		aw = iota + 10
		bw
		cw
		dw
	)
	ps := fmt.Println
	ps("jello aslises")
	ps("Enter the input")
	r := bufio.NewReader(os.Stdin)
	n, err := r.ReadString('\n')
	if err == nil {
		log.Println(n)
	} else {
		log.Fatalln(nil)
	}
}
