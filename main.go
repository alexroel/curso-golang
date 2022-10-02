package main

import "fmt"

func main() {
	a, b := 10, 5

	result := (a*b-2*b) >= 20 && !((a % b) != 0)

	fmt.Println("Resultado: ", result)
}
