package main

import (
	"fmt"
	"math"
)

func main() {
	var integer int = 25
	var float float32 = 1.618
	var boolean bool = true
	var text string = "Hola"
	fmt.Println(integer)
	fmt.Println(float)
	fmt.Println(boolean)
	fmt.Println(text)

	fmt.Println(math.MinInt64, math.MaxInt64)

	fullname := "Jonathan Steven Jaramillo Gracia \t(alias \"Stev\")\n"
	fmt.Println(fullname)

	var a byte = 'a'
	fmt.Println(a)

	s := `Hola mundo`
	fmt.Println(s[0])

	var r rune = '♥'
	fmt.Println(r)
}
