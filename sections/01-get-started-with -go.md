# Primeros pasos con Go
 
1. [¿Qué es Go?](#¿Qué-es-Go?)
2. [Herramientas de trabajo](#Herramientas-de-trabajo)
3. [Hola Mundo con Go](#Hola-Mundo-con-Go)
4. [Mi primer función](#mi-primer-función)
5. [Comandos de Go](#Comandos-de-Go)
6. [Comentarios en Go](#Comentarios-en-Go)
 
---
## ¿Qué es Go?
 
Go, también conocido como Golang, es un lenguaje de programación de código abierto diseñado para ser eficiente, seguro y fácil de usar. Fue desarrollado por Google y lanzado en el año 2009. Algunas de las principales características de Go son:

- Rendimiento: Go se diseñó para ser rápido y eficiente en el uso de recursos. Go tiene una gestión de memoria muy eficiente y un recolector de basura automático que ayuda a prevenir fugas de memoria.

- Confiabilidad: Go fue diseñado para ser un lenguaje seguro y confiable. Go incluye características como la comprobación de errores en tiempo de compilación, la verificación de tipos y la detección de condiciones de carrera en tiempo de ejecución.

- Concurrencia: Go tiene concurrencia nativa, lo que permite a los desarrolladores crear programas que pueden manejar múltiples tareas simultáneamente sin problemas. Go utiliza goroutines y canales para lograr la concurrencia de manera eficiente.

- Legibilidad: Go tiene una sintaxis clara y legible que facilita la lectura y el mantenimiento del código. Go enfatiza la simplicidad y la facilidad de uso, lo que hace que sea fácil para los desarrolladores nuevos aprender el lenguaje.

- Escalabilidad: Go es muy escalable y puede manejar grandes aplicaciones empresariales sin problemas. Go cuenta con herramientas integradas que permiten a los desarrolladores crear aplicaciones eficientes y escalables.

- Portabilidad: Go se puede compilar y ejecutar en una variedad de plataformas, lo que lo hace muy portátil. Go también tiene una amplia variedad de bibliotecas y herramientas disponibles para su uso, lo que lo hace muy flexible y adaptable.

En resumen, Go es un lenguaje de programación moderno, confiable y eficiente, que se utiliza ampliamente en la creación de aplicaciones empresariales escalables y en la creación de infraestructuras de software de alto rendimiento.
 
---
## Herramientas de trabajo
Para aprender Go podemos utilizar el Playground de Go, pero cuando empecemos a crear aplicaciones reales con Go, necesitaremos las herramientas de desarrollo de Go y un editor de código o IDE para Go.
 
### Instalar Go
- Descarga e Instala Go: https://go.dev/doc/install
 
### Editor e IDES para Go
- Visual Studio Code: https://code.visualstudio.com/download
- Vim: https://www.vim.org/
- Sublime Text:  https://www.sublimetext.com/
- Golang (IDE): https://www.jetbrains.com/go/
 
### Instalar VSCode y Extensiones
- Go (Extensión)
- Extensiones Opcionales
    - Material Icon Theme
    - Atom One Dark Theme
    - Thunder Client
---
## Hola Mundo con Go
Para escribir un programa "Hola Mundo" en Go, sigue los siguientes pasos:

- Abre un editor de texto y crea un nuevo archivo llamado "hola.go".

Escribe el siguiente código:
 
~~~go
package main
 
import "fmt"
 
func main() {
    fmt.Println("¡Hola Mundo de Go!")
}
~~~
 
- `package main`: Esta línea indica que este archivo es el paquete principal del programa. En Go, cada programa debe tener un paquete principal llamado "main".

- `import "fmt"`: La palabra clave "import" se utiliza en Go para importar paquetes. En este caso, estamos importando el paquete "fmt", que es uno de los paquetes estándar de Go y se utiliza para imprimir texto en la consola.
 
- `func main(){}`: Esta es la función principal del programa. Es el punto de entrada del programa y se ejecutará automáticamente cuando se inicie el programa. La función "main" es obligatoria en todos los programas de Go.

Dentro de la función "main", estamos llamando a la función "Println" del paquete "fmt" para imprimir el texto "Hola Mundo" en la consola.
 
- `fmt.Println("Hola")`: La función "Println" toma un argumento de tipo "string" y lo imprime en la consola seguido de un salto de línea.

Al ejecutar el programa, verás la salida "Hola Mundo" en la consola.

---
## Mi primer función
¡Excelente! Aquí hay un ejemplo de cómo escribir una función simple en Go:

~~~go
package main

import "fmt"

func miFuncion() {
    fmt.Println("Hola desde mi función")
}

func main() {

    fmt.Println("¡Hola Mundo de Go!")
    miFuncion()
}
~~~

En este ejemplo, hemos creado una función llamada "miFuncion" que imprime "Hola desde mi función" en la consola utilizando la función "Println" del paquete "fmt".

Dentro de la función principal "main", estamos llamando a la función "miFuncion" utilizando el nombre de la función seguido de paréntesis vacíos. Esto invoca la función y ejecuta el código dentro de ella, imprimiendo "Hola desde mi función" en la consola.

¡Eso es todo! Ahora tienes una función simple en Go que puedes llamar desde cualquier lugar en tu programa.

---
## Comandos de Go
Los comandos de Go no va ayudar para ejecutar aplicaciones que vamos a crear como también compilar aplicaciones de Go, pero también podemos hacer muchas cosas con los comandos de Go, como instalar paquetes de  terceros, crear manejador de módulos y muchas cosas  más.

Aquí hay una lista de algunos de los comandos más comunes de Go:
 
Doc de comandos de Go: https://pkg.go.dev/cmd/go
 
- `go help`: Podemos ver todos los comandos de go con su respectivo descripción.
- `go run`: compila y ejecuta un archivo Go. Por ejemplo, "go run archivo.go" compilará y ejecutará el archivo "archivo.go".
- `go build`: compila un archivo Go y lo convierte en un archivo ejecutable. Por ejemplo, "go build archivo.go" compilará el archivo "archivo.go" y creará un archivo ejecutable llamado "archivo".
- `go fmt`: formatea el código fuente en un archivo Go. Por ejemplo, "go fmt archivo.go" formateará el código en el archivo "archivo.go".
- `go get`: descarga y compila paquetes y dependencias. Por ejemplo, "go get github.com/nombre_usuario/paquete" descargará y compilará el paquete de GitHub "paquete" del usuario "nombre_usuario".
- `go vet`: analiza el código fuente en busca de errores comunes. Por ejemplo, "go vet archivo.go" analizará el código en el archivo "archivo.go" en busca de errores comunes.
- `go test`: ejecuta las pruebas en un paquete. Por ejemplo, "go test paquete" ejecutará las pruebas en el paquete "paquete".
- `go doc`: muestra la documentación de un paquete. Por ejemplo, "go doc paquete" mostrará la documentación del paquete "paquete".

Estos son solo algunos de los comandos más comunes de Go. La documentación oficial de Go proporciona una lista completa de todos los comandos y opciones disponibles.
 
---
## Comentarios en Go

En Go, los comentarios se utilizan para documentar el código y hacerlo más fácil de entender para otros desarrolladores. Hay dos tipos de comentarios en Go:

- Comentarios de línea: se utilizan para hacer comentarios breves en una sola línea. Se pueden colocar en cualquier lugar de la línea, antes o después del código. Se indican con dos barras inclinadas (//).
- Comentarios de bloque: se utilizan para hacer comentarios más largos que ocupan varias líneas. Se pueden colocar en cualquier lugar dentro del bloque de código y se indican con una barra inclinada seguida de un asterisco (/) y terminan con un asterisco seguido de una barra inclinada (/).

~~~go
// El paquete "main" es el punto de entrada de la aplicación
package main

/* Importamos el paquete "fmt", q
ue proporciona funciones para imprimir 
y formatear texto
*/
import "fmt"

// Definimos la función "miFuncion"
func miFuncion() {
    fmt.Println("Hola desde mi función")
}

// Función principal "main" de la aplicación
func main() {

    /* Imprimimos un mensaje en la 
    consola utilizando la función 
    "Println" del paquete "fmt"
    */
    fmt.Println("¡Hola Mundo de Go!")
    
    // Llamamos a la función "miFuncion"
    miFuncion()
}
~~~

Es importante utilizar comentarios para documentar tu código, especialmente para explicar partes del código que puedan ser confusas o difíciles de entender. Un buen código es aquel que puede ser comprendido y mantenido fácilmente por otros desarrolladores, y los comentarios son una herramienta valiosa para lograr ese objetivo.

 

 
