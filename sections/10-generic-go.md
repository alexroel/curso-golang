# Genericos en go 

1. [Introducción](#introducción)
2. [Uso de any](#uso-de-any)
3. [Parametros de tipos y restricciones](#parametros-de-tipos-y-restricciones)
4. [Crear restricciones](#crear-restricciones)
5. [Restricciones y operadores](#restricciones-y-operadores)
6. [Estructuras genéricas](#estructuras-genéricas)
7. [Resumen](#resumen)
---
## Introducción 
¡Bienvenido al curso de Go sobre Generics!

En esta sección del curso, exploraremos uno de los aspectos más emocionantes y esperados de Go: los Generics. Los Generics permiten la creación de funciones y estructuras que pueden trabajar con diferentes tipos de datos de manera segura y eficiente, sin tener que duplicar código.

Comenzaremos introduciendo el concepto de los parámetros de tipos y las restricciones. Los parámetros de tipos son marcadores o nombres de tipo que nos permiten crear funciones y estructuras genéricas. Las restricciones, por otro lado, son condiciones que podemos establecer en esos parámetros de tipos para garantizar ciertas propiedades o comportamientos de los tipos utilizados.

A continuación, profundizaremos en cómo crear restricciones en funciones genéricas. Veremos ejemplos prácticos de cómo especificar y limitar las propiedades de los tipos de datos utilizados, lo que nos permitirá escribir código más robusto y seguro.

Además, exploraremos cómo los operadores y funciones pueden desempeñar un papel importante al trabajar con tipos genéricos. Veremos cómo utilizar operadores aritméticos, operadores de comparación y operadores de igualdad con los tipos genéricos, y cómo esto puede facilitar las operaciones comunes en nuestra programación genérica.

Luego, abordaremos el tema de las estructuras genéricas. Aunque actualmente Go no admite estructuras genéricas en su biblioteca estándar, discutiremos algunas alternativas y patrones que se pueden utilizar para lograr un comportamiento similar. Esto nos permitirá aprovechar al máximo las capacidades existentes de Go mientras esperamos la implementación oficial de las estructuras genéricas en futuras versiones.

En resumen, en esta sección del curso, exploraremos los fundamentos de los Generics en Go. Aprenderemos sobre los parámetros de tipos, las restricciones y cómo crear funciones genéricas sólidas. También examinaremos los operadores y funciones que se pueden utilizar con tipos genéricos. Aunque Go no admite estructuras genéricas en este momento, exploraremos algunas alternativas prácticas.

¡Estás a punto de descubrir una poderosa herramienta que mejorará tu capacidad para escribir código flexible, reutilizable y seguro en Go! ¡Vamos a sumergirnos en el mundo de los Generics en Go y aprovechar al máximo esta emocionante funcionalidad!

---
## Uso de any 
Una función genérica es una función que puede ser utilizada con diferentes tipos de datos sin necesidad de escribir código específico para cada tipo. En otros términos, permite la reutilización de código al proporcionar una abstracción sobre los tipos de datos específicos.

La programación genérica es una característica común en muchos lenguajes de programación modernos y permite escribir algoritmos y estructuras de datos que funcionan con cualquier tipo de dato compatible.

Las funciones genéricas son especialmente útiles cuando se necesita realizar operaciones similares en diferentes tipos de datos. En lugar de escribir varias funciones casi idénticas para cada tipo, una función genérica puede ser escrita una sola vez y luego utilizada con diferentes tipos de datos.

La función `PrintList` que has proporcionado es un ejemplo de una función variádica en Go. La sintaxis `...any` indica que la función puede aceptar un número variable de argumentos de cualquier tipo. 

~~~go
func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func main() {
	PrintList("Alex", "Roel", "Juan", "Pedro")
	PrintList(100, 456, 789, 456, 452)
	PrintList("Hola", 452, 4.25, true)
}
~~~

En este ejemplo, se llama a `PrintList` con diferentes argumentos de diferentes tipos. La función iterará sobre cada uno de estos argumentos y los imprimirá en la consola utilizando `fmt.Println`.

---
## Parametros de tipos y restricciones

Cuando se trabaja con funciones genéricas, los parámetros de tipos y las restricciones son elementos importantes que permiten definir y limitar los tipos de datos que pueden utilizarse en una función genérica. Estos elementos brindan flexibilidad y al mismo tiempo garantizan ciertas propiedades o comportamientos de los tipos utilizados.

Un ejemplo básico de una función que suma N números: 
~~~go
func Sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(Sum(100, 456, 789, 456, 452))
}
~~~

### Declaración de parámetros de tipos
Los parámetros de tipos son marcadores o nombres de tipo que se utilizan en la definición de una función genérica para representar un tipo desconocido. Estos parámetros se indican entre corchetes angulares ([]) y se colocan antes de los parámetros regulares de la función.

- Restricción arbitrario 
~~~go
func Sum[T int](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(Sum[int](100, 456, 789, 456, 452))
}
~~~

- Restricción de unión de elementos 

~~~go
func Sum[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(Sum(100, 456, 789, 456, 452))
	fmt.Println(Sum(4.5, 5.5, 7.5))
}
~~~

- Restricción de aproximación 

~~~go
type integer int

func Sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	var num1 integer = 100
	var num2 integer = 300

	fmt.Println(Sum(100, 456, 789, 456, 452))
	fmt.Println(Sum(4.5, 5.5, 7.5))
	fmt.Println(Sum(num1 + num2))
}
~~~


---
## Crear restricciones 
Al crear restricciones en funciones genéricas, puedes especificar los requisitos que deben cumplir los tipos de datos utilizados en esa función. Las restricciones permiten garantizar propiedades o características específicas de los tipos utilizados, lo que ayuda a evitar errores y asegurar un comportamiento correcto en tiempo de compilación.

~~~go
type integer int

type Integer interface {
	~int | ~float64 | ~float32
}

func Sum[T Integer](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	var num1 integer = 100
	var num2 integer = 300

	fmt.Println(Sum(100, 456, 789, 456, 452))
	fmt.Println(Sum[float32](4.5, 5.5, 7.5))
	fmt.Println(Sum(num1, num2))
}
~~~

Utilizando el paquete de restricciones `constraints`

~~~go
import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type integer int

func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	var num1 integer = 100
	var num2 integer = 300

	fmt.Println(Sum(100, 456, 789, 456, 452))
	fmt.Println(Sum[float32](4.5, 5.5, 7.5))
	fmt.Println(Sum(num1, num2))
}
~~~

---
## Restricciones y operadores 
En el contexto de la programación genérica, las restricciones y los operadores se utilizan para especificar y limitar las propiedades de los tipos de datos que se pueden utilizar en una función genérica. Estas restricciones y operadores permiten establecer requisitos sobre los tipos de datos y realizar operaciones específicas con ellos.

~~~go
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func main() {

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}

	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "f"))
	fmt.Println(Includes(numbers, 4))
	fmt.Println(Includes(numbers, 5))

}
~~~

Utilizando otros operadores de comparación 

~~~go
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

func main() {

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

}
~~~

---
## Estructuras genéricas 

Para base de datos relacionales y base de datos documentales 

~~~go
type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"sadad-ad5d4-asd", "Zapatos", 50}

	fmt.Println(product1, product2)

}
~~~

---
## Resumen

En esta sección del curso, exploramos los Generics en Go. Introdujimos el concepto de parámetros de tipos y restricciones, que nos permiten crear funciones y estructuras genéricas. Aprendimos cómo especificar restricciones en los parámetros de tipos para garantizar propiedades específicas de los tipos utilizados.

Además, analizamos el uso de operadores y funciones con tipos genéricos, como los operadores aritméticos, de comparación y de igualdad. Estos operadores nos ayudan a realizar operaciones comunes con los tipos genéricos.

También abordamos el tema de las estructuras genéricas en Go. Aunque Go no admite estructuras genéricas en la biblioteca estándar, exploramos alternativas y patrones que se pueden utilizar para lograr un comportamiento similar.

En resumen, esta sección nos brindó una introducción a los Generics en Go. Aprendimos sobre los parámetros de tipos, las restricciones y cómo crear funciones genéricas sólidas. También exploramos el uso de operadores y funciones con tipos genéricos, así como alternativas para estructuras genéricas en Go.

Con estos conocimientos, hemos adquirido una poderosa herramienta para escribir código más flexible, reutilizable y seguro en Go. Los Generics nos permiten trabajar con diferentes tipos de datos de manera eficiente, sin tener que duplicar código. ¡Esperamos que aproveches al máximo esta funcionalidad y sigas aprendiendo y experimentando con los Generics en Go! 






