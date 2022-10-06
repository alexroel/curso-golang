package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Dibuja corazones
func drawHearts(lives int) string {
	var hearts string

	for i := 0; i < lives; i++ {
		hearts += "♥ "
	}
	return hearts
}

// Función jugar
func play(lives int) {
	//Generar numero aleatorio
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101)
	var chosenNumber int

	//Compara si el número ingresado por el jugador es igual al aleatorio
	for chosenNumber != randomNumber {

		//Mostrar vidas
		fmt.Println("-------------------------------")
		fmt.Println("Vidas: ", drawHearts(lives))
		fmt.Println("-------------------------------")

		//Optener número ingresado por el jugador
		fmt.Print("Ingrese un número entre 0 y 101: ")
		fmt.Scanln(&chosenNumber)

		//Encontrando el número
		if randomNumber < chosenNumber {
			fmt.Println("► El número es mas pequeño")
			lives-- //Quita una vida por intento
		} else if randomNumber > chosenNumber {
			fmt.Println("► El número es mas grande")
			lives-- //Quita una vida por intento
		}

		if lives == 0 {
			fmt.Println("-------------------------------")
			fmt.Println("|        ☺ GAME OVER ☺         |")
			fmt.Println("-------------------------------")
			break
		}

	}

	// Comprobando si ganaste el juego
	if chosenNumber == randomNumber {
		fmt.Println("-------------------------------")
		fmt.Println("|     ☻ Ganaste el juego ☻    |")
		fmt.Println("-------------------------------")
	}
}

func main() {
	for {
		fmt.Println("--------------------------------\n",
			"     Juego adivina un número\n",
			"--------------------------------\n",
			"1- Nivel Facil \n",
			"2- Nivel Intermedio \n",
			"3- Nivel Dificil\n",
			"4- Salir del juego")

		fmt.Print("► Ingrese una Opción: ")
		var option int
		fmt.Scanln(&option)

		if option == 1 {
			play(10)
		} else if option == 2 {
			play(7)
		} else if option == 3 {
			play(5)
		} else if option == 4 {
			fmt.Println("► Cerrando juego")
			break
		} else {
			fmt.Println("► Opción incorrecta")
		}
	}
}
