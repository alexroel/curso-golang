# Primeros pasos con Go

1. [¿Qué es Go?](#¿Qué-es-Go?)
2. [Herramientas de trabajo](#Herramientas-de-trabajo)
3. [Hola Mundo con Go](#Hola-Mundo-con-Go)
4. [Comandos de Go](#Comandos-de-Go)
5. [Comentarios en Go](#Comentarios-en-Goo)

---
## ¿Qué es Go?

Cree software rápido, confiable y eficiente a escala
- Go es un lenguaje de programación de código abierto compatible con Google
- Fácil de aprender y comenzar
- Simultaneidad incorporada y una biblioteca estándar robusta
- Ecosistema creciente de socios, comunidades y herramientas

---
## Herramientas de trabajo
Para aprender Go podemos utiliziar el Playground de Go, pero cuando empezemos a crear aplicaciones reales con Go, necesitaremos las herraminestas de dasrrollo de Go y un editor de codigo o IDE para Go. 

### Instalar Go 
- Descarga e Instala Go: https://go.dev/doc/install

### Editor e IDES para Go 
- Visual Studio Code: https://code.visualstudio.com/download
- Vim: https://www.vim.org/
- Sublime Text:  https://www.sublimetext.com/
- Golang (IDE): https://www.jetbrains.com/go/

---
## Hola Mundo con Go
Crear Hola Mundo con Go es muy simple como se muestra en el siguiente código. 

~~~go
package main

import "fmt"

func main() {
	fmt.Println("¡Hola Mundo de Go!")
}
~~~

- `package main`: Todos los archivos y aplicaciones de Go pertenecen a un paquete o podemos indicar que pertenece al paquete principal.
- `import "fmt"`: Importar paque que usamos en nuestra aplicación. 

- `func main(){}`: Todas las aplicaciones en Go inician su ejecucion desde la funcion `main`. 

- `fmt.Println("Hola")`: El paquete `fmt` y su función `Println()` es para imprimer datos en la consolo. 

---
## Comandos de Go
Los comandos de Go no va ayudar para ejecutar aplicaciones que vamos a crear como tambien compilar aplicaciones de Go, pero tambien podemos hacer muchas cosa con los comandos de Go, como instalar paquetes de  terceros, crear manejador de modulos y muschas cosas  mas. 

Doc de comandos de Go: https://pkg.go.dev/cmd/go

- `go help`: Podemos ver todos los comandos de go con su respercivo descripción.
- `go run app.go`: Compilar y ejecutar el programa Go 
- `go build app.go`: Compila paquetes y dependencias
- `go doc fmt`: Muestra la documentación del paquete o símbolo
- `go env`: Imprime la información del entorno Go
- `go get`: Agregue dependencias al módulo actual e instálelas
- `go mod`: Mantenimiento del módulo

---
## Comentarios en G
- Un comentario es un texto que se ignora al ejecutarse.
- Los comentarios se pueden utilizar para explicar el código y hacerlo más legible.
- Los comentarios también se pueden usar para evitar la ejecución de código al probar un código alternativo.
- Go admite comentarios de una o varias líneas.
- Los comentarios de una sola línea comienzan con dos barras inclinadas `//`.
- Los comentarios de varias líneas comienzan con `/*` y terminan con `*/`.
