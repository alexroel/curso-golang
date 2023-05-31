package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()

}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("¡Felicitaciones, adivinaste el número!")
			jugarNuvamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El número a adivinar es mayor.")
		} else if numIngresado > numAleatorio {
			fmt.Println("El número a adivinar es menor.")
		}
	}

	fmt.Println("Se acabaron los intentos. El número era:", numAleatorio)
	jugarNuvamente()
}

func jugarNuvamente() {
	var eleccion string
	fmt.Print("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Elección inválida. Inténtalo nuevamente.")
		jugarNuvamente()
	}
}
