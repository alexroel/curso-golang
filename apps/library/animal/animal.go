package animal

import "fmt"

type Animal interface {
	Sonido()
}

// Estructura de perro y sus métodos
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau guau")
}

// Estructura de gato y sus métodos
type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau")
}

// Función para imprimer sonido
func HacerSonido(animal Animal) {
	animal.Sonido()
}
