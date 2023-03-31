package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

type text struct {
	title string
	words int
}

type book struct {
	text
	isbn  string
	title string
}

func main() {
	s1 := song{"wonderwall", "oasis"}
	s2 := song{"super sonic", "oasis"}

	b := s1 == s2
	fmt.Println(b)

	s1 = s2
	b = s1 == s2
	fmt.Println(b)

	ss := []song{s1, s2}

	p := playlist{"rockAndRoll", ss}
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	moby := book{text: text{"Moby Dick", 10000}, isbn: "12334234"}

	fmt.Printf("%+v\n", moby)

	moby.title = "hunky babe"
	fmt.Printf("%+v\n", moby)
	moby.text.title = "hunky babe"
	fmt.Printf("%+v\n", moby)
}
