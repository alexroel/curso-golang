# Control de errores 

1. [Introducción](#introducción)
2. [Manejo de errores](#manejo-de-errores)
3. [Ejemplo de Manejo de Errores](#ejemplo-de-manejo-de-errores)
4. [Uso de Defer](#uso-de-defer)
5. [Uso de panic y recover](#uso-de-panic-y-recover)
6. [Registro de errores](#registro-de-errores)
7. [Proyecto de sección: Gestor de contactos](#proyecto-de-sección-gestor-de-contactos)
7. [Resumen](#resumen)

---
## Introducción 
Hay ocasiones en las que los programas que se escriben no se comportan de la manera esperada. A veces, hay factores externos que no se pueden controlar, como que otros procesos bloqueen un archivo o un intento de acceder a una dirección de memoria que ya no está disponible. Los errores son simplemente otro tipo de comportamiento que los programas pueden tener. Es mejor que se anticipe a esos errores para poder solucionar los problemas cuando se produzcan.

En primer lugar, aprenderemos sobre el manejo de errores en Go. Veremos cómo utilizar la sintaxis if err != nil para verificar y manejar errores de una manera concisa y eficiente. También exploraremos las mejores prácticas para devolver y manejar errores de manera adecuada en nuestras funciones.

A continuación, profundizaremos en el ejemplo de manejo de errores en Go. Mediante un escenario práctico, analizaremos cómo implementar el manejo de errores en una aplicación real. Exploraremos técnicas como el uso de variables de error, el registro de errores y la propagación adecuada de errores en diferentes capas de nuestra aplicación.

Uno de los aspectos más interesantes del manejo de errores en Go es el uso de la declaración defer. Descubriremos cómo usar defer para asegurarnos de que ciertas operaciones se realicen antes de salir de una función, incluso en caso de error. Veremos ejemplos prácticos que demuestran cómo defer puede facilitar el manejo de tareas como la liberación de recursos o el cierre adecuado de archivos.

Otro tema importante relacionado con el manejo de errores es el uso de panic y recover. Aprenderemos cómo utilizar estas dos funciones para manejar situaciones excepcionales y errores graves en nuestras aplicaciones. Exploraremos cómo panic puede ser utilizado para interrumpir el flujo normal del programa y cómo recover nos permite recuperarnos de dichos errores y continuar la ejecución de manera controlada.

Finalmente, en el proyecto de esta sección, crearemos un gestor de contactos utilizando Go. A través de este proyecto, aplicaremos todos los conceptos que hemos aprendido sobre el manejo de errores en Go. Implementaremos una funcionalidad robusta para agregar, mostrar y almacenar contactos, asegurándonos de manejar los posibles errores que puedan ocurrir durante el proceso.

¡Únete a esta sección del curso y domina el arte del manejo de errores en Go! Estarás preparado para crear aplicaciones confiables y resistentes que sepan manejar adecuadamente situaciones inesperadas.

---
## Manejo de errores 
El control de errores en Go es una parte fundamental de la programación para garantizar la robustez y confiabilidad de las aplicaciones. En Go, los errores se manejan de manera explícita y se utilizan tipos especiales llamados "errores".

El manejo de errores en Go implica capturar y manejar los posibles errores que puedan ocurrir durante la ejecución del programa. Un ejemplo común es el manejo de errores al convertir una cadena a un número utilizando la función `strconv.Atoi()` de la biblioteca estándar de Go.

La función `strconv.Atoi()` se utiliza para convertir una cadena que representa un número en su equivalente entero. Sin embargo, si la cadena no puede ser convertida correctamente, la función devuelve un error. Aquí hay un ejemplo que muestra cómo manejar este error:

~~~go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"

    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Número:", num)
}
~~~
En este ejemplo, se define una cadena `str` que representa un número. Luego, se utiliza la función `strconv.Atoi()` para convertir esa cadena a un entero. Si ocurre un error durante la conversión, la función `Atoi()` devuelve un error, que se captura en la variable `err`. Dentro del bloque `if err != nil`, se verifica si el error es diferente de `nil`, lo que indica que se produjo un error durante la conversión. En ese caso, se imprime el mensaje de error y se detiene la ejecución del programa.

### Errores personalizados 
En Go, un error es simplemente un tipo de dato que implementa la interfaz `error`, que consiste en un único método: `Error() string`. Esta interfaz permite a los programadores definir sus propios tipos de error personalizados o utilizar los errores predefinidos proporcionados por el lenguaje.

Para crear un nuevo error en Go, se puede utilizar el paquete `errors` de la biblioteca estándar. Aquí hay un ejemplo:

~~~go
package main

import (
    "errors"
    "fmt"
)

func divide(dividendo, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errors.New("No se puede dividir por cero")
    }
    return dividendo / divisor, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Resultado:", result)
}
~~~

En este ejemplo, la función `divide` recibe dos enteros y devuelve el resultado de la división y un error. Si el divisor es cero, se crea un nuevo error utilizando la función `errors.New` y se devuelve junto con un valor de resultado de cero. En la función `main`, se verifica si el error es diferente de `nil` y, en caso afirmativo, se imprime el mensaje de error.

Además de los errores personalizados, Go también proporciona el tipo de error `error` predefinido, que es simplemente una cadena de texto. Se puede utilizar utilizando el paquete `fmt` de la biblioteca estándar:

~~~go
package main

import "fmt"

func divide(dividendo, divisor int) (int, error) {
    if divisor == 0 {
        return 0, fmt.Errorf("no se puede dividir por cero")
    }
    return dividendo / divisor, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Resultado:", result)
}
~~~
En este caso, se utiliza la función `fmt.Errorf` para crear un error con formato que incluye el mensaje de error.

En resumen, en Go se manejan los errores utilizando tipos de datos que implementan la interfaz `error`. Los errores se pueden crear utilizando el paquete `errors` o el paquete `fmt`, y se pueden verificar y manejar adecuadamente en el código para controlar el flujo de ejecución y proporcionar mensajes claros al usuario.

---
## Ejemplo de Manejo de Errores
Este código es un ejemplo práctico que ilustra el manejo de errores en Go. Muestra cómo controlar y gestionar los errores de manera efectiva durante la ejecución de un programa.

El objetivo de este ejemplo es realizar una búsqueda en una colección de alimentos representados por emojis. La búsqueda se realiza utilizando un identificador numérico. Si se encuentra el alimento correspondiente al identificador, se muestra el emoji correspondiente; de lo contrario, se informa que no se encontró ningún alimento.

El código implementa diversas técnicas para el manejo de errores. Utiliza el paquete de errores de la biblioteca estándar de Go para definir y gestionar errores personalizados. Además, hace uso de la función `errors.Is()` para verificar si un error específico se ha producido.

A lo largo del código, se aplican buenas prácticas de manejo de errores, como devolver errores explícitos, proporcionar mensajes de error descriptivos y utilizar el enfoque de "fallar rápido" para evitar ejecuciones innecesarias cuando se produce un error esperado.

Este ejemplo te ayudará a comprender cómo implementar el manejo de errores en tus propios proyectos en Go, brindándote una base sólida para construir aplicaciones más robustas y confiables.

~~~go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// errNotFound representa el error de elemento no encontrado.
var errNotFound = errors.New("no encontrado")

// food es un mapa que contiene alimentos representados por emojis.
var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

func main() {
	// Realizar una búsqueda con la clave "34".
	encontrado, err := buscar("34")
	if errors.Is(err, errNotFound) {
		fmt.Println("Se pudo manejar el error correctamente")
		return
	}
	if err != nil {
		fmt.Println("buscar()", err)
		return
	}

	fmt.Println(encontrado)
}

// buscar realiza una búsqueda de un emoji de comida utilizando una clave.
func buscar(clave string) (string, error) {
	// Convertir la clave a un número entero.
	num, err := strconv.Atoi(clave)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %w", err)
	}

	// Encontrar el emoji de comida correspondiente al número.
	emoji, err := encontrarComida(num)
	if err != nil {
		return "", fmt.Errorf("encontrarComida(): %w", err)
	}

	return emoji, nil
}

// encontrarComida busca el emoji de comida correspondiente a un identificador.
func encontrarComida(id int) (string, error) {
	// Verificar si el identificador existe en el mapa de alimentos.
	valor, existe := food[id]
	if !existe {
		return "", errNotFound
	}

	return valor, nil
}
~~~

---
## Uso de Defer
En el lenguaje de programación Go, la palabra clave `defer` se utiliza para posponer la ejecución de una función hasta que la función que la contiene haya finalizado. Esto se logra agregando la palabra clave `defer` antes de la llamada a la función.

Cuando se utiliza `defer`, la función especificada se agrega a una pila de ejecución de funciones diferidas. Una vez que la función que contiene el `defer` ha finalizado, las funciones diferidas se ejecutan en el orden inverso al que se agregaron a la pila.

El uso de `defer` es útil en situaciones donde se necesita asegurar que ciertas operaciones se realicen antes de que una función se complete, independientemente de cómo salga esa función (por una excepción, un error o simplemente una finalización normal). Al posponer la ejecución de una función utilizando `defer`, se puede garantizar que se realice una limpieza o cierre adecuado de recursos, incluso si ocurren excepciones o errores en el camino.

Aquí tienes un ejemplo de cómo se puede usar `defer` en Go:
~~~go
defer fmt.Println(3)
fmt.Println(1)
fmt.Println(2)
~~~
En este ejemplo, la línea `fmt.Println(3)` se pospone hasta que la función main haya finalizado, por lo que se imprimirá como la última línea de salida. Las líneas antes y después del defer se ejecutarán en su orden original.

~~~go
defer fmt.Println(3)
defer fmt.Println(2)
defer fmt.Println(1)
~~~
El código que has proporcionado utiliza la palabra clave `defer` en tres llamadas consecutivas a la función `fmt.Println()` en orden descendente: primero se pospone la impresión del número 3, luego el número 2 y finalmente el número 1.

Cuando se utiliza `defer`, las funciones diferidas se agregan a una pila de ejecución y se ejecutan en orden inverso una vez que la función que las contiene ha finalizado. En este caso, las llamadas a `fmt.Println()` se agregarán a la pila en el orden inverso al que aparecen en el código.

### Ejemplo de uso de defer
Este código en Go muestra un ejemplo de creación de un archivo llamado "hola.txt" y escribir en él el texto "Hola, Alex Roel":

~~~go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hola, Alex Roel"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
~~~
- Se importan los paquetes fmt y os, que son necesarios para mostrar mensajes de error y realizar operaciones relacionadas con el sistema operativo, respectivamente.

- Dentro de main(), se utiliza `os.Create("hola.txt")` para crear un nuevo archivo llamado "hola.txt". Esta función devuelve un puntero a un objeto de tipo `os.File`, que representa el archivo creado, y un posible error (err).

- Se verifica si ocurrió algún error al crear el archivo. Si `err` no es `nil` (es decir, si hay un error), se muestra el mensaje de error utilizando `fmt.Println()` y se finaliza la ejecución de la función `main()` utilizando `return`. En este caso, no se cerrará el archivo creado.

- Si no hubo errores al crear el archivo, se utiliza `defer` para posponer la ejecución de `file.Close()`. Esto garantiza que el archivo se cerrará antes de que la función main() termine, independientemente de cómo se salga de ella (por ejemplo, si se encuentra un error).

- A continuación, se utiliza `file.Write([]byte("Hola, Alex Roel"))` para escribir el texto "`Hola, Alex Roel"` en el archivo. El método `Write()` acepta un `slice` de `bytes ([]byte)` como argumento, por lo que se utiliza `[]byte("Hola, Alex Roel")` para convertir el texto en un slice de bytes antes de escribirlo en el archivo.

Se verifica si ocurrió algún error al escribir en el archivo. Si `err` no es `nil`, se muestra el mensaje de error utilizando `fmt.Println()` y se finaliza la ejecución de la función main() utilizando return.

---
## Uso de panic y recover

En Go, las palabras clave panic y recover se utilizan para manejar situaciones excepcionales o errores graves que ocurren durante la ejecución de un programa. A continuación, explicaré brevemente cada una de estas palabras clave:

### Panic 
La palabra clave `panic` se utiliza para generar una interrupción inmediata en la ejecución normal del programa. Cuando se produce un `panic`, el programa se detiene y comienza a desenrollar (unwind) la pila de llamadas, ejecutando cualquier función diferida (defer) en el camino. El resultado final es la terminación del programa y la impresión de un mensaje de error.

Un `panic` se puede provocar intencionalmente mediante la función `panic()` o puede ocurrir automáticamente en respuesta a una situación excepcional, como un índice fuera de rango en un slice o una división por cero. En cualquier caso, cuando se produce un `panic`, se detiene la ejecución normal y se inicia el proceso de desenrollado de la pila.

~~~go
package main

import "fmt"

func main() {
	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(124, 8)
}

func divide(dividend, divisor int) {
	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("🤕 no puedes dividir por cero")
	}
}
~~~

### Recover
La palabra clave `recover` se utiliza en conjunto con `defer` para capturar y manejar un `panic`. Cuando se coloca `recover` en una función diferida (defer), se capturará cualquier `panic` que ocurra en esa función o en las funciones llamadas desde ella. Esto permite que el programa recupere el control y realice acciones adicionales antes de que la ejecución finalice.

Para utilizar `recover`, se debe llamar a la función `recover()` dentro de un bloque `defer`. Si no hay ningún `panic` activo en el momento de la llamada a `recover()`, esta devuelve `nil`. Si hay un `panic` activo, `recover()` detiene el desenrollado de la pila, devuelve el valor pasado al `panic` y la ejecución continúa normalmente.

Es importante destacar que `recover` solo puede capturar `panic` dentro de la misma goroutine. Si se produce un `panic` en una goroutine distinta, no se puede capturar y propagará a la goroutine principal, lo que podría finalizar el programa.

~~~go
func divide(dividend, divisor int) {
	/* Se utiliza defer para posponer la ejecución de una función anónima hasta que la función que la contiene haya finalizado.
	La función anónima tiene como objetivo capturar cualquier panic que ocurra en el bloque de código que la rodea y realizar una acción de recuperación.
	Si se produce un panic, el valor recuperado se almacena en 'r' y se verifica si no es nulo.
	Si 'r' no es nulo, significa que se produjo un panic y se imprime su valor utilizando fmt.Println().
	Esta construcción es útil para manejar errores y realizar tareas de limpieza antes de que la función finalice su ejecución. */
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r) // Imprimir el valor recuperado en caso de un panic
		}
	}()

	validateZero(divisor)
	fmt.Println(dividend / divisor)
}
~~~

El uso de `panic` y `recover` se reserva para casos excepcionales y errores graves, y no se recomienda su uso para el control de flujo normal del programa. Se deben utilizar principalmente para manejar situaciones inesperadas o para realizar tareas de limpieza y recuperación antes de finalizar la ejecución del programa.

--- 
## Registro de errores
Los registros desempeñan un papel importante en un programa, ya que se convierten en un origen de información que se puede comprobar cuando se produce algún problema. Normalmente, cuando se produce un error, los usuarios finales solo ven un mensaje que indica un problema en el programa. Desde el punto de vista del desarrollador, necesitamos más información que un mensaje de error simple. Esto se debe principalmente a que queremos reproducir el problema para escribir una corrección adecuada. En este módulo, aprenderá cómo funciona el registro en Go. También aprenderá algunas prácticas que siempre debe implementar.

### Paquete log
El paquete `log` en Go es parte del paquete estándar y se utiliza para el registro básico de mensajes en aplicaciones. Proporciona funcionalidades simples para imprimir mensajes de registro en la salida estándar y en la salida de errores estándar.

A continuación, se proporciona una breve descripción y algunos ejemplos de uso del paquete `log`:

~~~go
log.Print("Este es un mensaje de registro")
log.Println("Este es otro mensaje de registro")
~~~

El paquete log proporciona varias funciones para imprimir mensajes de registro en la salida estándar. `Print()` se utiliza para imprimir un mensaje sin agregar un salto de línea al final, mientras que `Println()` agrega automáticamente un salto de línea al final del mensaje.

Puede usar la función `log.Fatal() `para registrar un error y finalizar el programa como si hubiera usado `os.Exit(1)`. Para probarlo, utilicemos este fragmento de código:

~~~go
log.Fatal("¡Oye, soy un registro de errores!")
fmt.Print("¿Puedes verme?")
~~~

Se obtiene un comportamiento similar al usar la función `log.Panic()`, que también llama a la función `panic()` de la siguiente manera:

~~~go
log.Panic("¡Oye, soy un registro de errores!")
fmt.Print("¿Puedes verme?")
~~~

Otra función esencial es `log.SetPrefix()`. Puede utilizarla para agregar un prefijo a los mensajes de registro del programa. Por ejemplo, podría utilizar este fragmento de código:

~~~go
log.SetPrefix("main(): ")
log.Print("¡Oye, soy un Log!")
log.Fatal("¡Oye, soy un registro de errores!")
~~~

### Registro en un archivo
Además de imprimir registros en la consola, es posible que quiera enviar registros a un archivo para poder procesarlos más adelante o en tiempo real.

¿Por qué quiere enviar los registros a un archivo? En primer lugar, es posible que quiera ocultar información específica a los usuarios finales. Puede que no les interesen o que se exponga información confidencial. Si tiene registros en archivos, puede centralizar todos los registros en una sola ubicación y correlacionarlos con otros eventos. Este patrón es típico: tener aplicaciones distribuidas que puedan ser efímeras, como los contenedores.

Vamos a usar el código siguiente para probar el envío de registros a un archivo:

~~~go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("¡Oye, soy un Log!")
}
~~~


El código proporcionado muestra cómo utilizar el paquete log en Go para redirigir la salida de registro a un archivo y registrar mensajes en dicho archivo. 

- `os.O_CREATE|os.O_APPEND|os.O_WRONLY`: Estas son las banderas de modo utilizadas para el archivo. En este caso, se combinan tres banderas utilizando el operador OR (|):
- `os.O_CREATE`: Crea el archivo si no existe.
- `os.O_APPEND`: Abre el archivo en modo adjunto, lo que significa que los nuevos datos se agregarán al final del archivo.
- `os.O_WRONLY`: Abre el archivo en modo de solo escritura.
- `0644`: Es el valor octal utilizado para los permisos del archivo. En este caso, 0644 significa que el archivo tendrá permisos de lectura y escritura para el propietario y solo permisos de lectura para los demás.

---
## Proyecto de sección: Gestor de contactos
Desarrolla una aplicación que permita a los usuarios guardar y administrar contactos en un archivo. Asegúrate de manejar los errores que puedan ocurrir al agregar y mostrar contactos.

Este es un ejemplo básico de un gestor de contactos que utiliza archivos para almacenar los contactos en formato JSON. El programa ofrece las opciones de agregar un nuevo contacto, mostrar todos los contactos y salir. Los contactos se guardan en el archivo "contacts.json" y se cargan al inicio del programa.

Ten en cuenta que este ejemplo solo aborda la funcionalidad básica del gestor de contactos y no incluye todas las validaciones de entrada y manejo de errores que podrías implementar. Puedes mejorar y expandir este código según tus necesidades y requisitos específicos.

Recuerda crear el archivo "contacts.json" en el mismo directorio antes de ejecutar el programa, o asegúrate de tener permisos de escritura en el directorio para crearlo automáticamente.

~~~go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contactos
type Contact struct {
	Name  string
	Email string
	Phone string
}

func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	// Crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar opciones al usuario
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opción: ")
		// Leer la opción del usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		// Manejar la opción del usuario
		switch option {
		case 1:
			// Ingresar y crear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')
			// Agregar un contacto a Slice
			contacts = append(contacts, c)

			// Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}

		case 2:
			// Mostrar todos los contactos
			fmt.Println("=======================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=======================================")
		case 3:
			// Salir del programa
			return
		default:
			fmt.Println("Opción no válida")
		}

	}
}

// Guarda contatos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// Carga contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}
~~~

---
## Resumen
En la sección del curso de Go dedicada al manejo de errores, exploramos diferentes aspectos del manejo de errores en Go y cómo aplicarlos de manera efectiva en nuestras aplicaciones.

Comenzamos aprendiendo sobre el manejo de errores en Go, donde vimos cómo utilizar la sintaxis if err != nil para verificar y manejar errores de forma concisa y eficiente. También exploramos las mejores prácticas para devolver y manejar errores adecuadamente en nuestras funciones.

Luego profundizamos en un ejemplo práctico de manejo de errores en Go, donde analizamos cómo implementar el manejo de errores en una aplicación real. Exploramos técnicas como el uso de variables de error, el registro de errores y la propagación adecuada de errores en diferentes capas de nuestra aplicación.

Uno de los aspectos más interesantes del manejo de errores en Go es el uso de la declaración defer. Aprendimos cómo usar defer para asegurarnos de que ciertas operaciones se realicen antes de salir de una función, incluso en caso de error. Vimos ejemplos prácticos que demostraron cómo defer puede facilitar el manejo de tareas como la liberación de recursos o el cierre adecuado de archivos.

También exploramos el uso de panic y recover. Aprendimos cómo utilizar estas dos funciones para manejar situaciones excepcionales y errores graves en nuestras aplicaciones. Vimos cómo panic puede interrumpir el flujo normal del programa y cómo recover nos permite recuperarnos de dichos errores y continuar la ejecución de manera controlada.

En el proyecto de esta sección, creamos un gestor de contactos utilizando Go. Aplicamos todos los conceptos que aprendimos sobre el manejo de errores en Go para implementar una funcionalidad robusta para agregar, mostrar y almacenar contactos, asegurándonos de manejar los posibles errores que pudieran ocurrir durante el proceso.

Concluimos esta sección con una sólida comprensión del manejo de errores en Go. Ahora estamos preparados para crear aplicaciones confiables y resistentes que manejen adecuadamente situaciones inesperadas. 
