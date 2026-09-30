package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hola Mundo")
	fmt.Println("Hola Mundo")

	// name := "Stev"
	// age := 29

	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scan(&name)

	fmt.Print("Ingrese su edad: ")
	fmt.Scan(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años  \n", name, age)

	greting := fmt.Sprintf("Hola, me llamo %s y tengo %d años", name, age)

	fmt.Println(greting)

	fmt.Printf("El tipo de name es: %T\n", name)
	fmt.Printf("El tipo de age es: %T\n", age)
}
