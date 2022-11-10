package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		fmt.Println("--------------------------------\n",
			"     🎮 Juego de Adivinanza 🎮\n",
			"📘 Adivina un número aleatorio entre 1 y 100\n",
			"--------------------------------\n",
			"1- Nivel Facil \n",
			"2- Nivel Intermedio \n",
			"3- Nivel Dificil \n",
			"4- Salr del Juego")

		var option int
		fmt.Print("👉 Ingrese una Opción: ")
		fmt.Scanln(&option)

		if option == 1 {
			play(10)
		} else if option == 2 {
			play(7)
		} else if option == 3 {
			play(5)
		} else if option == 4 {
			fmt.Println("🙋‍♂️ Cerrando el juego")
			break
		} else {
			fmt.Println("🙅‍♂️ Opción incorrecta")
		}
	}
}

// Dibuja corazones
func drawHearts(lives int) string {
	var hearts string

	for i := 0; i < lives; i++ {
		hearts += "💚"
	}

	return hearts
}

// Función Jugar
func play(lives int) {
	//Genera números aleatorios
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	var chosenNumber int

	for chosenNumber != randomNumber {

		fmt.Println(" --------------------------\n",
			"Vidas:", drawHearts(lives), "\n",
			"--------------------------")

		fmt.Print("👉 Ingrese un número: ")
		fmt.Scanln(&chosenNumber)

		if randomNumber < chosenNumber {
			fmt.Println("😎 El número es mas pequeño")
			lives--
		} else if randomNumber > chosenNumber {
			fmt.Println("😎 El número es mas grande")
			lives--
		}

		if lives == 0 {
			fmt.Println(" --------------------------\n",
				"    💀 GAME OVER 💀     \n",
				"--------------------------")
			fmt.Printf("😎 El numero era %d\n", randomNumber)
			break
		}
	}

	if chosenNumber == randomNumber {
		fmt.Println(" --------------------------\n",
			"    😀 Ganaste el juego 😀     \n",
			"--------------------------")
	}
}
