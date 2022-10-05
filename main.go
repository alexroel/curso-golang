package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Generar numero aleatorio
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101)

	fmt.Println(randomNumber)
}
