# Estructura de datos

1. [Arrays](#Arrays)
2. [Slices](#Slices)
3. [Maps](#Maps)
4. [Estructura](#Estructura)
5. [Punteros y métodos](#Punteros-y-métodos)
6. [Intefaces](#Intefaces)
7. [Proyecto de la sección](#Proyecto-de-la-sección)

---
## Arrays
En Go, los arrays son una estructura de datos que se utiliza para almacenar un conjunto de elementos del mismo tipo. Se definen con una longitud fija y pueden ser de cualquier tipo de datos, incluidos tipos definidos por el usuario.

La sintaxis para definir un array en Go es la siguiente:

~~~go
var nombre [longitud]tipo
~~~
Donde nombre es el nombre del array, longitud es el número de elementos que contendrá el array y tipo es el tipo de datos que contendrán los elementos del array.

Por ejemplo, para crear un array de enteros con 5 elementos, se puede escribir:

~~~go
var miArray [5]int
~~~
Para asignar valores a los elementos del array, se puede utilizar la siguiente sintaxis:

~~~go
miArray[0] = 10
miArray[1] = 20
miArray[2] = 30
miArray[3] = 40
miArray[4] = 50
~~~

También es posible inicializar un array con valores en la misma línea en la que se define:

~~~go
var miArray = [5]int{10, 20, 30, 40, 50}
~~~

Para acceder a los elementos de un array, se utiliza la siguiente sintaxis:

~~~	
valor := miArray[indice]
~~~

Donde indice es el número del elemento al que se quiere acceder. El primer elemento tiene un índice de 0.

Los arrays también se pueden utilizar en combinación con bucles para realizar operaciones en todos sus elementos. Por ejemplo:

~~~go
for i := 0; i < len(miArray); i++ {
    fmt.Println(miArray[i])
}
~~~

Este bucle imprimirá todos los elementos del array en la consola.

En resumen, los arrays en Go son una estructura de datos útil para almacenar conjuntos de elementos del mismo tipo con una longitud fija. Se pueden inicializar con valores y se accede a sus elementos utilizando índices.

---
## Slices
En Go, los slices son estructuras de datos que permiten almacenar y manipular secuencias de elementos de cualquier tipo de datos. A diferencia de los arrays, los slices no tienen una longitud fija y pueden crecer y reducir su tamaño dinámicamente durante la ejecución del programa.

La sintaxis para definir un slice es similar a la de un array, pero sin especificar la longitud:

~~~go
var nombre []tipo
~~~

Por ejemplo, para crear un slice de enteros, se puede escribir:

~~~go
var miSlice []int
~~~
Para asignar valores a un slice, se utiliza la siguiente sintaxis:

~~~go
miSlice = append(miSlice, 10)
miSlice = append(miSlice, 20)
miSlice = append(miSlice, 30)
~~~
La función `append()` se utiliza para agregar elementos al final del slice. También es posible inicializar un slice con valores en la misma línea en la que se define:

~~~go
var miSlice = []int{10, 20, 30}
~~~
Para acceder a los elementos de un slice, se utiliza la misma sintaxis que para los arrays:

~~~go
valor := miSlice[indice]
~~~

donde indice es el número del elemento al que se quiere acceder. Los slices también se pueden utilizar en combinación con bucles para realizar operaciones en todos sus elementos.

Los slices también tienen otras funciones útiles, como `len()` para obtener la longitud del slice y `cap()` para obtener la capacidad del slice, que es el número máximo de elementos que puede contener el slice sin necesidad de redimensionarse.

En resumen, los slices en Go son una estructura de datos poderosa y flexible que permite almacenar y manipular secuencias de elementos de cualquier tipo de datos con una longitud variable. Los slices se utilizan comúnmente en lugar de los arrays en Go debido a su capacidad para crecer y reducir su tamaño dinámicamente durante la ejecución del programa.

---
## Maps
En Go, los maps son una estructura de datos que permiten almacenar y recuperar pares clave-valor de manera eficiente. Los maps se definen utilizando la siguiente sintaxis:

~~~go
miMap := make(map[<tipo de clave>]<tipo de valor>)
~~~

Donde `<tipo de clave>` es el tipo de datos que se utilizará como clave en el mapa, y `<tipo de valor>` es el tipo de datos que se almacenará como valor en el mapa.

Por ejemplo, para crear un mapa que almacene nombres de colores y sus códigos hexadecimales correspondientes, se podría escribir:

~~~go
colores := make(map[string]string)
~~~
Luego, se pueden agregar pares clave-valor al mapa utilizando la siguiente sintaxis:

~~~go
colores["rojo"] = "#FF0000"
colores["verde"] = "#00FF00"
colores["azul"] = "#0000FF"
~~~
Para recuperar el valor correspondiente a una clave específica en el mapa, se puede utilizar la siguiente sintaxis:

~~~go
valor := miMap[clave]
~~~
Por ejemplo, para obtener el código hexadecimal correspondiente al color "rojo" del mapa colores, se puede escribir:

~~~go
codigoRojo := colores["rojo"]
~~~
También es posible verificar si una clave específica existe en un mapa utilizando la siguiente sintaxis:

~~~go
valor, ok := miMap[clave]
~~~
En este caso, valor contiene el valor correspondiente a la clave clave, si existe en el mapa, y ok es un valor booleano que indica si la clave existe o no.

Los maps son una estructura de datos muy útil en Go, ya que permiten almacenar y recuperar datos de manera eficiente utilizando claves que pueden ser de diferentes tipos de datos. Los maps también son muy flexibles y se pueden utilizar en una amplia variedad de aplicaciones, como la creación de índices, el almacenamiento de configuraciones y la implementación de algoritmos de búsqueda y ordenación.

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
    fmt.Println("Hola, mi nombre es", p.Name)
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
## Intefaces
En Go, una interfaz es un tipo de dato abstracto que define un conjunto de métodos. Una interfaz describe un comportamiento que un tipo puede tener, pero no define cómo se implementa. Los métodos de una interfaz no tienen cuerpo, solo tienen la firma de la función.

Una interfaz se define mediante la palabra clave "interface" seguida de un conjunto de métodos que describe. Por ejemplo:

~~~go
type Animal interface {
    Comer()
    Mover()
    Hablar()
}
~~~

En este ejemplo, hemos definido una interfaz llamada "Animal" que describe tres métodos: Eat(), Move() y Speak(). Cualquier tipo que implemente estos tres métodos será considerado un Animal.

Para implementar una interfaz, un tipo debe proporcionar implementaciones para cada uno de los métodos de la interfaz. Por ejemplo:

~~~ go
type Perro struct {
    Nombre string
}

func (d *Perro) Comer() {
    fmt.Println(d.Name, "está comiendo")
}

func (d *Perro) Mover() {
    fmt.Println(d.Name, "se esta moviendo")
}

func (d *Perro) Hablar() {
    fmt.Println(d.Name, "está ladrando")
}
~~~

~~~ go
type Gato struct {
    Nombre string
}

func (d *Gato) Comer() {
    fmt.Println(d.Name, "está comiendo")
}

func (d *Gato) Mover() {
    fmt.Println(d.Name, "se esta moviendo")
}

func (d *Gato) Hablar() {
    fmt.Println(d.Name, "está maúllando")
}
~~~
En este ejemplo, hemos definido un tipo llamado "Dog" que implementa la interfaz "Animal" proporcionando implementaciones para los métodos Eat(), Move() y Speak().

Las interfaces son útiles porque permiten al código trabajar con diferentes tipos de manera genérica. Por ejemplo, supongamos que tenemos una función que espera un parámetro de tipo Animal:

~~~go
func HacerAlgo(a Animal) {
    a.Comer()
    a.Mover()
    a.SpeHablarak()
}
~~~
Cualquier tipo que implemente la interfaz "Animal" puede ser pasado a esta función y funcionará correctamente. Esto significa que podemos escribir código genérico que funciona con cualquier tipo que implemente una interfaz específica.

En resumen, las interfaces son un mecanismo poderoso en Go que permite al código trabajar con diferentes tipos de manera genérica.

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

