# Comenzamdo a codificar

1. [Paquetes](#Paquetes)
2. [Variables](#Variables)
3. [Tipos Básicos](#Tipos-Básicos)
4. [Valores cero](#valores-cero)
7. [Funciones](#Funciones)

---
## Paquetes 
- Cada programa de Go se compone de paquetes.
- Los programas comienzan a ejecutarse en el paquete `main`.
- Por convención, el nombre del paquete es el mismo que el último elemento de la ruta de importación. Por ejemplo, el "math/rand"paquete comprende archivos que comienzan con la instrucción package rand.

~~~go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("¡Hola Mundo de Go!")

	fmt.Println("Valor de PI: ", math.Pi)
	fmt.Println("Valor de 4^2: ", math.Pow(4, 2))

	//Generar numero aleatorio
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(101))
}
~~~

---
## Variables
- Una variable almacena datos temporalmente.
- La instrucción `var` declara una lista de variables.
- Una declaración `var` puede estar a nivel de paquete o de función.

~~~go
package main

import "fmt"

//Variables a nivel del paquete
var a, b int

func main() {

	//Variable a nivel de función
	var name string

	//Asignado valor a las variables
	a, b = 10, 20
	name = "Alex Roel"

	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor de b: ", b)
	fmt.Println("Nombre: ", name)
}
~~~

### Variables inicializadas 
- Podemos inicilizar una variable 
- Si hay un inicializador, se puede omitir el tipo de dato.
- Dependiendo del dato que se asigne tomara el tipado. 

~~~go
//Variables a nivel del paquete
var a, b = 10, 20

func main() {

	//Variable a nivel de función
	var name = "Alex Roel"

	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor de b: ", b)
	fmt.Println("Nombre: ", name)
}
~~~
### Declaraciones de variables cortas
- Dentro de una función, la declaración `:=` de asignación corta se puede usar en lugar de una declaración `var` con tipo implícito.
- Fuera de una función, cada instrucción comienza con una palabra clave ( var, func, etc.) y, por lo tanto, la construcción ´:=´ no está disponible.

~~~go
//Variables a nivel del paquete
var a, b = 10, 20

func main() {

	//Variable a nivel de función
	name := "Alex Roel"

	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor de b: ", b)
	fmt.Println("Nombre: ", name)
}
~~~

---
## Tipos Básicos
- `bool` Para booleanos con `true` y `false`
- `string` Para cadena de caracteres 
- `int` Para números enteros 32-bit o 64-bits
	- int8 
	- int16
	- int32
	- int64 
- `uint` Para números enteros postivos 
	- uint8
	- uint16
	- uint32
	- uint64
- `byte` que alias para `uint8`
- `rune` que es alias int32
- `float32` y `float64` para números flotantes con punto decimal
- `complex64` y `complex128` para números complejos


~~~go
	option, name, a, b, c, d := true, "Alex", -27, 45, 4.5, cmplx.Sqrt(-5+12i)

	fmt.Printf("Tipo %T valor %v\n", option, option)
	fmt.Printf("Tipo %T valor %v\n", name, name)
	fmt.Printf("Tipo %T valor %v\n", a, a)
	fmt.Printf("Tipo %T valor %v\n", b, b)
	fmt.Printf("Tipo %T valor %v\n", c, c)
	fmt.Printf("Tipo %T valor %v\n", d, d)
~~~

---
## Valores cero
Las variables declaradas sin un valor inicial explícito reciben su valor cero .

~~~go
	var (
		option bool
		name   string
		a      int
		b      uint
		c      float32
	)
	fmt.Printf("%v %q %v %v %v\n", option, name, a, b, c)
~~~
---
## Funciones 
- Una función contiene un bloque de código que se ejecuta solo cuando llamas la función.
- Una funioón puede ejecutar el bloque de código que tenga y devolver un valor. 
- Una función puede tomar cero o más argumentos.
- Observe que el tipo viene después del nombre de la variable.

~~~go
//Creando mis funciónes
func hello(name string) {
	//fmt.Println("Hola desde la función")
	fmt.Println("Hola ", name)
}

func sum(a, b int) int {
	return a + b
}

func main() {
	//llmar una función
	hello("Alex")
	fmt.Println("La suma es: ", sum(40, 50))
}
~~~

