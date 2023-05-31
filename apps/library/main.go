package main

import "library/animal"

func main() {
	// var myBook = book.NewBook("Pedro Páramo", "Juan Rulfo", 128)

	// var myTextBook = book.NewTextBook("Comunicación", "Jaime Gamarra", 261,
	// 	"Santillana SAC", "Secundaria")

	// //myBook.PrintInfo()
	// //myTextBook.PrintInfo()

	// book.Print(myBook)
	// book.Print(myTextBook)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
