# Entrada y salida en la consola

1. [Salida por consolo](#Salida-por-consolo)
2. [Entrada por consola](#Entrada-por-consola )
3. [Formater datosr](#Formater-datos)
4. [Entrada datos en el juego](#Entrada-datos-en-el-juego)

---
## Salida por consolo
- `Print()`: Imprime datos en la consola, y devuelve el número de bytes escritos y cualquier error de escritura encontrado.
- `Println()`: Imprime datos en la consola y salta a una nueva linea, y devuelve el número de bytes escritos y cualquier error de escritura encontrado.
- `Printf()`: Formatea de acuerdo con un especificador de formato y escribe en la salida estándar. Devuelve el número de bytes escritos y cualquier error de escritura encontrado.

### Verbos para formatear datos  
- General 
    - `%v` Formato predeterminado 
    - `%T` Muestra el tipo de dato 
    - `%t` Booleano 
    - `%d` Números enteros 
    - `%g` Números flotantes 
    - `%s` Cadena de caracteres 

Ejemplo:

~~~go
	// Entra y salida por consola
    var (
        name string
        age int
        size float64
    )

	// name := "Alex Roel"
	// age := 27
	// size := 1.67

	fmt.Println("Datos de usuario")
	fmt.Printf("Nombre: %s\nEdad: %d\nTalla: %g\n", name, age, size)
~~~

--- 
## Entrada por consola 

Para ingresar datos por la consola podemos utilzar dunciones del paquete `fmt`.

- `Scan()` y `Scanln()`: Escanea el texto leído desde la entrada estándar, almacenando valores sucesivos separados por espacios en argumentos sucesivos. Las nuevas líneas cuentan como espacio. Devuelve el número de elementos escaneados con éxito. Si es menor que el número de argumentos, err informará por qué.

- `Scanf()`: Escanea el texto leído desde la entrada estándar, almacenando valores sucesivos separados por espacios en argumentos sucesivos según lo determine el formato. Devuelve el número de elementos escaneados con éxito. Si es menor que el número de argumentos, err informará por qué. Los saltos de línea en la entrada deben coincidir con los saltos de línea en el formato. La única excepción: el verbo %c siempre escanea la siguiente runa en la entrada, incluso si es un espacio (o tabulación, etc.) o una nueva línea.

Ejemplo:
~~~go
	// Entra y salida por consola
	var (
		name1 string
		name2 string
		age   int
		size  float64
	)
	// name := "Alex Roel"
	// age := 27
	// size := 1.67

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name1, &name2)

	fmt.Print("Ingrese su Edad: ")
	fmt.Scanln(&age)

	fmt.Print("Ingrese su Talla: ")
	fmt.Scanln(&size)

	fmt.Println("Datos de usuario")
	fmt.Printf("Nombre: %s %s\nEdad: %d\nTalla: %g\n", name1, name2, age, size)
~~~

---
### Formater datos 

- `Sprint()`, `Sprintln()`: Formatea usando los formatos predeterminados para sus operandos y devuelve la cadena resultante. Se agregan espacios entre operandos cuando ninguno es una cadena.

- `Sprintf()`: Formatea de acuerdo con un especificador de formato y devuelve la cadena resultante.

Ejemplo: 
~~~go
// Entra y salida por consola
	var (
		name1 string
		name2 string
		age   int
		size  float64
	)
	// name := "Alex Roel"
	// age := 27
	// size := 1.67

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name1, &name2)

	fmt.Print("Ingrese su Edad: ")
	fmt.Scanln(&age)

	fmt.Print("Ingrese su Talla: ")
	fmt.Scanln(&size)

	userData := fmt.Sprintf("Nombre: %s %s\nEdad: %d\nTalla: %g\n", name1, name2, age, size)
	//fmt.Printf("Nombre: %s %s\nEdad: %d\nTalla: %g\n", name1, name2, age, size)
	fmt.Println("Datos de usuario")
	fmt.Println(userData)
~~~

---
## Entrada datos en el juego
Para que el usuario adivine el número aleatorio tiene que ingresar por cosolo, ahora vamos implementar entrada de datos y guardar en una variable. 

Tambien el número aleatorio que generamos guardaremos en una variable para poder comparar luego. 

~~~go
    //Generar numero aleatorio
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101)

	var chosenNumber int
	//Optener número ingresado por el jugador
	fmt.Print("Ingrese un número entre 0 y 101: ")
	fmt.Scanln(&chosenNumber)

	fmt.Println("Resultado: ", randomNumber == chosenNumber)
~~~