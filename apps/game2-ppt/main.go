package main

import "fmt"

func main() {
	// fmt.Println("Juego de PIEDRA, PAPEL o TIJERA")

	// // 0 -> Piedra -->
	// // 1 -> Papel --> 0
	// // 2 -> Tigera

	// rand.Seed(time.Now().UnixNano())
	// computerValue := rand.Intn(3)
	// values := [3]string{"🪨Piedra", "📰Papel", "✂️Tijera"}

	// //Datos de usuario
	// var userValue int
	// fmt.Print("Ingres una opción: ")
	// fmt.Scanln(&userValue)

	// fmt.Println(values[computerValue])

	// //piedra gana a tijera
	fmt.Println("La suma es: ", sumar(4, 4.5))

}

func sumar(a int, b float32) float32 {
	n1 := float32(a) //Convertir a flotante

	return n1 + b
}
