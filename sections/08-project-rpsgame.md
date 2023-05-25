# Aplicación web sencilla

Código del archivo `main.go`

~~~go
package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Result struct {
	PlayerChoice   string
	ComputerChoice string
	Winner         string
}

func playGame(playerChoice string) Result {
	choices := []string{"piedra", "papel", "tijera"}
	computerChoice := choices[rand.Intn(len(choices))]

	playerChoice = strings.ToLower(playerChoice)

	switch playerChoice {
	case "piedra":
		switch computerChoice {
		case "piedra":
			return Result{playerChoice, computerChoice, "Empate"}
		case "papel":
			return Result{playerChoice, computerChoice, "La computadora gana"}
		case "tijera":
			return Result{playerChoice, computerChoice, "¡Tú ganas!"}
		}
	case "papel":
		switch computerChoice {
		case "piedra":
			return Result{playerChoice, computerChoice, "¡Tú ganas!"}
		case "papel":
			return Result{playerChoice, computerChoice, "Empate"}
		case "tijera":
			return Result{playerChoice, computerChoice, "La computadora gana"}
		}
	case "tijera":
		switch computerChoice {
		case "piedra":
			return Result{playerChoice, computerChoice, "La computadora gana"}
		case "papel":
			return Result{playerChoice, computerChoice, "¡Tú ganas!"}
		case "tijera":
			return Result{playerChoice, computerChoice, "Empate"}
		}
	}

	return Result{}
}

func play(w http.ResponseWriter, r *http.Request) {
	playerChoice := r.FormValue("choice")

	result := playGame(playerChoice)

	tmpl := template.Must(template.ParseFiles("game.html"))
	err := tmpl.Execute(w, result)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/play", play)

	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
~~~


Este es un código en Go (Golang) que implementa un juego de piedra, papel y tijera en un servidor web. A continuación, te explico las diferentes partes del código:

El primer bloque de importaciones es donde se importan las librerías necesarias para el funcionamiento del programa. Estas librerías incluyen funcionalidades para imprimir en la consola (fmt), trabajar con plantillas HTML (html/template), manejar solicitudes HTTP (net/http), realizar registros (log), generar números aleatorios (math/rand), trabajar con cadenas de texto (strings), y manejar el tiempo (time).

Luego se define una estructura llamada Result que representa el resultado de cada juego. Esta estructura tiene tres campos: PlayerChoice (elección del jugador), ComputerChoice (elección de la computadora) y Winner (ganador del juego).

La función playGame es la encargada de ejecutar un juego de piedra, papel y tijera. Recibe la elección del jugador como argumento y devuelve un resultado de tipo Result. La función selecciona aleatoriamente la elección de la computadora y luego realiza una serie de comparaciones utilizando una declaración switch para determinar el ganador del juego.

A continuación, se define la función play que maneja la solicitud HTTP para jugar al juego. Obtiene la elección del jugador a través de la solicitud y luego llama a la función playGame para obtener el resultado. Luego, utiliza una plantilla HTML llamada "game.html" para renderizar el resultado y enviarlo al cliente.

En la función main, se inicializa el generador de números aleatorios utilizando la función rand.Seed con el tiempo actual en nanosegundos. Luego, se configuran los manejadores de rutas HTTP utilizando http.HandleFunc. La ruta raíz ("/") muestra el archivo "index.html" y la ruta "/play" llama a la función play para procesar el juego. Finalmente, el servidor se inicia utilizando http.ListenAndServe en el puerto 8080.

Por último, se imprime un mensaje en la consola indicando que el servidor está en ejecución.

Código de ´index.html´

~~~html
<!DOCTYPE html>
<html>
<head>
    <title>Juego de Piedra, Papel o Tijera</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0/dist/css/bootstrap.min.css">
</head>
<body>
    <div class="container">
        <h1 class="mt-5">Juego de Piedra, Papel o Tijera</h1>
        <form action="/play" method="post">
            <div class="mb-3">
                <label for="choice" class="form-label">Elige una opción:</label>
                <select name="choice" id="choice" class="form-select">
                    <option value="piedra">Piedra</option>
                    <option value="papel">Papel</option>
                    <option value="tijera">Tijera</option>
                </select>
            </div>
            <button type="submit" class="btn btn-primary">Jugar</button>
        </form>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
~~~

Código de game.html´

~~~html
<!DOCTYPE html>
<html>
<head>
    <title>Resultado del juego</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0/dist/css/bootstrap.min.css">
</head>
<body>
    <div class="container">
        <h1 class="mt-5">Resultado del juego</h1>
        <p>Tu elección: {{.PlayerChoice}}</p>
        <p>Elección de la computadora: {{.ComputerChoice}}</p>
        <h2>{{.Winner}}</h2>
        <a href="/" class="btn btn-primary">Volver a jugar</a>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
~~~