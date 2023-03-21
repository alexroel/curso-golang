# Control de flujos

1. [Operadores booleanos](#Operadores-booleanos) 
2. [Declaración If - Else](#Declaración-If---Else)
3. [Declaración Switch](#Declaración-Switch)
4. [Bucle For](#Bucle-For)
5. [Break y continue](#Break-y-continue)
6. [Funciones](#Funciones)
7. [Proyecto de la sección](#Proyecto-de-la-sección)
   
---
## Operadores booleanos 
Los operadores relacionales y lógicos son utilizados en conjunto en las expresiones lógicas de Go para evaluar condiciones complejas y producir un resultado booleano (verdadero o falso).

### Operadores relacionales
Aquí te proporciono una lista de los operadores relacionales disponibles en Go:

- `==` : igual a
- !`=` : diferente de
- `>`: mayor que
- `>=` : mayor o igual que
- `<` : menor que
- `<=` : menor o igual que
Estos operadores son utilizados para comparar valores y producir un resultado booleano (verdadero o falso) dependiendo de la relación entre los valores.

~~~go
x := 5
y := 10
z := x < y
fmt.Println(z) // Imprime true
~~~

### Operadores lógicos 
Los operadores lógicos en Go son utilizados para evaluar expresiones lógicas y producir un resultado booleano (verdadero o falso). En Go, existen tres operadores lógicos: AND lógico (&&), OR lógico (||) y NOT lógico (!).

#### Operador AND lógico (&&):
El operador && evalúa dos expresiones booleanas y devuelve verdadero (true) si ambas expresiones son verdaderas, y devuelve falso (false) si alguna de las expresiones es falsa. Por ejemplo:
~~~go
x := true
y := false
z := x && y
fmt.Println(z) // Imprime false
~~~

En este ejemplo, la variable z tendrá el valor de falso (false), ya que la expresión x && y evalúa a falso, debido a que la variable y es falsa.

#### Operador OR lógico (||):
El operador || evalúa dos expresiones booleanas y devuelve verdadero (true) si al menos una de las expresiones es verdadera, y devuelve falso (false) si ambas expresiones son falsas. Por ejemplo:
~~~go
x := true
y := false
z := x || y
fmt.Println(z) // Imprime true
~~~

En este ejemplo, la variable z tendrá el valor de verdadero (true), ya que la expresión x || y evalúa a verdadero, debido a que la variable x es verdadera.

#### Operador NOT lógico (!):
El operador ! niega el valor booleano de una expresión, es decir, si una expresión es verdadera, la niega y devuelve falso, y si una expresión es falsa, la niega y devuelve verdadero. Por ejemplo:
~~~go
x := true
z := !x
fmt.Println(z) // Imprime false
~~~

En este ejemplo, la variable z tendrá el valor de falso (false), ya que la expresión !x niega el valor de x, que es verdadero, y devuelve falso.

---
## Declaración If - Else
En Go, la estructura de control if ... else se utiliza para ejecutar un bloque de código si se cumple una condición booleana, y otro bloque de código si la condición no se cumple.

La sintaxis básica de la estructura if ... else es la siguiente:

~~~go
if condicion {
  // código a ejecutar si la condición es verdadera
} else {
  // código a ejecutar si la condición es falsa
}
~~~

Por ejemplo, se puede utilizar la estructura if ... else para determinar si es mañana, tarde o noche:

~~~go
	t := time.Now()

	if t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}
~~~
Estructura de condición If- Else acanzado.
~~~go

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}
~~~

---
## Declaración Switch
En Go, la declaración switch se utiliza para tomar decisiones basadas en el valor de una expresión. La sintaxis básica de la declaración switch es la siguiente:

~~~go
switch expression {
case value1:
  // código a ejecutar si expression == value1
case value2:
  // código a ejecutar si expression == value2
default:
  // código a ejecutar si ninguno de los casos anteriores se cumple
}
~~~
La expresión que se pasa a la declaración switch se evalúa y luego se compara con cada uno de los casos. Si el valor de la expresión coincide con uno de los valores de caso, se ejecuta el bloque de código correspondiente. Si ninguno de los casos coincide con el valor de la expresión, se ejecuta el bloque de código dentro del bloque default (opcional).

Aquí hay un ejemplo que muestra cómo se puede utilizar la declaración switch en Go:

~~~go
	//os := runtime.GOOS;
	//os := "linux"
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Go run -> Linux.")
	case "windows":
		fmt.Println("Go run -> Windows.")
	case "darwin":
		fmt.Println("Go run -> macOS.")
	default:
		fmt.Println("Go run -> Otro SO.")
	}
~~~
En este ejemplo, la variable dia se compara con cada uno de los casos dentro de la declaración switch. Si el valor de la variable dia coincide con uno de los casos, se imprime un mensaje correspondiente. Si no hay coincidencia, se imprime un mensaje dentro del bloque default.

También se puede utilizar una expresión en lugar de un valor en cada caso, lo que permite evaluar una expresión más compleja:

~~~go

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}
~~~

---
## Bucle For
En Go, el bucle for es la única estructura de control de repetición disponible. Sin embargo, se puede utilizar de diferentes maneras para lograr diferentes funcionalidades.

### Bucle infinito:
Para crear un bucle infinito en Go, se puede usar la sintaxis for {}:

~~~go
for {
    // Código que se repetirá infinitamente
}
~~~

### Bucle con una sola condición:
La sintaxis para crear un bucle con una sola condición en Go es la siguiente:

~~~go
for condición {
    // Código que se repetirá mientras la condición sea verdadera
}
~~~
Por ejemplo, el siguiente bucle se repetirá mientras i sea menor que 10:

~~~go
for i < 10 {
    fmt.Println(i)
    i++
}
~~~

### Bucle típico de for:
La sintaxis típica de un bucle for en Go es la siguiente:

~~~go
for inicialización; condición; actualización {
    // Código que se repetirá mientras la condición sea verdadera
}
~~~

Por ejemplo, el siguiente bucle imprimirá los números del 1 al 10:

~~~go
for i := 1; i <= 10; i++ {
    fmt.Println(i)
}
~~~
En este caso, la inicialización es i := 1, la condición es i <= 10, y la actualización es i++. El bucle se repetirá mientras la condición sea verdadera, e incrementará i en 1 en cada iteración.

---
## Break y continue
En Go, break y continue son palabras clave que se utilizan dentro de los bucles for para controlar el flujo de ejecución.

La palabra clave break se utiliza para salir de un bucle antes de que la condición de finalización se haya alcanzado. Cuando se ejecuta break, el control se transfiere a la siguiente instrucción después del bucle. Por ejemplo:

~~~go
for i := 1; i <= 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
~~~

En este caso, el bucle for se repetirá hasta que i sea igual a 5. Cuando i sea igual a 5, se ejecutará la instrucción break, lo que detendrá la ejecución del bucle y transferirá el control a la siguiente instrucción después del bucle.

Por otro lado, la palabra clave continue se utiliza para saltar a la siguiente iteración de un bucle sin ejecutar el código restante del cuerpo del bucle para la iteración actual. Cuando se ejecuta continue, el control se transfiere directamente al inicio del bucle para comenzar la siguiente iteración. Por ejemplo:

~~~go
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
~~~

En este caso, el bucle for imprimirá sólo los números impares del 1 al 10. Cuando i es un número par, se ejecutará la instrucción continue, lo que saltará directamente a la siguiente iteración del bucle sin ejecutar el código restante del cuerpo del bucle para la iteración actual.

---
## Funciones
En Go, se puede declarar una función utilizando la siguiente sintaxis:

~~~go
func nombreFuncion(parametro1 tipoParametro1, parametro2 tipoParametro2) tipoRetorno {
    // Código de la función
    return valorRetorno
}
~~~

Donde:

- `func`: es la palabra clave que se utiliza para declarar una función.
- `nombreFuncion`: es el nombre que se le da a la función.
- `parametro1`, `parametro2`, etc.: son los parámetros que recibe la función, cada uno con su tipo correspondiente separado por coma.
- `tipoRetorno`: es el tipo de dato que devuelve la función. Si la función no devuelve nada, se utiliza la palabra clave void.
- `valorRetorno`: es el valor que devuelve la función si es que tiene un tipo de retorno definido.

Por ejemplo, para declarar una función que sume dos números enteros y devuelva el resultado, se podría utilizar el siguiente código:

~~~go
func sumar(a int, b int) int {
    return a + b
}
~~~

En este ejemplo, la función se llama sumar, recibe dos parámetros de tipo int llamados a y b, y devuelve un valor de tipo int que es la suma de a y b.

### Funciones que devuelven múltiples valores
En Go, es posible que una función devuelva múltiples valores. Esto se logra simplemente especificando los tipos de los valores de retorno separados por comas en la declaración de la función.

Por ejemplo, la siguiente función llamada dividir toma dos números y devuelve el resultado de la división y el resto de la división como dos valores distintos:

~~~go
func dividir(dividendo, divisor int) (int, int) {
    cociente := dividendo / divisor
    resto := dividendo % divisor
    return cociente, resto
}
~~~

En la declaración de la función, se especifica que la función devuelve dos valores de tipo int separados por coma: (int, int).

Dentro de la función, se realiza la división y se calcula el resto y se devuelven ambos valores utilizando la palabra clave return seguida de los valores separados por coma: return cociente, resto.

Luego, al llamar a esta función, se pueden asignar los valores de retorno a dos variables separadas:

~~~go
cociente, resto := dividir(20, 3)
fmt.Println(cociente) // Output: 6
fmt.Println(resto) // Output: 2
~~~

En este ejemplo, se llama a la función dividir con los parámetros 20 y 3, lo que devuelve los valores 6 y 2. Luego, se asignan estos valores a las variables cociente y resto, respectivamente, utilizando la sintaxis de asignación múltiple.

---
## Proyecto de la sección 
Desarrollar un programa que permita al usuario jugar el famoso juego "Adivina el número". El programa deberá incluir los siguientes elementos:

1. Generación de un número aleatorio entre 1 y 100 que será el número que el usuario deberá adivinar.
2. Preguntar al usuario que ingrese un número.
Verificar si el número ingresado es igual, mayor o menor que el número generado aleatoriamente.
3. Si el número ingresado es igual al número generado, mostrar un mensaje de felicitación y preguntar si el usuario desea volver a jugar.
4. Si el número ingresado es mayor o menor que el número generado, mostrar un mensaje indicando la situación y permitir al usuario ingresar un nuevo número.
5. El usuario tendrá un número limitado de intentos para adivinar el número, por lo que se deberá mostrar cuántos intentos le quedan.
6. Si el usuario no adivina el número después de los intentos permitidos, mostrar un mensaje indicando que el juego ha terminado y preguntar si desea volver a jugar.
7. Si el usuario desea volver a jugar, el programa debe volver a generar un número aleatorio y comenzar el juego nuevamente.

El programa deberá utilizar los siguientes elementos de control de flujo:

1. Generar un número aleatorio usando la biblioteca "math/rand".
2. Utilizar la declaración "if - else" para verificar si el número ingresado es igual, mayor o menor que el número generado aleatoriamente.
3. Utilizar la declaración "for" para permitir al usuario tener un número limitado de intentos para adivinar el número.
4. Utilizar la declaración "break" para salir del bucle "for" si el usuario adivina el número o si se agotan los intentos.
5. Utilizar la declaración "switch" para preguntar al usuario si desea volver a jugar y ejecutar la acción correspondiente.
6. Utilizar funciones para estructurar el código y hacerlo más legible.

### Código de proyecto 
~~~go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	jugar()
}

func jugar() {
	numeroAleatorio := rand.Intn(100) + 1
	var numeroIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numeroIngresado)

		if numeroIngresado == numeroAleatorio {
			fmt.Println("¡Felicitaciones, adivinaste el número!")
			jugarNuevamente()
			return
		} else if numeroIngresado < numeroAleatorio {
			fmt.Println("El número a adivinar es mayor.")
		} else if numeroIngresado > numeroAleatorio {
			fmt.Println("El número a adivinar es menor.")
		}
	}

	fmt.Println("Se acabaron los intentos. El número era:", numeroAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Print("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Elección inválida. Inténtalo nuevamente.")
		jugarNuevamente()
	}
}
~~~

### Generar número aleatorio
Go tiene una biblioteca integrada para trabajar con números aleatorios. Esta biblioteca se llama `math/rand`. Para generar un número aleatorio en Go, primero necesitamos crear una fuente de números aleatorios o un generador de números aleatorios. Luego, podemos utilizar el generador para generar números aleatorios.

 
~~~go
// Crea una nueva fuente de números aleatorios usando el tiempo actual como semilla
rand.Seed(time.Now().UnixNano())

// Genera un número aleatorio entre 0 y 99
randomNumber := rand.Intn(100)

fmt.Println(randomNumber)
~~~

1. Primero, se llama a la función `time.Now()` que devuelve el tiempo actual en formato `time.Time`.
2. Luego, se llama a la función `UnixNano()` de la estructura time.Time para obtener el tiempo actual en nanosegundos.
3. A continuación, se utiliza el valor obtenido como semilla para la fuente de números aleatorios llamando a la función `rand.Seed()`. Esto inicializa la fuente de números aleatorios con una semilla única que se basa en el tiempo actual.
4. Se utiliza la función `rand.Intn(100)` para generar un número aleatorio entero entre 0 y 99. Esta función devuelve un número entero pseudoaleatorio dentro del rango especificado.
5. El número aleatorio generado se almacena en la variable `randomNumber`.
6. Finalmente, se imprime el número aleatorio generado en la consola utilizando la función `fmt.Println()`.

En resumen, este código genera un número aleatorio entre 0 y 99 y lo imprime en la consola cada vez que se ejecuta, utilizando el tiempo actual como semilla para la fuente de números aleatorios. Esto garantiza que se generen números diferentes cada vez que se ejecuta el programa.