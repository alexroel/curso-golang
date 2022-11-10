package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("CALCULADORA BÁSICA")

	// Entrada de datos
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("► ")
	scanner.Scan() //Escanear

	//Convertir a texto
	expression := scanner.Text()
	fmt.Println(expression)

	values := strings.Split(expression, "+")
	var suma int

	for i := 0; i < len(values); i++ {
		value, _ := strconv.Atoi(values[i])
		suma += value
	}

	fmt.Println("► ", suma)
}
