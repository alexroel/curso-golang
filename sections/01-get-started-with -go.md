# Primeros pasos con Go

1. [¿Qué es Go?](#¿Qué-es-Go?)
2. [Herramientas de trabajo](#Herramientas-de-trabajo)
3. [Hola Mundo con Go](#Hola-Mundo-con-Go)
4. [Comandos de Go](#Comandos-de-Go)

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
    - Descarga e Instala Go: 

### Editor e IDES para Go 
    - Visual Studio Code 
    - Vim 
    - Sublime Text 
    - Golang (IDE)

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

- `go help`: Podemos ver todos los comandos de go con su respercivo descripción.
- `go run app.go`: Compilar y ejecutar el programa Go 
- `go build app.go`: Compila paquetes y dependencias
- `go doc fmt`: Muestra la documentación del paquete o símbolo
- `go env`: Imprime la información del entorno Go
- `go get`: Agregue dependencias al módulo actual e instálelas
- `go mod`: Mantenimiento del módulo
