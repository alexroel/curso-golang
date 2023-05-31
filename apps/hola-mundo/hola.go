package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())
	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d añon.\n", name, age)

	fmt.Printf("El tipo de name es: %T\n", name)
	fmt.Printf("El tipo de age es: %T\n", age)

}
