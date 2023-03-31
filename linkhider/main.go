package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatalln("Please provide and input")
	}

	const (
		link  = "http"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = os.Args[1]
		size = len(text)
		buf  = make([]byte, 0, size)
	)

	in := false
	fmt.Printf("Lenght -> %d\n", len(text))
	for i := 0; i < size; i++ {
		c := text[i]
		switch c {
		case ' ', '\t', '\n':
			in = false
		}
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			buf = append(buf, link...)
			i += nlink
			continue
		}
		if in {
			c = mask
		}
		buf = append(buf, c)
	}
	fmt.Printf("%s\n", buf)
}
