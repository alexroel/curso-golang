# Comenzando a codificar
 
1. [Paquetes](#Paquetes)
2. [Variables](#Variables)
3. [Tipos Básicos](#Tipos-Básicos)
4. [Valores cero](#valores-cero)
5. [Constantes](#Constantes)
6. [Operadores](#Operadores)
 
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
)
 
func main() {
    fmt.Println("¡Hola Mundo de Go!")
 
    fmt.Println("Valor de PI: ", math.Pi)
    fmt.Println("Valor de 4^2: ", math.Pow(4, 2))
}
~~~
 
### Genera números aleatorios
Para generar números  aleatorios utilizados más paquetes de Go.
 
- `math/rand`: El paquete rand implementa generadores de números pseudoaleatorios inadecuados para trabajos sensibles a la seguridad.
- `time`: El paquete proporciona funcionalidad para medir y mostrar el tiempo.
 
~~~go
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {
    //Generar numero aleatorio
    rand.Seed(time.Now().UnixNano())
    fmt.Println(rand.Intn(101))
}
~~~
- `rand.Intn(101)`: Devuelve, como int, un número pseudoaleatorio no negativo en el intervalo semiabierto [0,n) del origen predeterminado.
- `rand.Seed()`: Se utiliza el valor de `Seed()` proporcionado para inicializar el origen predeterminado a un estado determinista. Si no se llama `Seed()`, el generador se comporta como si fuera sembrado por `Seed(1)`.
 
- `time.Now()`: Devuelve la hora local actual.
- `UnixNano()`: Devuelve tiempo como un tiempo Unix, el número de nanosegundos transcurridos desde el 1 de enero de 1970 UTC.
 
 
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
- Podemos inicializar una variable
- Si hay un inicializador, se puede omitir el tipo de dato.
- Dependiendo del dato que se asigne tomará el tipado.
 
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
- Fuera de una función, cada instrucción comienza con una palabra clave ( var, func, etc.) y, por lo tanto, la construcción `:=` no está disponible.
 
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
- `uint` Para números enteros positivos
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
## Constantes
- Los constantes no se puede modificar su valor
- Las constantes se declaran como variables, pero con la palabra clave `cosnt`.
- Las constantes pueden ser caracteres, cadenas, booleanos o valores numéricos.
- Las constantes no se pueden declarar usando la :=sintaxis.
 
~~~go
const (
    a = 10
    b = 20
)
 
func main() {
    const pi = 3.1415
 
    fmt.Println("Valor de PI: ", pi)
 
    fmt.Println("Valor de a: ", a)
    fmt.Println("Valor de b: ", b)
}
~~~
 
---
## Operadores
Los operadores no van ayudar a manipular los datos, realizar lógicas de programación entre otros.
 
### Operadores aritméticos
Los operadores aritméticos nos permiten realizar operaciones básicas sobre nuestras variables y constantes.
 
- `+`: Realiza una suma de dos valores o también puede concatenar caracteres
- `-`: Realiza un resta
- `*`: Realiza una multiplicación
- `/`: Realiza una división
- `%`: Devuelve el resto de una división
- `++`: Operador de incremento. Incrementa en 1 el operador de la izquierda. No permite el decremento prefijo (pre-decremento).
- `--`: Operador de decremento. Decrementa en 1 el operador de la izquierda, es decir, no permite el decremento prefijo.
 
### Operadores de asignación
Los operadores de asignación nos permiten modificar el valor de variables a lo largo de la ejecución de nuestro programa.
 
- `+=`: Suma en asignación
- `-=`: Resta en asignación
- `*=`: Multiplicación en asignación
- `/=`: División en asignación
- `%=`: Módulo en asignación
 
 
### Operadores relacionales
Los operadores relacionales son aquellos que devuelven un valor de verdad (verdadero o falso) cuando se resuelve una expresión:
 
- `==`: Devuelve un valor booleano dependiendo si son iguales o no.
- `!=`: Devuelve un valor booleano dependiendo si son distintos o no.
- `>` : Devuelve verdadero cuando el operando de la izquierda es mayor que el de la derecha.
- `<`: Devuelve verdadero cuando el operando de la izquierda es menor que el de la derecha.
- `>=` : Devuelve verdadero cuando el operador de la izquierda es mayor o igual al de la derecha.
- `<=`: Devuelve verdadero cuando el operador de izquierda es menor o igual a el de la derecha.
 
### Operadores lógicos
- `&&`: Operador lógico de conjunción (AND). Compara dos valores `bool` o expresiones relacionales y devuelve `true` cuando ambos son `true`.
- `||`: operador lógico de disyunción (OR). Compara dos valores `bool` o expresiones relacionales y devuelve `true` cuando uno de ellos es `true`.
- `!`: Operador de negación. Invierte (niega) el valor `bool` del operando.
 
### Expresiones anidadas
   - Paréntesis
   - Multiplicación / División
   - Suma / Resta
   - Operadores relacionales
   - Operadores lógicos
 
 
   ~~~go
    a, b := 10, 5
 
    result := (a*b-2*b) >= 20 && !((a % b) != 0)
    // a*b-2*b -> 50 - 10 = 40
    //a % b -> 0
    // 0 != 0 -> false -> true
    // (a*b-2*b) >= 20 => true
 
    fmt.Println("Resultado: ", result)
   ~~~

