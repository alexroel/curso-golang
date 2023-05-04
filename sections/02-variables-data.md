# Usu de variables, funciones y paquetes
 
1. [Variables](#Variables)
2. [Tipos de datos](#Variables)
3. [Valores cero](#valores-cero)
4. [Constantes](#Constantes)
5. [Operadores matemáticos](#Operadores-matemáticos)
6. [Paquetes Math](#Paquetes-Math)
7. [Proyecto de la sección](#Proyecto-de-la-sección)

--- 
## Variables 
En Go, las variables son utilizadas para almacenar y manipular datos. Una variable es un espacio en la memoria donde se puede almacenar un valor. En Go, todas las variables tienen un tipo de dato, el cual determina qué tipo de valor puede ser almacenado en esa variable.

### Declaración de variables
Para declarar una variable en Go, se utiliza la sintaxis `var nombreDeVariable tipoDeDato`. Por ejemplo, para declarar una variable de tipo entero (int) llamada "edad", se utilizaría el siguiente código:

~~~go
var edad int
~~~

### Variables inicializadas
También es posible inicializar una variable al mismo tiempo que se declara, utilizando la sintaxis `var nombreDeVariable tipoDeDato = valor`. Por ejemplo, para declarar e inicializar una variable de tipo cadena (string) llamada "nombre" con el valor "Juan", se utilizaría el siguiente código:

~~~go
var nombre string = "Juan"
~~~

### Declaraciones de variables cortas
Go también cuenta con la sintaxis de declaraciones de variables cortas `nombreDeVariable := valor`, la cual permite declarar e inicializar una variable en la misma línea de código. Go deduce automáticamente el tipo de dato basado en el valor asignado a la variable. Por ejemplo, para declarar e inicializar una variable de tipo flotante (float32) llamada "precio" con el valor 2.50, se utilizaría el siguiente código:

~~~
precio := 2.50
~~~

### Múltiples variables
Go también permite declarar múltiples variables en una sola línea, separando cada variable con una coma. Por ejemplo, para declarar dos variables de tipo entero llamadas "edad" y "altura", se utilizaría el siguiente código:

~~~go
var edad, altura int
~~~

También es posible inicializar múltiples variables en una sola línea, utilizando la sintaxis `var nombreDeVariable1, nombreDeVariable2 = valor1, valor2`.

~~~go
var nombre, apellido = "Juan", "Pérez"
~~~

Espero que esta información adicional te haya sido útil para entender mejor cómo trabajar con variables en Go.

---
## Tipos de datos 
En Go, existen diferentes tipos de datos que se pueden utilizar para almacenar diferentes tipos de valores. A continuación, se describen algunos de los tipos de datos más comunes en Go:

### Tipos de datos numéricos
- `int`: representa valores enteros, puede ser de diferentes tamaños (`int8`, `int16`, `int32`, `int64`)
- `uint`: representa valores enteros positivos sin signo, también puede ser de diferentes tamaños (`uint8`, `uint16`, `uint32`, `uint64`)
- `float32`, `float64`: representan números decimales con diferentes niveles de precisión

### Tipo de datos booleano
- `bool`: representa valores booleanos (verdadero o falso)

### Tipo de datos de texto
- `string`: representa una cadena de caracteres

### Tipos de datos compuestos
- `array`: es una colección de elementos del mismo tipo y tamaño fijo
- `slice`: es una colección de elementos del mismo tipo y tamaño variable
- `map`: es una colección de pares clave-valor
- `struct`: es una colección de campos con nombre, cada uno con un tipo de dato

Go también cuenta con tipos de datos adicionales para manejar direcciones de memoria, errores, funciones y más. Es importante recordar que en Go, todas las variables tienen un tipo de dato y que el tipo de dato determina qué valores se pueden almacenar en la variable y qué operaciones se pueden realizar con ella.

---
## Valores cero
En Go, todas las variables tienen un valor predeterminado, conocido como valor cero. El valor cero depende del tipo de dato de la variable y se utiliza cuando no se ha asignado un valor explícitamente a la variable.

A continuación, se muestran los valores cero de algunos de los tipos de datos más comunes en Go:

- `int`: `0`
- `float`: `0.0`
- `bool`: `false`
- `string`: cadena vacía (`""`)
- `array`: un array vacío
- `slice`: un slice nulo (`nil`)
- `map`: un map nulo (`nil`)
- `struct`: todas las variables dentro de la estructura tienen su propio valor cero

Es importante tener en cuenta que el valor cero no siempre es el mismo que el valor nulo (`nil`). Por ejemplo, el valor cero de un slice es un slice nulo, no un slice vacío. Es importante comprender la diferencia para evitar errores en la programación.

En resumen, es importante conocer los valores cero de los diferentes tipos de datos en Go para comprender cómo funcionan las variables en diferentes contextos y evitar errores al programar.

---
## Constantes 
En Go, las constantes son variables cuyo valor no cambia durante la ejecución del programa. Se definen utilizando la palabra clave `const` seguida de un identificador y un valor. Por ejemplo:

~~~go
const pi = 3.14159
~~~
En este ejemplo, `pi` es una constante que se utiliza para representar el valor de Pi, con un valor de `3.14159`.

Es importante tener en cuenta que, a diferencia de las variables normales, las constantes no se pueden modificar una vez que se han definido. Intentar cambiar el valor de una constante generará un error de compilación.

Las constantes pueden ser de cualquier tipo de dato básico, incluyendo enteros, flotantes, booleanos y cadenas. También se pueden definir constantes numéricas en diferentes bases, como octal y hexadecimal.

~~~go
const (
    x = 100
    y = 0b1010  // binario
    z = 0o12    // octal
    w = 0xFF    // hexadecimal
)
~~~

En este ejemplo, se definen varias constantes numéricas en diferentes bases.

En resumen, las constantes en Go son variables cuyo valor no cambia y se definen utilizando la palabra clave `const`. Son útiles para definir valores que no deben ser modificados durante la ejecución del programa.

---
## Operadores matemáticos 
En Go, hay varios tipos de operadores matemáticos que se utilizan para realizar operaciones aritméticas, incrementar o decrementar valores, y asignar valores a variables. A continuación se describen los operadores más comunes:

### Operadores aritméticos:
- `+` : suma dos valores.
- `-` : resta dos valores.
- `*` : multiplica dos valores.
- `/` : divide el primer valor por el segundo valor.
- `%` : devuelve el resto de la división entre dos valores (módulo).

~~~go
a := 10
b := 3
fmt.Println(a + b)  // 13
fmt.Println(a - b)  // 7
fmt.Println(a * b)  // 30
fmt.Println(a / b)  // 3
fmt.Println(a % b)  // 1
~~~

### Operadores de incremento y decremento:
- `++` : incrementa el valor de una variable en 1.
- `--` : decrementa el valor de una variable en 1.

~~~go
a := 10
a++
fmt.Println(a)  // 11

b := 5
b--
fmt.Println(b)  // 4
~~~

### Operadores en asignación:
- `=` : asigna el valor de la derecha a la variable de la izquierda.
- +`=` : suma el valor de la derecha a la variable de la izquierda y luego asigna el resultado a la variable de la izquierda.
- `-=` : resta el valor de la derecha a la variable de la izquierda y luego asigna el resultado a la variable de la izquierda.
- `*=` : multiplica el valor de la derecha a la variable de la izquierda y luego asigna el resultado a la variable de la izquierda.
- `/=` : divide el valor de la variable de la izquierda por el valor de la derecha y luego asigna el resultado a la variable de la izquierda.
- `%=` : devuelve el resto de la división entre la variable de la izquierda y la variable de la derecha y luego asigna el resultado a la variable de la izquierda.

~~~go
a := 10
a += 5
fmt.Println(a)  // 15

b := 20
b -= 3
fmt.Println(b)  // 17

c := 7
c *= 3
fmt.Println(c)  // 21

d := 100
d /= 5
fmt.Println(d)  // 20

e := 15
e %= 4
fmt.Println(e)  // 3
~~~
En resumen, los operadores matemáticos en Go se utilizan para realizar operaciones aritméticas, incrementar o decrementar valores, y asignar valores a variables. Es importante conocer los diferentes operadores y cómo se utilizan para realizar operaciones matemáticas y asignar valores a variables en Go.

---
## Paquetes Math
El paquete math en Go proporciona funciones matemáticas comunes que se utilizan para realizar cálculos matemáticos complejos. Este paquete incluye funciones para trabajar con valores numéricos, como operaciones trigonométricas, logaritmos, exponenciales, funciones de redondeo y más.

Para utilizar el paquete math, simplemente se debe importar en el archivo Go y luego llamar a la función apropiada. A continuación se muestra un ejemplo de algunas de las funciones disponibles en el paquete math:

### Constantes en Math
El paquete math no proporciona constantes predefinidas, pero sí tiene varias funciones que retornan constantes matemáticas, por ejemplo:

- `math.Pi`: retorna el valor de Pi (π).
- `math.E`: retorna el valor de la constante matemática e.
Estas constantes se pueden utilizar en cualquier programa que utilice el paquete math. Aquí un ejemplo:ç

~~~go 
fmt.Println("PI: ", math.Pi) //PI:  3.141592653589793
fmt.Println("E: ", math.E)   //E:  2.718281828459045
~~~

### Funciones de potencia
El paquete math proporciona varias funciones para calcular potencias de un número en Go. A continuación se muestran las funciones de potencia disponibles:

- `math.Pow(x, y float64) float64`: devuelve x elevado a la potencia de y.
- `math.Sqrt(x float64) float64`: devuelve la raíz cuadrada de x.
- `math.Cbrt(x float64) float64`: devuelve la raíz cúbica de x.
- `math.Pow10(n int) float64`: devuelve 10 elevado a la potencia de n.


~~~go
// Ejemplo de potencias
fmt.Println(math.Pow(2, 3)) // Imprime: 8
fmt.Println(math.Sqrt(64)) // Imprime: 8
fmt.Println(math.Cbrt(27)) // Imprime: 3
fmt.Println(math.Pow10(2)) // Imprime: 100
~~~

En resumen, el paquete math es muy útil para realizar operaciones matemáticas en Go, ya que proporciona una gran cantidad de funciones matemáticas estándar que se pueden utilizar en cualquier programa com: 
- Funciones de redondear un número
- Funciones trigonométricas
- Funciones exponenciales y logarítmicas

---
## Proyecto de la sección
Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.

El programa debe:

- Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.
- Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.
- Calcular el área del triángulo usando la fórmula base x altura / 2.
- Calcular el perímetro del triángulo sumando los lados.
- Imprimir el área y el perímetro del triángulo con dos decimales de precisión.
- El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro. También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.

Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.

Ejemplo de entrada y salida:

~~~
Ingrese lado 1: 3.5
Ingrese lado 2: 4.2

Área: 7.35
Perímetro: 12.20
~~~
Por supuesto, aquí te presento el código para crear un programa en Go que calcule el área y el perímetro de un triángulo rectángulo:

~~~go
package main

import (
    "fmt"
    "math"
)

func main() {
    var lado1, lado2 float64
    const precision = 2

    fmt.Print("Ingrese lado 1: ")
    fmt.Scanln(&lado1)

    fmt.Print("Ingrese lado 2: ")
    fmt.Scanln(&lado2)

    hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
    area := (lado1 * lado2) / 2
    perimetro := lado1 + lado2 + hipotenusa

    fmt.Printf("\nÁrea: %.*f\n", precision, area)
    fmt.Printf("Perímetro: %.*f\n", precision, perimetro)
}
~~~

Para imprimir los resultados, usamos la función fmt.Printf() para formatear la salida y mostrar los valores calculados. La cadena de formato "%.*f" se utiliza para indicar que se mostrará un número de punto flotante con un número determinado de decimales (especificado por la constante precision), y el operador % se utiliza para indicar qué valor se utilizará en ese lugar de la cadena de formato.



