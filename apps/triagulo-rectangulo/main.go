package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64
	const precision = 2

	//Entrada de datos
	fmt.Print("Ingrese lado 1:")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese lado 2:")
	fmt.Scanln(&lado2)

	// Proceos
	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	//Salida de datos
	fmt.Printf("\nÁrea: %.*f \n", precision, area)
	fmt.Printf("Perímetro: %.*f \n", precision, perimetro)
}
