package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(integer16 + int16(integer32))
	fmt.Println(int32(integer16) + integer32)

	s := "10"
	i, _ := strconv.Atoi(s)
	fmt.Println(i + i)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s)
}
