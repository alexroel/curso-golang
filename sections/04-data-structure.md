# Estructura de datos

1. [Introducción](#introducción)
2. [Matrices](#matrices)
3. [Rebanadas](#rebanadas)
4. [Mapas](#mapas)
5. [Estructura](#estructura)
6. [Punteros y métodos](#punteros-y-métodos)
7. [Intefaces](#intefaces)
8. [Proyecto de la sección](#proyecto-de-la-sección)
9. [Resumen](#resumen)

---
## Introducción 

¡Bienvenidos a la sección de Estructura de Datos en el curso de Go! En esta sección, exploraremos una variedad de herramientas y técnicas para manejar datos de manera eficiente y efectiva en programas Go.

Comenzaremos examinando las matrices, que son colecciones de elementos del mismo tipo que se almacenan en una estructura de datos rectangular. A continuación, estudiaremos los segmentos, que son subconjuntos de una matriz y se utilizan a menudo para manipular partes específicas de una estructura de datos.

Continuaremos con los mapas, que son estructuras de datos más complejas que permiten almacenar y recuperar valores asociados a claves específicas. A continuación, examinaremos las estructuras, que son una forma de agrupar diferentes tipos de datos en una sola unidad lógica.

Además, exploraremos los punteros y los métodos, que son herramientas poderosas para trabajar con datos en Go. Los punteros permiten acceder y manipular directamente la memoria de una computadora, mientras que los métodos son funciones asociadas a un tipo de datos específico que pueden realizar operaciones sobre ese tipo de datos.

Por último, estudiaremos las interfaces, que son una forma de definir comportamientos que deben cumplir diferentes tipos de datos en un programa. Las interfaces permiten escribir código que es más genérico y reutilizable, lo que puede simplificar significativamente el proceso de desarrollo de software.

¡Prepárense para sumergirse en el emocionante mundo de la estructura de datos en Go!

---
## Matrices 
Las matrices de Go son una estructura de datos que se utiliza para almacenar un conjunto de elementos del mismo tipo. Se definen con una longitud fija y pueden ser de cualquier tipo de datos, incluidos tipos definidos por el usuario.

### Declaración de matrices
En Go, se pueden declarar matrices utilizando la siguiente sintaxis:

~~~go
var matriz [numeroElementos]tipoDato
~~~
Donde nombre es el nombre del array, longitud es el número de elementos que contendrá el array y tipo es el tipo de datos que contendrán los elementos del array.

Por ejemplo, para crear un array de enteros con 5 elementos, se puede escribir:

~~~go
var a [5]int
~~~
Para asignar valores a los elementos del array, se puede utilizar la siguiente sintaxis:

~~~go
a[0] = 10
a[1] = 20
a[2] = 30
a[3] = 40
a[4] = 50
~~~

### Inicialización de matrices
También es posible inicializar un array con valores en la misma línea en la que se define:

~~~go
var a = [5]int{10, 20, 30, 40, 50}
~~~


### Puntos suspensivos en matrices
Otra forma de declarar e inicializar una matriz cuando no se sabe cuántas posiciones se van a necesitar, pero sí cuál es el conjunto de elementos de datos, es usar puntos suspensivos `(...)`, como en este ejemplo:
~~~go
a := [...]int{10, 20, 30, 40, 50}
~~~

### Acceder a los elementos 
Para acceder a los elementos de un array, se utiliza la siguiente sintaxis:
~~~go	
valor := matriz[indice]
~~~

### Iterando matriz con for
Los arrays también se pueden utilizar en combinación con bucles para realizar operaciones en todos sus elementos. Por ejemplo:

~~~go
for i := 0; i < len(miArray); i++ {
    fmt.Println(miArray[i])
}
~~~

Este bucle imprimirá todos los elementos del array en la consola.

### Matrices multidimensionales
Go admite las matrices multidimensionales cuando se necesita trabajar con estructuras de datos complejas. Vamos a crear un programa en el que se declare e inicialice una matriz bidimensional. Use el código siguiente:

En Go, se pueden declarar matrices utilizando la siguiente sintaxis:
~~~go
var matriz [numeroFilas][numeroColumnas]tipoDato
~~~
Por ejemplo, para declarar una matriz de enteros de 3x5, se puede escribir:
~~~go 
package main

import "fmt"

func main() {
    var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = (i + 1) * (j + 1)
        }
        fmt.Println("Row", i, twoD[i])
    }
    fmt.Println("\nAll at once:", twoD)
}
~~~

---
## Rebanadas
En Go, una "rebanada" o "slice" es un tipo de datos que representa una sección contigua de un arreglo. Las rebanadas proporcionan una forma conveniente y flexible de trabajar con partes de un arreglo, ya que permiten acceder a elementos específicos del arreglo y también permiten agregar o eliminar elementos de manera dinámica.

### Declaración e inicialización de una rebanada 
La declaración de un segmento se hace lo mismo que la de una matriz. Por ejemplo, el código siguiente representa lo que ha visto en la imagen del segmento:

~~~go
var nombre []tipo
~~~

Por ejemplo, para crear un slice de enteros, se puede escribir:

~~~go
var miRebanada []int
//Declaraciín e incializacion de una rebanada
diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

// Crear una rebanada apartir de otra
diasRebanada := diasSemana[0:5]
diasRebanada:= diasSemana[0:5]

//Motrar tambien la longitud y la capacidad
fmt.Println(diasRebanada)
fmt.Println("Longitud: ", len(diasRebanada))
fmt.Println("Capacidad: ", cap(diasRebanada))
~~~

### Anexión de elementos
En Go, ``append()`` es una función predefinida que se utiliza para agregar uno o más elementos al final de una rebanada. La función toma como argumentos una o más rebanadas seguidas de los elementos que se desean agregar.

La función `append()` devuelve una nueva rebanada que contiene todos los elementos de la rebanada original, así como los nuevos elementos agregados al final. Si la capacidad de la rebanada original no es suficiente para almacenar todos los elementos, `append()` creará una nueva rebanada con capacidad suficiente y copiará todos los elementos de la rebanada original a la nueva rebanada antes de agregar los nuevos elementos.

Para asignar valores a un slice, se utiliza la siguiente sintaxis:

~~~go
	// Agregar datos
	diasRebanada = append(diasRebanada, "Jueves")
	diasRebanada = append(diasRebanada, "Viernes")
	fmt.Println(diasRebanada)
~~~

### Quitar elementos
En Go, se puede quitar elementos de una rebanada utilizando el operador de corte `([:])` y la función ``append()``.

El operador de corte se utiliza para seleccionar una porción de una rebanada. La sintaxis general es `rebanada[inicio:fin]`, donde inicio es el índice del primer elemento que se desea incluir y fin es el índice del primer elemento que se desea excluir. Si se omite inicio, se utilizará el valor 0 como valor predeterminado. Si se omite fin, se utilizará el valor `len(rebanada)` como valor predeterminado.

Para quitar elementos de una rebanada, se puede utilizar el operador de corte para seleccionar los elementos que se desean mantener y luego asignar el resultado a la rebanada original. Alternativamente, se puede utilizar la función `append()` para crear una nueva rebanada que contenga solo los elementos que se desean mantener.

~~~go
	// Quitar elementos
	diasRebanada = append(diasRebanada[:2], diasRebanada[3:]...)
	fmt.Println(diasRebanada)
~~~

### Funciones Make y Copy 

En Go, las funciones `make()` y `copy()` se utilizan para trabajar con rebanadas.

La función `make()` se utiliza para crear una nueva rebanada con una longitud y capacidad especificadas. La sintaxis de `make()` es la siguiente: `make([]Tipo, longitud, capacidad)`, donde Tipo es el tipo de elementos que se almacenarán en la rebanada, longitud es el número de elementos que inicialmente tendrá la rebanada, y capacidad es el número máximo de elementos que la rebanada puede contener. La capacidad puede ser mayor o igual a la longitud. Si se omite la capacidad, se utilizará la longitud como valor predeterminado.

Aquí hay un ejemplo que muestra cómo utilizar la función `make()` para crear una nueva rebanada:

~~~go
	// Crear una rebanada con make
	nombres := make([]string, 5, 10)
	fmt.Println(nombres)
	fmt.Println("Longitud: ", len(nombres))
	fmt.Println("Capacidad: ", cap(nombres))
~~~

La función `copy()` se utiliza para copiar elementos de una rebanada a otra. La sintaxis de c`opy()` es la siguiente: `copy(destino []Tipo, origen []Tipo)`, donde destino es la rebanada de destino y origen es la rebanada de origen. Los elementos de origen se copiarán en destino. La función devuelve el número de elementos que se copiaron.

~~~go
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	rebanadas := copy(rebanada1, rebanada2)
	fmt.Println(rebanadas)
~~~

---
## Mapas
En Go, los mapas son una estructura de datos que permiten almacenar y recuperar pares clave-valor de manera eficiente. Los mapas se definen utilizando la siguiente sintaxis:

~~~go
map[<tipo de clave>]<tipo de valor>
~~~

Donde `<tipo de clave>` es el tipo de datos que se utilizará como clave en el mapa, y `<tipo de valor>` es el tipo de datos que se almacenará como valor en el mapa.

Por ejemplo, para crear un mapa que almacene nombres de colores y sus códigos hexadecimales correspondientes, se podría escribir:

~~~go
	//Declaración e inicialización de una mapa
	colors := map[string]string{
		"rojo":   "#FF0000",
		"verder": "#00FF00",
		"azul":   "#0000FF",
	}

	fmt.Println(colors["verde"])
~~~
Luego, se pueden agregar pares clave-valor al mapa utilizando la siguiente sintaxis:

~~~go
	//Agregar nuevo elemento
	colors["negro"] = "#000000"
	fmt.Println(colors)
~~~
Para recuperar el valor correspondiente a una clave específica en el mapa, se puede utilizar la siguiente sintaxis:

~~~go
valor := miMap[clave]
~~~
Por ejemplo, para obtener el código hexadecimal correspondiente al color "rojo" del mapa colores, se puede escribir:

~~~go
fmt.Println(colors["rojo"])
~~~
También es posible verificar si una clave específica existe en un mapa utilizando la siguiente sintaxis:

~~~go
//valor, ok := miMap[clave]

valor, ok := colors["rojo"]
if ok {
    fmt.Println("Rojo en hexadecimal es: ", valor)
} else {
    fmt.Println("No existe este clave")
}
~~~
Eliminar un elemento de la siguiente forma:

~~~go
delete(colors, "negro")
fmt.Println(colors)
~~~

---
## Estructura
En Go, un struct es una estructura de datos que permite definir un tipo de datos personalizado compuesto por diferentes campos con nombres y tipos específicos. Los structs se definen utilizando la siguiente sintaxis:

~~~go
type <nombre del struct> struct {
    <nombre del campo> <tipo de dato>
    <nombre del campo> <tipo de dato>
    ...
}
~~~

Por ejemplo, para crear un struct que represente a una persona con nombre, edad y correo electrónico, se podría escribir:

~~~go
type Persona struct {
    Nombre string
    Edad int
    CorreoElectronico string
}
~~~
Luego, se pueden crear instancias de este struct y asignar valores a sus campos utilizando la siguiente sintaxis:

~~~go
var p Persona
p.Nombre = "Juan"
p.Edad = 30
p.CorreoElectronico = "juan@example.com"
~~~
También es posible crear e inicializar una instancia de un struct en una sola línea utilizando la siguiente sintaxis:

~~~go
p := Persona{"Juan", 30, "juan@example.com"}
~~~
Para acceder a los campos de una instancia de un struct, se puede utilizar la siguiente sintaxis:

~~~go
valor := instancia.<nombre del campo>
~~~
Por ejemplo, para obtener el nombre de la persona p definida anteriormente, se puede escribir:

~~~go
nombre := p.Nombre
~~~

Los structs son una estructura de datos muy útil en Go, ya que permiten definir tipos de datos personalizados compuestos por diferentes campos. Los structs se pueden utilizar en una amplia variedad de aplicaciones, como la representación de objetos en un juego, el almacenamiento de datos en una base de datos o la definición de mensajes en una aplicación de red.

---
## Punteros y métodos
En Go, se pueden definir métodos en estructuras para realizar operaciones en las variables de la estructura. Los métodos se definen como funciones que tienen un receptor, que es un puntero o una variable de la estructura. Los punteros son comúnmente utilizados como receptores de método en Go, ya que permiten que los métodos modifiquen directamente la variable original.

### Punteros 
En Go, un puntero es una variable que almacena la dirección de memoria de otra variable. Los punteros se utilizan para referenciar o acceder a la variable original a través de su dirección de memoria.

Para declarar un puntero en Go, se utiliza el operador & seguido del nombre de la variable a la que se desea obtener la dirección de memoria. Por ejemplo, si se tiene una variable x de tipo int, se puede obtener un puntero a x de la siguiente manera:

~~~go
var x int = 10
var p *int = &x
~~~

En este ejemplo, la variable p es un puntero a x. El tipo de p es *int, que se lee como "un puntero a un valor de tipo int".

Para acceder al valor almacenado en una variable a través de un puntero, se utiliza el operador *. Por ejemplo, si se tiene un puntero p a una variable x, se puede obtener el valor de x de la siguiente manera:

~~~go
var x int = 10
var p *int = &x
var y int = *p
~~~

En este ejemplo, la variable y toma el valor de x a través del puntero p.

Los punteros son muy útiles en Go para evitar la copia de datos al pasar argumentos a funciones y para permitir que las funciones modifiquen variables fuera de su alcance. Además, los punteros son necesarios para trabajar con estructuras de datos complejas como slices y maps. Sin embargo, también pueden ser fuente de errores si se usan incorrectamente, por lo que es importante tener cuidado al trabajar con punteros.

### Métodos
Para definir un método en Go, se sigue la siguiente sintaxis:

~~~go
func (p *Persona) DiHola() {
    fmt.Println("Hola, mi nombre es", p.Nombre)
}
~~~

En este ejemplo, se define un método llamado SayHello() en la estructura Person. El receptor del método es un puntero a Person, que se indica colocando el tipo *Person antes del nombre del método entre paréntesis. Dentro del cuerpo del método, se puede acceder a las variables de la estructura utilizando el operador de indirección *.

Para llamar a un método en una estructura, se utiliza la siguiente sintaxis:

~~~go
var p Persona
p.Nombre = "John"
p.DiHola()
~~~
En este ejemplo, se crea una variable p de tipo Person y se establece su nombre en "John". Luego, se llama al método SayHello() en p, que imprimirá "Hello, my name is John".

Los punteros son particularmente útiles al trabajar con métodos porque permiten modificar directamente la variable original en lugar de hacer una copia. Por ejemplo, si se tiene un método que agrega un elemento a una slice, la slice original se modificará si se usa un puntero como receptor, pero no si se usa una variable normal como receptor.

En resumen, los punteros son esenciales en Go para trabajar con estructuras de datos complejas y para permitir que los métodos modifiquen directamente la variable original. Los métodos con punteros como receptores pueden ser particularmente útiles para modificar variables en lugar de hacer copias y para trabajar con estructuras de datos complejas como slices y maps.

---

---
## Proyecto de la sección
Proyecto gestor de tareas: Crea un programa que permita gestionar una lista de tareas pendientes. Utiliza una estructura para representar cada tarea (título, descripción, completado.) y un slice para almacenar todas las tareas. Utiliza métodos para agregar, eliminar y actualizar tareas, y para marcar una tarea como completada. Este proyecto se tiene que hacer con entrada de datos por la consola. 

1. Crea una estructura para representar cada tarea, con campos para el título, la descripción y el estado de completado.
2. Crea un slice para almacenar todas las tareas.
3. Implementa un método para agregar una tarea al slice.
4. Implementa un método para eliminar una tarea del slice.
5. Implementa un método para actualizar una tarea en el slice.
6. Implementa un método para marcar una tarea como completada.
7. Implementa una función para mostrar la lista de tareas al usuario.
8. Implementa una función para recibir la entrada del usuario por la consola, que permita agregar, eliminar, actualizar y marcar tareas como completadas.
9. Implementa un loop que solicite continuamente la entrada del usuario y muestre la lista de tareas actualizada después de cada operación de gestión de tareas.

Código del proyecto
~~~go
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Estructura de Tarea
type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

// Estructura para lista de teareas
type ListaTareas struct {
	tareas []Tarea
}

// Método para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Método para marcar como completado
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// Método para editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// Métdodo para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//Instancia de Lista de tareas
	lista := ListaTareas{}

	// Instancia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese el descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el índice de la tarea que desea actualizar:")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese el descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción inválida.")
		}

		// Listar todas las tareas
		fmt.Println("Lista de tareas:")
		fmt.Println("===================================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.descripcion, t.completado)
		}
		fmt.Println("===================================================")
	}
}

~~~ 

---
## Resumen 
En esta sección del curso de Go, se cubrieron varios temas importantes de estructuras de datos y su implementación en el lenguaje de programación Go. Se explicó cómo declarar y utilizar matrices en Go, incluyendo la inicialización de matrices y la indexación de elementos en matrices.

También se abordó el tema de las rebanadas, que son una forma flexible y poderosa de trabajar con secciones de matrices y otras estructuras de datos en Go. Se mostró cómo declarar y utilizar rebanadas, así como cómo modificarlas y crear nuevas rebanadas a partir de rebanadas existentes.

Además, se cubrió la estructura de datos de los mapas en Go, que permiten almacenar pares clave-valor y proporcionan una forma rápida y eficiente de buscar valores por su clave. Se explicó cómo declarar y utilizar mapas en Go, así como cómo agregar, acceder y eliminar elementos de los mapas.

Otro tema importante fue la estructura de datos de las estructuras en Go, que permiten agrupar diferentes tipos de datos en una sola entidad. Se mostró cómo declarar y utilizar estructuras en Go, así como cómo acceder y modificar sus campos.

También se cubrió el tema de los punteros y métodos en Go, que permiten trabajar con referencias a memoria y definir funciones que operan en una estructura en particular. Se explicó cómo declarar y utilizar punteros en Go, así como cómo definir y utilizar métodos en Go.

Por último, se abordó el tema de las interfaces en Go, que permiten definir un conjunto de métodos que deben ser implementados por una estructura en particular. Se mostró cómo declarar e implementar interfaces en Go, así como cómo utilizar interfaces para escribir código más genérico y reutilizable.

En el proyecto de la sección, se desarrolló una lista de tareas en Go utilizando varios de los conceptos y estructuras de datos cubiertos en la sección, lo que permitió poner en práctica los conocimientos adquiridos y aplicarlos a un proyecto práctico.
