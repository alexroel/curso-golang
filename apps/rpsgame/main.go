package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
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

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/play", play)

	fmt.Println("Servidor en ejecución en http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
