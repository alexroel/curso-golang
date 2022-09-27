package main

import "fmt"

//Variables a nivel del paquete
var a, b = 10, 20

func main() {

	//Variable a nivel de función
	var name = "Alex Roel"

	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor de b: ", b)
	fmt.Println("Nombre: ", name)
}
