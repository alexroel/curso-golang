# Primeros pasos con Go
 
1. [Introdución](#introdución)
2. [¿Qué es y por que usar Go?](#¿qué-es-y-por-que-usar-go)
3. [¿Qué necesitamos?](#¿qué-necesitamos)
4. [Playground de Go](#playground-de-go)
5. [Instalar y configurar Go](#instalar-y-configurar-go)
6. [Hola mundo con Go](#hola-mundo-con-go)
7. [Paquetes de terceros](#paquetes-de-terceros)
8. [Resumen](#resumen)
 
---
## Introdución
¡Bienvenidos a la sección de Primeros Pasos con Go! En esta sección, vamos a introducirnos al mundo de Go, un lenguaje de programación que ha ganado popularidad en los últimos años debido a su simplicidad, eficiencia y versatilidad.

En primer lugar, vamos a hablar sobre qué es Go y por qué deberías considerar usarlo para tus proyectos. Hablaremos sobre su historia, su sintaxis y sus principales características, y cómo se compara con otros lenguajes de programación.

A continuación, exploraremos el playground de Go, una herramienta en línea que nos permitirá experimentar con el lenguaje de programación Go sin necesidad de instalar nada en nuestro equipo.

Después, veremos cómo instalar y configurar Go en tu sistema operativo. Veremos los requisitos necesarios y los pasos a seguir para que puedas comenzar a escribir código en Go lo antes posible.

Luego, vamos a crear nuestro primer programa en Go: el famoso "Hola Mundo". Veremos cómo escribir, compilar y ejecutar nuestro programa en Go, para que puedas familiarizarte con el proceso de desarrollo.

Por último, hablaremos sobre paquetes de terceros en Go. Veremos cómo utilizarlos, instalarlos y cómo pueden ayudarnos a ampliar las funcionalidades de nuestro programa.

---
## ¿Qué es y por que usar Go?
Go es un lenguaje de programación de código abierto que fue creado en el año 2009 por Google. En este video, vamos a conocer qué es Go, para qué se utiliza y cuáles son sus características principales.

### ¿Qué es Go?
Go es un lenguaje de programación compilado, lo que significa que el código fuente se convierte en un archivo ejecutable antes de ser ejecutado en una computadora. Fue diseñado para ser rápido y eficiente, especialmente para aplicaciones de red y servidores web. Go se considera un lenguaje de programación de alto nivel, lo que significa que tiene una sintaxis más fácil de leer y entender que otros lenguajes de programación.

### Caracteristicas de Go
Go tiene varias características que lo hacen único en comparación con otros lenguajes de programación. Una de ellas es su sistema de recolección de basura, que hace que el programador no tenga que preocuparse por la administración de memoria. Go también tiene un sistema de tipos estáticos que ayuda a evitar errores de programación. Además, Go es un lenguaje concurrente, lo que significa que puede manejar múltiples tareas simultáneamente, lo que lo hace ideal para aplicaciones de red y servidores web.

### Como funciona Go
Para crear aplicaciones en Go, primero necesitamos escribir nuestro código en un archivo con extensión .go. Luego, el código se compila en un archivo ejecutable utilizando el compilador de Go. Una vez que el archivo ejecutable está listo, podemos ejecutar nuestro programa en la computadora.

### Aplicaciones de Go
Go se utiliza en una variedad de aplicaciones, desde la creación de servidores web hasta la programación de sistemas operativos. Empresas como Google, Uber y Dropbox utilizan Go para construir sus aplicaciones y servicios. Go también se utiliza para construir herramientas de desarrollo, como el popular sistema de control de versiones Git.

---
## ¿Qué necesitamos?
Para aprender y crear aplicaciones con el lenguaje de programación Go, necesitarás los siguientes elementos:

- Conocimientos sobre programación: Aunque Go es conocido por su facilidad de aprendizaje, es útil tener conocimientos básicos de programación. Familiarizarse con conceptos como variables, tipos de datos, estructuras de control (como bucles y condicionales) y funciones te proporcionará una base sólida para aprender Go.

- [Editor de código o texto](https://go.dev/doc/editors): Puedes utilizar cualquier editor de código o texto para programar en Go. Algunos editores populares para trabajar con Go incluyen Visual Studio Code (VSCode), Vim, Emacs y GoLand. Estos editores a menudo tienen complementos o extensiones disponibles que facilitan el desarrollo en Go, como resaltado de sintaxis, autocompletado y depuración.

- Instalación de Go: Debes instalar el lenguaje de programación Go en tu computadora. Puedes descargar la última versión estable de Go desde el sitio web oficial de Go (https://golang.org). Sigue las instrucciones de instalación proporcionadas para tu sistema operativo específico.

- Uso de la terminal: Ser capaz de usar la terminal o línea de comandos es esencial para compilar y ejecutar programas escritos en Go. Debes aprender a navegar por los directorios, ejecutar comandos de compilación y ejecución, y gestionar tu entorno de desarrollo desde la línea de comandos.

Además de estos elementos básicos, también es beneficioso tener acceso a recursos de aprendizaje, como la documentación oficial de Go (https://golang.org/doc/), tutoriales en línea, libros y comunidades de programadores de Go. Estos recursos pueden ayudarte a comprender mejor el lenguaje y sus características, y a resolver dudas o problemas que puedan surgir durante tu aprendizaje.

Además de los elementos mencionados anteriormente, también es útil familiarizarse con Git y GitHub para aprender y crear aplicaciones con Go. Aquí hay una descripción de estos conceptos:

- Git: Git es un sistema de control de versiones distribuido que te permite realizar un seguimiento de los cambios en tu código a lo largo del tiempo. Con Git, puedes crear ramas para trabajar en nuevas características o solucionar problemas sin afectar la versión principal de tu código. También facilita la colaboración en proyectos de desarrollo de software, ya que varios desarrolladores pueden trabajar en paralelo y combinar sus cambios de manera eficiente.

- GitHub: GitHub es una plataforma de alojamiento de repositorios basada en la nube que utiliza Git. Permite a los desarrolladores almacenar y compartir su código fuente, colaborar en proyectos con otros desarrolladores y realizar un seguimiento de los cambios en el código. En GitHub, puedes crear repositorios públicos o privados para tus proyectos y utilizar diversas funciones, como solicitudes de extracción (pull requests) y problemas (issues), para facilitar la colaboración y la comunicación con otros miembros del equipo.

---
## Playground de Go
¡Bienvenido al Playground de Go! Este es un ambiente en línea donde puedes escribir, compilar y ejecutar código en Go sin tener que instalar nada en tu computadora.

El Playground de Go es una herramienta muy útil para experimentar con el lenguaje de programación y para compartir código con otros programadores. Es muy fácil de usar y te permite probar diferentes ideas y soluciones sin preocuparte por la configuración de tu entorno de desarrollo.

Para empezar a utilizar el Playground de Go, simplemente dirígete a la página oficial de Go Playground en tu navegador web. Allí verás un editor de código en el que puedes escribir tu programa en Go.

Una vez que hayas escrito tu programa, puedes hacer clic en el botón "Run" para compilar y ejecutar el código. Si el programa no tiene errores, el Playground de Go mostrará el resultado de la ejecución en la parte inferior de la pantalla.

El Playground de Go también te permite compartir tu código con otros programadores a través de un enlace. Simplemente haz clic en el botón "Share" para obtener un enlace único que puedes compartir con tus colegas.

Además, el Playground de Go tiene una serie de características útiles, como la capacidad de guardar y cargar programas, la integración con GitHub, y la posibilidad de ver el resultado del programa en formato HTML.

### Limitaciones de Playground de Go

El Playground de Go es una herramienta muy útil para probar y compartir código en Go, pero también tiene algunas limitaciones. A continuación, se presentan algunas cosas que no se pueden hacer con el Playground de Go:

- No se puede interactuar con el sistema operativo: el Playground de Go se ejecuta en un entorno aislado en línea y no tiene acceso al sistema operativo en el que se ejecuta. Por lo tanto, no puedes realizar operaciones de E/S en el sistema de archivos, como leer o escribir archivos.

- No se pueden instalar paquetes externos: el Playground de Go viene con un conjunto predefinido de paquetes estándar, pero no permite instalar paquetes externos. Si necesitas utilizar un paquete que no está incluido en el conjunto estándar, deberás instalarlo en tu propia computadora.

- No se pueden ejecutar programas que requieren entrada de usuario: debido a que el Playground de Go se ejecuta en un entorno en línea, no es posible interactuar con el usuario para solicitar entrada en tiempo de ejecución. Por lo tanto, cualquier programa que requiera entrada del usuario no se puede ejecutar en el Playground de Go.

---
## Instalar y configurar Go
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
    - Monokai Pro
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

### Comentarios en Go

En Go, los comentarios se utilizan para documentar el código y hacerlo más fácil de entender para otros desarrolladores. Hay dos tipos de comentarios en Go:

- Comentarios de línea: se utilizan para hacer comentarios breves en una sola línea. Se pueden colocar en cualquier lugar de la línea, antes o después del código. Se indican con dos barras inclinadas (//).
- Comentarios de bloque: se utilizan para hacer comentarios más largos que ocupan varias líneas. Se pueden colocar en cualquier lugar dentro del bloque de código y se indican con una barra inclinada seguida de un asterisco (/) y terminan con un asterisco seguido de una barra inclinada (/).

Es importante utilizar comentarios para documentar tu código, especialmente para explicar partes del código que puedan ser confusas o difíciles de entender. Un buen código es aquel que puede ser comprendido y mantenido fácilmente por otros desarrolladores, y los comentarios son una herramienta valiosa para lograr ese objetivo.

---
## Paquetes de terceros
En Go, los paquetes de terceros son paquetes que no están incluidos en la biblioteca estándar de Go y que han sido desarrollados y mantenidos por la comunidad de desarrolladores de Go. Estos paquetes pueden ser utilizados para ampliar las funcionalidades de tu programa o para resolver problemas específicos.

### Manejador de modulos 

En Go, el manejador de módulos es una herramienta que te permite gestionar las dependencias de tu proyecto. Los módulos son colecciones de paquetes de Go que se agrupan y versionan juntos.

Para utilizar el manejador de módulos en tu proyecto de Go, primero debes inicializar un nuevo módulo. Esto se hace ejecutando el siguiente comando en la línea de comandos en el directorio raíz de tu proyecto:

~~~go
go mod init nombre_del_modulo
~~~

El archivo `go.mod` es un archivo que se utiliza para definir y gestionar los módulos y sus dependencias en un proyecto de Go. Es parte del sistema de manejo de módulos de Go, que se introdujo en Go 1.11 y se volvió obligatorio en Go 1.16.

El archivo `go.sum` es un archivo que se utiliza para registrar las sumas de verificación de los módulos y sus dependencias en un proyecto de Go. 

### Instalar un paquete externo 
`rsc.io/quote` es un paquete de terceros para el lenguaje de programación Go que proporciona una serie de citas famosas. Este paquete fue creado por Russ Cox, uno de los desarrolladores principales de Go.

Para utilizar el paquete rsc.io/quote en tu programa de Go, primero debes instalarlo en tu sistema utilizando el comando go get. Puedes hacerlo ejecutando el siguiente comando en la línea de comandos:

~~~go
go get rsc.io/quote
~~~

Después de instalar el paquete, puedes utilizarlo en tu programa de Go importando el paquete en tu código. Por ejemplo, si quisieras imprimir la cita "Hello, world" en tu programa, podrías agregar el siguiente código a tu archivo Go:

~~~go
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Hello())
}
~~~

`rsc.io/quote` también proporciona otras funciones para imprimir citas famosas, como `Go() `para imprimir una cita relacionada con el lenguaje de programación Go, o `Opt()` para imprimir una cita optimista. Puedes consultar la documentación del paquete para obtener más información sobre estas funciones y cómo utilizarlas.

---
## Resumen 

En resumen, en esta sección se cubrieron los siguientes temas:

- ¿Qué es y por qué usar Go? Go es un lenguaje de programación de código abierto desarrollado por Google que es rápido, eficiente, seguro y fácil de usar. Se utiliza comúnmente para construir aplicaciones de servidor, herramientas de línea de comandos y aplicaciones de alto rendimiento.
- Playground de Go: una herramienta en línea que permite a los usuarios escribir, ejecutar y compartir código de Go sin necesidad de instalar nada en su propia computadora.
- Instalar y configurar Go: se explicó cómo descargar e instalar Go en una computadora, y cómo configurar la variable de entorno GOPATH.
- Hola mundo con Go: se mostró cómo escribir y ejecutar un programa simple de "Hola mundo" en Go.
- Paquetes de terceros: se explicó cómo agregar paquetes de terceros a un proyecto de Go utilizando el comando go get, y cómo administrar las dependencias del proyecto utilizando los archivos go.mod y go.sum.

En conjunto, estos temas proporcionan una introducción útil al desarrollo de aplicaciones en Go y permiten a los programadores comenzar a escribir y ejecutar programas en Go en su propia computadora.


 
