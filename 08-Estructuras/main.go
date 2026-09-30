package main

import "fmt"

// Definimos una estructura para persona
type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	/*var p Persona
	p.nombre = "Jonathan"
	p.edad = 2
	p.correo = "stevdevelop@gmail.com"*/

	//Definimos una instacia el cual mostramos y modificamos
	p := Persona{"Jonathan", 29, "stevdevelop@gmail.com"}
	fmt.Println(p)
	fmt.Println(p.nombre)
	p.edad = 30
	fmt.Println(p)

	//Definimos otra instancia
	p2 := Persona{"Steven", 29, "Jonst970420@gmail.com"}
	fmt.Println(p2)
}
