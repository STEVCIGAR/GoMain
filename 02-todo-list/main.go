package main

import "fmt"

func main() {

	//Declaración de variables
	/*var firstName, lastName string
	var age int

	var (
		firstName = "Steven"
		lastName  = "Jaramillo"
		age       = 29
	)*/

	//var firstName, lastName, age = "Steven", "Jaramillo", "29"
	firstName, lastName, age := "Steven", "Jaramillo", "29"

	fmt.Println("Nombre: ", firstName, lastName, "Edad: ", age)
}
