package main

import (
	"fmt"
	"log"

	"github.com/STEVCIGAR/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Jonathan", "Ana", "Michael"} //hace que cualquier mensaje que se imprima con log comience con ese texto.
	msgs, err2 := greetings.Hellos(names)           //elimina la marca de tiempo y otra información adicional que por defecto muestra log.
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(msgs)

	fmt.Println("*****************************************")
	message, err := greetings.Hello("Jonathan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
