package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	a := "ØÙÚâæ😀😉😊"
	b := []byte(a)
	b[0] = '1'
	b[1] = '3'
	fmt.Printf("%v % [1]x\n", []byte(a))
	fmt.Printf("%v\n", b)
	fmt.Printf("%v % [1]x\n", []byte(a))

	for i, v := range a {
		fmt.Printf("str[%2d] = % -12x = %[2]q \n", i, v)
	}

	fmt.Printf("%T %#[1]v %[2]T %#[2]v\n", a[:2], a[:1])

	r := []rune(a)
	fmt.Printf("Size of rune slice -> %d\n", int(unsafe.Sizeof(r[0]))*len(r))
	fmt.Printf("Size of byte slice -> %d\n", len(a))

	for n := 1055; n <= 1300; n++ {
		fmt.Printf("%c\t%[1]q\n", n)
	}

	bt := "ПРС"

	fmt.Printf("Size in bytes -> %d Size in Runes -> %d\n", len(bt), utf8.RuneCountInString(bt))

	rs := []rune(bt)
	sofrs := int(unsafe.Sizeof(rs[0])) * len(rs)

	fmt.Printf("Size in Runes -> %d Size of the Runes Slice -> %d\n", len(rs), sofrs)

	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	for i := 0; i < len(text); {
		runeInString, size := utf8.DecodeRuneInString(text[i:])
		i += size
		fmt.Printf("%c", runeInString)
	}
	fmt.Printf("\n%s\n", strings.Repeat(" -*- ", 25))
	for _, r := range text {
		fmt.Printf("%c", r)
	}

	word := []byte("düşünürler.")
	fmt.Printf("\n%s = % [1]x\n", word)

	_, sz := utf8.DecodeRune(word)

	nt := copy(word[:sz], bytes.ToUpper(word[:sz]))

	fmt.Printf("\n%s = % [1]x\n", word)
	fmt.Println(nt)
	print()
}
