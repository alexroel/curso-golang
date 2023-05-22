# Manejo de paquetes y módulos 
 1. [Manejador de mòdulos](#manejador-de-mòdulos)
 2. [Paquetes en Go](#paquetes-en-go)
 3. [Manejo de error](#manejo-de-error)  
 4. [Saludo aleatorio](#saludo-aleatorio)
 5. [Subir a GitHub](#subir-a-github) 
 6. [Usar desde otro aplicación](#usar-desde-otro-aplicación)

---
## Manejador de mòdulos
El manejador de módulos en Go se encarga de la gestión de dependencias de los proyectos y se introdujo en la versión 1.11 de Go. Permite a los desarrolladores especificar las dependencias de sus proyectos y manejar fácilmente las actualizaciones de esas dependencias.

Aquí hay un ejemplo de cómo utilizar un módulo de terceros en Go:

Crear un nuevo proyecto y habilitar el manejo de módulos:
~~~bash
mkdir greetings
cd greetings
go mod init greetings
~~~

Agregar una dependencia de un módulo de terceros al proyecto. En este ejemplo, usaremos el paquete "go get rsc.io/quote" que es un framework de Go para el desarrollo web.

~~~bash
go get rsc.io/quote
~~~

Importar el paquete en el archivo principal de nuestro proyecto y utilizarlo:

~~~go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())
}
~~~

Cuando agregamos una dependencia de un módulo de terceros, Go automáticamente descarga el paquete y lo almacena en el directorio `go.mod`. También crea un archivo `go.su`m que contiene una lista de las sumas de verificación de cada archivo descargado, lo que garantiza la integridad de las dependencias.

---
## Paquetes en Go

En Go, un paquete es un conjunto de archivos que definen un conjunto de funciones, tipos y variables que pueden ser utilizados por otros programas. Los paquetes en Go proporcionan una forma de organizar el código y de separar la funcionalidad en componentes reutilizables.

Para crear un paquete en Go, se debe crear un directorio con el nombre del paquete y colocar los archivos que definen las funciones y tipos en ese directorio. El nombre del paquete debe ser el mismo que el nombre del directorio.

Por ejemplo, para crear un paquete llamado "geometry" que contiene una función llamada "geometry", se puede crear el siguiente directorio y archivo:

~~~bash
greeting/
    greeting.go
~~~

El archivo "greeting.go" podría contener el siguiente código:

~~~go
package greeting

import "fmt"

// Función de saludo por nombre
func Hello(name string) string {
	message := fmt.Sprintf("Hola, %v. Bienvenido!", name)
	return message
}
~~~

Para utilizar este paquete en otro programa, se debe importar el paquete en el archivo principal y luego llamar a la función "Hello": 

~~~go
package main

import (
	"fmt"
	"greetings/greeting"
)

func main() {

	// Saludo desde el paquete
	fmt.Println(greeting.Hello("Alex"))
}
~~~

---
## Manejo de error
Go proporciona un mecanismo de manejo de errores integrado que permite a los programadores controlar el flujo del programa cuando ocurren errores.

En Go, el manejo de errores se realiza utilizando el tipo de dato "error", que es una interfaz que tiene un solo método "Error()" que devuelve un mensaje de error.

Para manejar un error, se utiliza el valor de retorno del tipo de dato error y se verifica si su valor es nulo. Si es nulo, significa que no se ha producido ningún error y el programa puede continuar con su flujo normal. Si no es nulo, se ha producido un error y el programa debe manejarlo adecuadamente.

~~~go
package greeting

import (
	"errors"
	"fmt"
)

// Función de saludo por nombre
func Hello(name string) (string, error) {

	// Si name es vació
	if name == "" {
		return "", errors.New("nombre vacío")
	}

	message := fmt.Sprintf("Hola, %v. Bienvenido!", name)
	return message, nil
}
~~~

Este código define una función llamada "Hello" que toma un parámetro "name" de tipo string y devuelve una cadena de saludo personalizada y un error.

La función comienza verificando si el valor del parámetro "name" es una cadena vacía utilizando un condicional "if". Si el valor es una cadena vacía, la función devuelve una cadena vacía y un error creado con la función "errors.New()", que crea un nuevo valor de error con un mensaje de error personalizado. El mensaje de error en este caso es "nombre vacío".

~~~go
package main

import (
	"fmt"
	"greetings/greeting"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	var name string
	// Saludo desde el paquete
	if message, error := greeting.Hello(name); error != nil {
		log.Fatal(error)
	} else {
		fmt.Println(message)
	}

}
~~~

---
## Saludo aleatorio 
Para devolver un saludo aleatorio en Go, podemos utilizar una función que elija aleatoriamente una cadena de saludo de un conjunto de opciones predefinidas. Podemos implementar esto de la siguiente manera:

~~~go

// Retorna saludo aleatorio
func randomFormat() string {
	formats := []string{
		"Hola, %s. Bienvenido!",
		"Hola, %s. Ya eres Gopher",
		"Que bueno verte, %s!",
	}

	return formats[rand.Intn(len(formats))]
}
~~~

Este código define una función llamada "randomFormat()" que devuelve una cadena de formato de saludo aleatoria de un conjunto predefinido de opciones.

La función comienza definiendo un arreglo de cadenas llamado "formats" que contiene tres opciones de saludo diferentes, cada una con un marcador de posición para un nombre. Luego, la función utiliza la función "rand.Intn()" para elegir aleatoriamente un índice del arreglo "formats". La función "len()" se utiliza para determinar la longitud del arreglo, lo que nos permite asegurarnos de que el índice aleatorio seleccionado no exceda la longitud del arreglo.

Finalmente, la función devuelve la cadena de formato de saludo aleatoria seleccionada utilizando el índice aleatorio como índice del arreglo.

En resumen, esta función nos permite generar saludos aleatorios personalizados en Go utilizando una lista predefinida de opciones de saludo y una función de generación de números aleatorios. Esta implementación puede ser utilizada en conjunto con otras funciones para generar saludos aleatorios personalizados en Go.

---
## Subir a GitHub 
Para subir un módulo de Go a GitHub, puedes seguir los siguientes pasos:

1. Crea un repositorio vacío en GitHub.
2. Crea un directorio para tu módulo de Go. Este directorio debe ser el directorio principal de tu módulo y debe contener el archivo "go.mod" que define las dependencias del módulo.
3. Inicializa un repositorio de Git en el directorio de tu módulo y haz el primer commit:
4. Agrega el repositorio de GitHub como un control remoto:
5. Empuja tus cambios al repositorio de GitHub:
6. Ahora puedes agregar etiquetas a tu repositorio para indicar las versiones de tu módulo. 7. Para agregar una etiqueta para la versión 1.0.0, puedes ejecutar el siguiente comando:
    ~~~go
    git tag v1.0.0
    ~~~
8. Luego, debes enviar tus etiquetas al repositorio de GitHub:
    ~~~go
    git push --tags
    ~~~

Ahora tu módulo de Go está disponible en el repositorio de GitHub y puede ser utilizado por otros proyectos de Go como una dependencia.
Es importante asegurarse de que tu módulo de Go tenga una estructura adecuada y que siga las convenciones de nomenclatura de paquetes de Go. También es recomendable agregar documentación para tu módulo y sus funciones, para que otros desarrolladores puedan entender cómo utilizarlo.

---
## Usar desde otro aplicación
Para utilizar un módulo de Go que se encuentra en un repositorio en otro proyecto, puedes seguir los siguientes pasos:

1. Asegúrate de tener instalado Go en tu computadora.
2. Crea un nuevo proyecto de Go y utiliza el comando "go mod init" para inicializar el archivo "go.mod" del proyecto:

    ~~~bash
    mkdir mi-proyecto
    cd mi-proyecto
    go mod init mi-proyecto
    ~~~

3. Edita el archivo "go.mod" y agrega una línea que apunte al módulo en tu repositorio de GitHub. La línea debe seguir el siguiente formato:
    ~~~go
    require <url-del-repositorio>#<nombre-del-módulo> <versión>
    ~~~

    Por ejemplo, si tu módulo se encuentra en el repositorio "https://github.com/tu-usuario/mi-modulo" y tiene el nombre "mi-modulo", puedes agregar la siguiente línea al archivo "go.mod":

    ~~~go
    require github.com/tu-usuario/mi-modulo mi-modulo v1.0.0
    ~~~

4. Ejecuta el comando "go mod tidy" para descargar las dependencias de tu módulo y agregarlas al archivo "go.sum":
    ~~~bash 
    go mod tidy
    ~~~

5. Ahora puedes utilizar las funciones de tu módulo de Go en tu proyecto. Para hacerlo, simplemente importa el paquete en tu código:
    ~~~
    import "github.com/tu-usuario/mi-modulo/mi-paquete"
    ~~~
6. Luego, puedes llamar a las funciones de tu módulo como lo harías con cualquier otra función de Go:

Ten en cuenta que si tu módulo de Go tiene dependencias adicionales, también deberás agregarlas al archivo "go.mod" de tu proyecto y ejecutar "go mod tidy" para descargarlas.
