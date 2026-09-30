package main

import "fmt"

//const pi float32 = 3.14159265
const Pi = 3.14159265

const (
	X = 100
	Y = 0b1010 // binario
	Z = 0o12   // octal
	W = 0xFF   // hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println(Pi)
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
	fmt.Println(W)
	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)
}
