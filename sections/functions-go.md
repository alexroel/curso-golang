# Funciones avanzadas

1. [Funciones variádicas](#Funciones-variádicas)
2. [Funciones recursivas](#Funciones-recursivas) 
3. [Funciones anónimas](#Funciones-anónimas)
4. [Funciones de orden superior](#funciones-de-orden-superior)
5. [Closures en Go](#closures-en-go)
6. [Funciones defer, Panic y Recover](#funciones-defer-panic-y-recover) 

---
## Funciones variádicas
En Go, las funciones variádicas son aquellas que pueden tomar un número variable de argumentos. Esto se logra utilizando la sintaxis ...tipo después del tipo del último argumento de la función. El valor tipo especifica el tipo de los argumentos variables que se pueden pasar a la función.

Por ejemplo, la función suma a continuación toma un número variable de argumentos de tipo int y devuelve la suma de todos los argumentos:

~~~go
func suma(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
~~~

En la declaración de la función, el parámetro nums es de tipo ...int, lo que indica que puede tomar cualquier número de argumentos de tipo int.

Dentro de la función, se inicializa una variable total en 0 y se recorren todos los elementos del slice de argumentos nums utilizando el bucle range. En cada iteración, se suma el elemento actual al total.

Luego, se devuelve el valor total.

La función suma se puede llamar con cualquier número de argumentos de tipo int:
~~~go
fmt.Println(suma(1, 2, 3)) // Output: 6
fmt.Println(suma(4, 5, 6, 7, 8)) // Output: 30
~~~

En este ejemplo, la función suma se llama con dos conjuntos diferentes de argumentos de tipo int. En ambos casos, la función devuelve la suma de todos los argumentos.

---
## Funciones recursivas
La recursión es un concepto en programación en el que una función se llama a sí misma de manera repetitiva hasta que se cumple una determinada condición de salida. Una función recursiva es una función que contiene una llamada a sí misma dentro de su definición.

En Go, las funciones recursivas se definen de manera similar a cualquier otra función, pero se incluye una llamada a sí misma dentro del cuerpo de la función. Por ejemplo, a continuación se muestra una función recursiva que calcula el factorial de un número entero:

~~~go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
~~~

En este ejemplo, la función factorial toma un parámetro entero `n` y devuelve el factorial de `n`. El factorial de un número entero `n` se define como `n * (n-1) * (n-2) * ... * 1`.

La función factorial utiliza una condición `if` para comprobar si el valor de `n` es `0`. Si es así, la función devuelve `1`, que es el factorial de `0`. Si `n` es diferente de `0`, la función devuelve el producto de `n` y el factorial de `n-1`. La llamada a `factorial(n-1)` es la llamada recursiva, que se repite hasta que la condición `n == 0` se cumple y la función devuelve 1.

Es importante tener en cuenta que las funciones recursivas pueden ser ineficientes y consumir mucha memoria si se usan de manera inadecuada. Por lo tanto, siempre se debe asegurar de que la recursión tenga una condición de salida clara y que no se llame a sí misma infinitamente.

---
## Funciones anónimas
Las funciones anónimas son funciones que no tienen un nombre identificador y que pueden ser definidas en el lugar donde son necesitadas, en lugar de ser definidas globalmente. En Go, las funciones anónimas se crean utilizando la palabra clave `func` seguida de los parámetros y el cuerpo de la función entre llaves `{}`. A continuación se muestra un ejemplo de una función anónima:

~~~go
salud := func(name string) {
	fmt.Printf("Hola, %s!\n", name)
}
salud("Alex")
~~~

### Funciones como valores
Por ejemplo, la siguiente función saludar toma una función como argumento y luego llama a la función pasada con el nombre `name`:

~~~go
// Mis funciones
func saludar(name string, f func(string)) {
	f(name)
}

func main() {
	salud := func(name string) {
		fmt.Printf("Hola, %s!\n", name)
	}
	salud("Alex")
	saludar("Roel", salud)
}
~~~

Otro ejemplo de uso de funciones como valores es el siguiente:

~~~go
func duplicar(n int) int {
    return n * 2
}

func triplicar(n int) int {
    return n * 3
}

func aplicar(f func(int) int, n int) int {
    return f(n)
}

func main() {
    resultado1 := aplicar(duplicar, 5) // resultado1 = 10
    resultado2 := aplicar(triplicar, 5) // resultado2 = 15

    fmt.Println(resultado1, resultado2)
}
~~~

En este ejemplo, se definen dos funciones `duplicar` y `triplicar`, que toman un argumento `n` de tipo `int` y devuelven el valor de `n` multiplicado por 2 y 3, respectivamente.

Luego, se define una función aplicar que toma dos argumentos: `f` de tipo `func(int) int` y n de tipo int. Dentro de la función, se llama a la función f pasando el argumento n.

En la función main, se llama a la función aplicar dos veces con diferentes funciones y argumentos. El resultado de cada llamada se almacena en una variable y se imprime en la consola.

---
## Funciones de orden superior
En Go, una función de orden superior es una función que acepta otra función como argumento y/o devuelve una función como resultado. Esto permite una programación más modular y flexible, ya que las funciones de orden superior pueden utilizarse para crear funciones personalizadas en tiempo de ejecución.

Las funciones de orden superior se utilizan para implementar patrones de programación comunes, como el mapeo, filtrado y reducción de datos en colecciones. Por ejemplo, la función map() es una función de orden superior que toma una función como argumento y la aplica a cada elemento de una colección, devolviendo una nueva colección con los resultados de la aplicación de la función a cada elemento.

Aquí hay un ejemplo de una función de orden superior que acepta una función f y un entero x como argumentos, y devuelve el resultado de aplicar la función f al entero x multiplicado por dos:

~~~go
func double(f func(int) int, x int) int {
    return f(x * 2)
}
~~~

Esta función puede llamarse con cualquier función que tome un entero como argumento y devuelva un entero como resultado. Por ejemplo:

~~~go
func addOne(x int) int {
    return x + 1
}

result := double(addOne, 3) // Devuelve 7, ya que addOne(6) = 7
~~~

En este ejemplo, la función addOne se utiliza como argumento para la función double(), que se encarga de aplicar la función addOne al entero 6 (calculado a partir del entero 3 multiplicado por 2) y devuelve el resultado 7.

En resumen, las funciones de orden superior son una característica poderosa de Go que permiten una programación más modular y flexible al permitir la creación de funciones personalizadas en tiempo de ejecución mediante el uso de otras funciones como argumentos y/o resultados.

---
## Closures en Go

En Go, un closure (o clausura) es una función anónima que se define dentro de otra función y que puede acceder y modificar variables definidas en el ámbito de la función externa. Esto significa que un closure puede "recordar" el valor de las variables del ámbito de la función en el momento en que se define y utilizarlas en cualquier momento posterior en que se llame al closure.

Un closure en Go se define de la siguiente manera:

~~~go
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
~~~

En este ejemplo, la función incrementer devuelve una función anónima que tiene acceso a la variable local i definida en la función `incrementer`. La función anónima incrementa i en uno cada vez que se llama, y devuelve el valor actual de i.

Aquí hay un ejemplo de cómo se puede utilizar la clausura incrementer para crear una secuencia de números enteros:
~~~go
func main() {
	nextInt := incrementer()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

~~~

En este ejemplo, se llama a incrementer para obtener la función anónima que mantiene el estado de la variable i. Luego, se llama a la función anónima tres veces utilizando la variable nextInt. Cada vez que se llama a la función, la variable i se incrementa y se devuelve el nuevo valor.

En resumen, las clausuras son una característica poderosa de Go que permiten definir funciones anidadas que tienen acceso a las variables locales de la función padre. Se pueden utilizar para implementar funciones que mantienen un estado interno o para definir funciones genéricas que aceptan funciones como argumentos y las adaptan a un contexto específico.

---
## Funciones defer, Panic y Recover 
En Go, las funciones defer, panic y recover se utilizan juntas para manejar errores y controlar el flujo de ejecución en programas.

### Defer 
La función defer se utiliza para programar una función que se ejecutará justo antes de que una función que la contiene devuelva su valor. Esto se utiliza a menudo para asegurarse de que ciertas operaciones se realicen antes de que una función termine, independientemente de si la función se termina normalmente o por un error. Por ejemplo:

~~~go
func main() {
    defer fmt.Println("Adiós!")
    fmt.Println("Hola!")
}
~~~

En este ejemplo, la función fmt.Println("Adiós!") se ejecutará justo antes de que la función main() termine, incluso si la ejecución se termina normalmente o por un error.

### Panic
La función panic se utiliza para detener la ejecución normal de un programa y lanzar un error. Esto se utiliza cuando ocurre algo inesperado o cuando se encuentra un error que no se puede manejar en la función actual. Por ejemplo:

~~~go
func dividir(a, b int) int {
    if b == 0 {
        panic("No se puede dividir por cero!")
    }
    return a / b
}

func main() {
    resultado := dividir(10, 0)
    fmt.Println(resultado)
}
~~~

En este ejemplo, si el segundo argumento de la función dividir() es cero, la función lanzará un error con el mensaje "No se puede dividir por cero!" y detendrá la ejecución del programa.

### Recover
La función recover se utiliza para manejar errores lanzados por la función panic y controlar el flujo de ejecución en el programa. La función recover debe utilizarse dentro de una función defer para poder recuperar el error y controlar su manejo. Por ejemplo:

~~~go
func dividir(a, b int) (resultado int, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("Error recuperado: %v", p)
        }
    }()
    if b == 0 {
        panic("No se puede dividir por cero!")
    }
    resultado = a / b
    return resultado, nil
}

func main() {
    resultado, err := dividir(10, 0)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(resultado)
}

~~~
En este ejemplo, si el segundo argumento de la función dividir() es cero, la función lanzará un error con el mensaje "No se puede dividir por cero!" y detendrá la ejecución del programa. La función recover se utiliza dentro de una función defer para recuperar el error lanzado por la función panic y asignarlo a la variable err. Si se produce un error, el valor de err será distinto de nil y se imprimirá un mensaje de error en la salida estándar. Si no se produce un error, se imprimirá el resultado de la división.

En resumen, las funciones defer, panic y recover se utilizan juntas en Go para manejar errores y controlar el flujo de ejecución en programas. La función defer se utiliza para programar una función que se ejecutará justo antes de que una función que la contiene devuelva su valor, mientras que la función panic se utiliza para detener la