# Crear un módulo 

1. [Introducción](#introducción)
2. [Iniciar un módulo que otros puedan utilizar](#iniciar-un-módulo-que-otros-puedan-utilizar)
3. [Llamar a tu código desde otro módulo](#llamar-a-tu-código-desde-otro-módulo)
4. [Devolver y manejar un error](#devolver-y-manejar-un-error)
5. [Devolver un saludo aleatorio](#devolver-un-saludo-aleatorio)
6. [Agregar una prueba](#agregar-una-prueba)
7. [Compilar e instalar la aplicación](#compilar-e-instalar-la-aplicación)
8. [Desplegar a GitHub](#desplegar-a-github)
9. [Resumen](#resumen)

---
## Introducción 

¡Bienvenidos a la sección "Crear un módulo" del curso de Go!

En esta sección, aprenderemos cómo iniciar un módulo en Go para que otros desarrolladores puedan utilizarlo en sus proyectos. También exploraremos cómo llamar a nuestro código desde otro módulo y cómo manejar y devolver errores adecuadamente.

Además, descubriremos cómo añadir una divertida funcionalidad a nuestro módulo: devolver un saludo aleatorio. Esto le dará un toque especial a nuestras aplicaciones.

Por supuesto, no podemos olvidarnos de la importancia de las pruebas en el desarrollo de software. Aprenderemos cómo agregar pruebas a nuestro módulo para asegurarnos de que funciona correctamente en diferentes escenarios.

Una vez que hayamos construido nuestro módulo y probado su funcionalidad, aprenderemos cómo compilar e instalar nuestra aplicación. Esto nos permitirá utilizar nuestro módulo en cualquier proyecto sin problemas.

Finalmente, descubriremos cómo desplegar nuestro módulo en GitHub, una plataforma popular para compartir y colaborar en proyectos de código abierto.

¡Comencemos a explorar el emocionante mundo de la creación de módulos en Go y a compartir nuestro código con otros desarrolladores!

---
## Iniciar un módulo que otros puedan utilizar
Comienza creando un módulo en Go. En un módulo, se recopilan uno o más paquetes relacionados que contienen un conjunto de funciones discretas y útiles. Por ejemplo, podrías crear un módulo con paquetes que tengan funciones para realizar análisis financieros, de modo que otras personas que escriban aplicaciones financieras puedan utilizar tu trabajo. Para obtener más información sobre el desarrollo de módulos, consulta Desarrollo y publicación de módulos.

El código en Go se agrupa en paquetes, y los paquetes se agrupan en módulos. Tu módulo especifica las dependencias necesarias para ejecutar tu código, incluyendo la versión de Go y el conjunto de otros módulos que requiere.

A medida que agregas o mejoras la funcionalidad de tu módulo, publicas nuevas versiones del mismo. Los desarrolladores que escriben código que llama a funciones de tu módulo pueden importar los paquetes actualizados del módulo y probar con la nueva versión antes de utilizarla en producción.

- Abre una ventana de comandos y navega hasta tu directorio principal.
En Linux o Mac:
~~~bash
cd
~~~
En Windows:
~~~bash
cd %HOMEPATH%
~~~
- Crea un directorio llamado "greetings" para el código fuente de tu módulo Go.
Por ejemplo, desde tu directorio principal, usa los siguientes comandos:
~~~bash
mkdir greetings
cd greetings
~~~

- Inicia tu módulo utilizando el comando `go mod init`.
Ejecuta el comando `go mod init`, indicándole la ruta de tu módulo. Aquí, utiliza "example.com/greetings". Si publicas un módulo, esta debe ser una ruta desde la cual tu módulo pueda ser descargado por las herramientas de Go. Esa sería la ubicación de tu repositorio de código.

Para obtener más información sobre cómo nombrar tu módulo con una ruta de módulo, consulta Gestión de dependencias.

~~~bash
$ go mod init example.com/greetings
go: creando archivo go.mod: módulo example.com/greetings
~~~

El comando go mod init crea un archivo go.mod para rastrear las dependencias de tu código. Hasta ahora, el archivo solo incluye el nombre de tu módulo y la versión de Go que admite tu código. Pero a medida que agregues dependencias, el archivo go.mod listará las versiones en las que se basa tu código. Esto permite que las compilaciones sean reproducibles y te brinda un control directo sobre qué versiones de módulos utilizar.

- En tu editor de texto, crea un archivo llamado "greetings.go" en el que escribirás tu código y guárdalo.
- Pega el siguiente código en tu archivo "greetings.go" y guárdalo.

~~~go
import "fmt"

// Hello devuelve un saludo para la persona especificada.
func Hello(name string) string {
    // Devuelve un saludo que incluye el nombre en un mensaje.
    message := fmt.Sprintf("¡Hola, %v! ¡Bienvenido!", name)
    return message
}
~~~

La función Hello es una función que recibe un parámetro "name" de tipo string y devuelve un string. Esta función se utiliza para generar un saludo personalizado para la persona especificada.

Dentro de la función Hello, se utiliza la función fmt.Sprintf para formatear un mensaje de saludo. Se utiliza un verbo de formato "%v" para insertar el valor del parámetro "name" en el mensaje. Luego, se asigna el mensaje formateado a la variable "message".

Finalmente, se devuelve el mensaje formateado como resultado de la función Hello.

En resumen, este código importa el paquete "fmt" para utilizar la función Sprintf y crea una función Hello que genera un saludo personalizado utilizando el nombre proporcionado.

---
## Llamar a tu código desde otro módulo
En la sección anterior, creaste un módulo llamado greetings. En esta sección, escribirás código para hacer llamadas a la función Hello en el módulo que acabas de crear. Escribirás código que se puede ejecutar como una aplicación y que llama al código en el módulo greetings.

Crea un directorio hello para el código fuente de tu módulo Go. Aquí es donde escribirás el código del módulo que realizará la llamada.

Después de crear este directorio, deberías tener tanto un directorio hello como un directorio `greetings` en el mismo nivel en la jerarquía, de la siguiente manera:

~~~bash
<home>/
|-- greetings/
|-- hello/
~~~

Por ejemplo, si tu símbolo del sistema está en el directorio greetings, puedes usar los siguientes comandos:
~~~bash
cd ..
mkdir hello
cd hello
~~~

Habilita el seguimiento de dependencias para el código que estás a punto de escribir.

Para habilitar el seguimiento de dependencias en tu código, ejecuta el comando `go mod init`, proporcionándole el nombre del módulo en el que estará tu código.

Para este tutorial, utiliza example.com/hello como la ruta del módulo.
~~~bash
$ go mod init example.com/hello
go: creando nuevo go.mod: módulo example.com/hello
~~~

En tu editor de texto, en el directorio hello, crea un archivo en el que escribirás tu código y llámalo `hello.go`.

Escribe el código para llamar a la función Hello y luego imprimir el valor de retorno de la función.

Para hacer eso, copia y pega el siguiente código en hello.go:

~~~go
package main

import (
"fmt"
"example.com/greetings"
)

func main() {
    // Obtén un mensaje de saludo e imprímelo.
    mensaje := greetings.Hello("Gladys")
    fmt.Println(mensaje)
}
~~~

En este código:

Se declara un paquete main. En Go, el código que se ejecuta como una aplicación debe estar en un paquete main.
Se importan dos paquetes: example.com/greetings y el paquete fmt. Esto permite que tu código tenga acceso a las funciones de esos paquetes. Al importar example.com/greetings (el paquete contenido en el módulo que creaste anteriormente), tienes acceso a la función Hello. También se importa fmt, que contiene funciones para manejar texto de entrada y salida (como imprimir texto en la consola).
Se obtiene un saludo llamando a la función Hello del paquete greetings.
Se edita el módulo example.com/hello para usar el módulo example.com/greetings local.

Para su uso en producción, publicarías el módulo example.com/greetings desde su repositorio (con una ruta de módulo que refleje su ubicación publicada), donde las herramientas de Go podrían encontrarlo y descargarlo. Pero como aún no has publicado el módulo, debes adaptar el módulo example.com/hello para que pueda encontrar el código example.com/greetings en tu sistema de archivos local.

Para hacer eso, usa el comando `go mod edit` para editar el módulo example.com/hello y redirigir las herramientas de Go desde su ruta de módulo (donde no está el módulo) al directorio local (donde sí está).

Desde el símbolo del sistema en el directorio hello, ejecuta el siguiente comando:

~~~bash
go mod edit -replace example.com/greetings=../greetings
~~~

El comando especifica que example.com/greetings debe ser reemplazado por ../greetings a efectos de ubicar la dependencia. Después de ejecutar el comando, el archivo go.mod en el directorio hello debe incluir una directiva de reemplazo:

~~~go
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings
~~~

Desde el símbolo del sistema en el directorio hello, ejecuta el comando `go mod tidy` para sincronizar las dependencias del módulo example.com/hello, agregando aquellas requeridas por el código, pero que aún no se han rastreado en el módulo.
~~~bash
$ go mod tidy
go: se encontró example.com/greetings en example.com/greetings v0.0.0-00010101000000-000000000000
~~~
Después de que el comando se complete, el archivo `go.mod `del módulo `example.com/hello` debería lucir así:

~~~
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
~~~

El comando encontró el código local en el directorio greetings y luego agregó una directiva `require` para especificar que `example.com/hello` requiere `example.com/greetings`. Creaste esta dependencia cuando importaste el paquete greetings en hello.go.

El número que sigue a la ruta del módulo es un número de` pseudo-versión`: un número generado que se utiliza en lugar de un número de versión semántica (que el módulo aún no tiene).

Para hacer referencia a un módulo publicado, un archivo `go.mod `normalmente omitiría la directiva replace y utilizaría una directiva require con un número de versión etiquetado al final.

~~~
require example.com/greetings v1.1.0
~~~
Para obtener más información sobre los números de versión, consulta la numeración de versiones de módulos.

En el símbolo del sistema en el directorio hello, ejecuta tu código para confirmar que funciona.

~~~
$ go run .

¡Hola, Gladys! ¡Bienvenido!
~~~

¡Felicidades! Has escrito dos módulos funcionales.

En el próximo tema, agregarás manejo de errores.

---
## Devolver y manejar un error
El manejo de errores es una característica esencial de un código sólido. En esta sección, agregarás un poco de código para devolver un error desde el módulo greetings y luego manejarlo en el llamador.

En el archivo greetings/greetings.go, agrega el código resaltado a continuación.
No tiene sentido enviar un saludo si no sabes a quién saludar. Devuelve un error al llamador si el nombre está vacío. Copia el siguiente código en greetings.go y guarda el archivo.

~~~go
package greetings

import (
"errors"
"fmt"
)

// Hello devuelve un saludo para la persona especificada.
func Hello(name string) (string, error) {
    // Si no se proporcionó ningún nombre, devuelve un error con un mensaje.
    if name == "" {
        return "", errors.New("nombre vacío")
    }

    // Si se recibió un nombre, devuelve un valor que incluye el nombre
    // en un mensaje de saludo.
    message := fmt.Sprintf("¡Hola, %v! ¡Bienvenido!", name)
    return message, nil
}
~~~

En este código:

Cambiaste la función para que devuelva dos valores: una cadena y un error. El llamador verificará el segundo valor para ver si se produjo un error. (Cualquier función de Go puede devolver múltiples valores. Para obtener más información, consulta Effective Go).
Importaste el paquete errors de la biblioteca estándar de Go para poder utilizar su función errors.New.
Agregaste una declaración if para verificar una solicitud inválida (una cadena vacía donde debería estar el nombre) y devolver un error si la solicitud es inválida. La función errors.New devuelve un error con tu mensaje dentro.
Agregaste nil (que significa sin error) como segundo valor en la devolución exitosa. De esta manera, el llamador puede ver que la función tuvo éxito.

En tu archivo hello/hello.go, maneja el error ahora devuelto por la función Hello, junto con el valor no error.
Copia y pega el siguiente código en hello.go.

~~~go
package main

import (
"fmt"
"log"
"example.com/greetings"
)

func main() {
    // Establece las propiedades del Logger predefinido, incluyendo
    // el prefijo de las entradas de registro y una bandera para deshabilitar la impresión
    // de la hora, el archivo fuente y el número de línea.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Solicita un mensaje de saludo.
    mensaje, err := greetings.Hello("")
    // Si se devolvió un error, imprímelo en la consola y
    // sal del programa.
    if err != nil {
        log.Fatal(err)
    }

    // Si no se devolvió un error, imprime el mensaje devuelto
    // en la consola.
    fmt.Println(mensaje)
}
~~~

En este código:

Configuraste el paquete log para imprimir el nombre del comando ("greetings: ") al inicio de sus mensajes de registro, sin una marca de tiempo ni información del archivo fuente.
Asignaste ambas devoluciones de la función Hello, incluido el error, a variables.
Cambió el argumento de Hello de nombre de Gladys a una cadena vacía, para probar el código de manejo de errores.
Buscaste un valor de error no nulo. No tiene sentido continuar en este caso.
Utilizaste las funciones del paquete log de la biblioteca estándar para mostrar información de error. Si se produce un error, utilizas la función Fatal del paquete log para imprimir el error y detener el programa.
En la línea de comandos en el directorio hello, ejecuta hello.go para confirmar que el código funciona.
Ahora que estás pasando un nombre vacío, obtendrás un error.

~~~bash
$ go run .
greetings: nombre vacío
exit status 1
~~~


Ese es un manejo común de errores en Go: Devolver un error como valor para que el llamador pueda verificarlo.

A continuación, utilizarás una slice de Go para devolver un saludo seleccionado al azar.

---
## Devolver un saludo aleatorio
En esta sección, cambiarás tu código para que, en lugar de devolver un saludo único cada vez, devuelva uno de varios mensajes de saludo predefinidos.

Para hacer esto, utilizarás un slice en Go. Un slice es similar a un array, excepto que su tamaño cambia dinámicamente a medida que agregas y eliminas elementos. El slice es uno de los tipos más útiles de Go.

Agregarás un pequeño slice para contener tres mensajes de saludo y luego harás que tu código devuelva uno de los mensajes de forma aleatoria. Para obtener más información sobre los slices, consulta los slices en el blog de Go.

En greetings/greetings.go, cambia tu código para que se vea así:

~~~go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello devuelve un saludo para la persona especificada.
func Hello(name string) (string, error) {
    // Si no se proporciona un nombre, devuelve un error con un mensaje.
    if name == "" {
        return name, errors.New("nombre vacío")
    }
    // Crea un mensaje utilizando un formato aleatorio.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de saludo. El mensaje devuelto
// se selecciona de forma aleatoria.
func randomFormat() string {
    // Un slice de formatos de mensajes.
    formats := []string{
        "¡Hola, %v! ¡Bienvenido!",
        "¡Qué bueno verte, %v!",
        "¡Saludos, %v! ¡Encantado de conocerte!",
    }

    // Devuelve un formato de mensaje seleccionado aleatoriamente especificando
    // un índice aleatorio para el slice de formatos.
    return formats[rand.Intn(len(formats))]
}
~~~
En este código:

- Agregas una función randomFormat que devuelve un formato seleccionado aleatoriamente para un mensaje de saludo. Ten en cuenta que randomFormat comienza con una letra minúscula, lo que significa que solo es accesible para el código en su propio paquete (es decir, no se exporta).
- En randomFormat, declaras un slice llamado formats con tres formatos de mensaje. Al declarar un slice, omites su tamaño entre corchetes, de esta manera: []string. Esto le indica a Go que el tamaño del array subyacente del slice puede cambiarse dinámicamente.
- Utilizas el paquete math/rand para generar un número aleatorio para seleccionar un elemento del slice.
- Agregas una función init para inicializar el paquete rand con la hora actual. Go ejecuta las funciones init automáticamente al iniciar el programa, después de que las variables globales hayan sido inicializadas. Para obtener más información sobre las funciones init, consulta Effective Go.
- En Hello, llamas a la función randomFormat para obtener un formato para el mensaje que devolverás, luego utilizas el formato y el valor del nombre para crear el mensaje.
- Devuelves el mensaje (o un error) como lo hiciste anteriormente.
En hello/hello.go, cambia tu código para que se vea así:

~~~go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Establece las propiedades del Logger predefinido, incluido
    // el prefijo de entrada de registro y una marca para deshabilitar la impresión
    // de la hora, el archivo fuente y el número de línea.
    log.SetPrefix("saludos: ")
    log.SetFlags(0)

    // Solicita un mensaje de saludo.
    message, err := greetings.Hello("Gladys")
    // Si se devuelve un error, imprímelo en la consola y
    // finaliza el programa.
    if err != nil {
        log.Fatal(err)
    }

    // Si no se devuelve un error, imprime el mensaje devuelto
    // en la consola.
    fmt.Println(message)
}

~~~

En este código:

- Configuras el paquete log para imprimir el nombre del comando ("saludos: ") al inicio de sus mensajes de registro, sin una marca de tiempo ni información del archivo fuente.
- Asignas los dos valores de retorno de Hello, incluido el error, a variables.
- Cambias el argumento de Hello de "Gladys" a un string vacío para probar tu código de manejo de errores.
- Buscas un valor de error que no sea nulo. No tiene sentido continuar en este caso.
- Utilizas las funciones del paquete log de la biblioteca estándar para imprimir información de error. Si hay un error, utilizas la función Fatal del paquete log para imprimir el error y detener el programa.
En la línea de comandos, en el directorio hello, ejecuta hello.go para confirmar que el código funciona. Ejecútalo varias veces y observa que el saludo cambia.

~~~bash
$ go run .
¡Qué bueno verte, Gladys!

$ go run .
¡Hola, Gladys! ¡Bienvenido!

$ go run .
¡Saludos, Gladys! ¡Encantado de conocerte!

~~~

A continuación, utilizarás un slice para saludar a varias personas.

---
### Devolver saludos para varias personas
En los últimos cambios que harás en el código de tu módulo, agregarás soporte para obtener saludos para múltiples personas en una sola solicitud. En otras palabras, manejarás una entrada de múltiples valores y luego emparejarás esos valores de entrada con una salida de múltiples valores. Para hacer esto, deberás pasar un conjunto de nombres a una función que pueda devolver un saludo para cada uno de ellos.

Sin embargo, hay un problema. Cambiar el parámetro de la función Hello de un solo nombre a un conjunto de nombres cambiaría la firma de la función. Si ya hubieras publicado el módulo example.com/greetings y los usuarios ya hubieran escrito código que llama a Hello, ese cambio rompería sus programas.

En esta situación, una mejor opción es escribir una nueva función con un nombre diferente. La nueva función tomará múltiples parámetros. Esto preservará la antigua función para mantener la compatibilidad hacia atrás.

En greetings/greetings.go, cambia tu código para que se parezca al siguiente:

~~~go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}

~~~

En este código, tú:

- Agregas una función llamada Hellos cuyo parámetro es un slice de nombres en lugar de un solo nombre. Además, cambias uno de sus tipos de retorno de una cadena a un mapa para poder devolver nombres asociados con mensajes de saludo.

- Haces que la nueva función Hellos llame a la función Hello existente. Esto ayuda a reducir la duplicación y también deja ambas funciones en su lugar.

- Creas un mapa llamado messages para asociar cada uno de los nombres recibidos (como clave) con un mensaje generado (como valor). En Go, inicializas un mapa con la siguiente sintaxis: make(map[key-type]value-type). Haces que la función Hellos devuelva este mapa al llamador. Para obtener más información sobre los mapas, consulta "Go maps in action" en el blog de Go.

- Recorres los nombres que recibió tu función, verificando que cada uno tenga un valor no vacío, y luego asocias un mensaje con cada uno. En este bucle for, range devuelve dos valores: el índice del elemento actual en el bucle y una copia del valor del elemento. No necesitas el índice, por lo que utilizas el identificador en blanco de Go (un guion bajo) para ignorarlo. Para obtener más información, consulta "The blank identifier" en "Effective Go".

En tu código de llamada hello/hello.go, pasas un slice de nombres y luego imprimes el contenido del `mapa names/message`s que obtienes.

En hello.go, cambia tu código para que se parezca al siguiente:
~~~go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Establece las propiedades del Logger predefinido, incluyendo
    // el prefijo de las entradas de registro y una bandera para desactivar
    // la impresión del tiempo, el archivo fuente y el número de línea.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Un slice de nombres.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Solicita mensajes de saludo para los nombres.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // Si no se devolvió ningún error, imprime el mapa de
    // mensajes devuelto en la consola.
    fmt.Println(messages)
}

~~~

Con estos cambios, tú:

Creas la variable names como un slice que contiene tres nombres.

Pasas la variable names como argumento a la función Hellos.

En la línea de comandos, ve al directorio que contiene hello/hello.go y utiliza go run para confirmar que el código funciona.

La salida debería ser una representación en cadena del mapa que asocia nombres con mensajes, algo como esto:

~~~bash
$ go run .
~~~

Este tema introdujo los mapas para representar pares de nombre/valor. También introdujo la idea de preservar la compatibilidad hacia atrás mediante la implementación de una nueva función para una funcionalidad nueva o cambiada en un módulo. Para obtener más información sobre la compatibilidad hacia atrás, consulta "Keeping your modules compatible".

A continuación, utilizarás las características integradas de Go para crear una prueba unitaria para tu código.

---
## Agregar una prueba
Ahora que has llevado tu código a un estado estable (¡buen trabajo, por cierto!), es hora de agregar una prueba. Probar tu código durante el desarrollo puede revelar errores que se cuelan mientras realizas cambios. En este tema, agregarás una prueba para la función Hello.

El soporte integrado de Go para pruebas unitarias facilita la prueba a medida que avanzas. Específicamente, utilizando convenciones de nomenclatura, el paquete de pruebas de Go y el comando go test, puedes escribir y ejecutar rápidamente pruebas.

En el directorio `greetings`, crea un archivo llamado `greetings_test.go`.
Al finalizar el nombre de un archivo con _test.go, le indica al comando `go test` que este archivo contiene funciones de prueba.

En `greetings_test.go`, pega el siguiente código y guarda el archivo:

~~~go
package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

~~~

En este código, tú:

- Implementas funciones de prueba en el mismo paquete que el código que estás probando.

- Creas dos funciones de prueba para probar la función Hello de greetings. Los nombres de las funciones de prueba tienen la forma TestNombre, donde Nombre describe algo sobre la prueba específica. Además, las funciones de prueba toman como parámetro un puntero al tipo testing.T del paquete de pruebas. Utilizas los métodos de este parámetro para informar y registrar desde tu prueba.

- Implementas dos pruebas:

    - TestHelloName llama a la función Hello, pasando un nombre con el que la función debería poder devolver un mensaje de respuesta válido. Si la llamada devuelve un error o un mensaje de respuesta inesperado (uno que no incluye el nombre que pasaste), utilizas el método Fatalf del parámetro t para imprimir un mensaje en la consola y finalizar la ejecución.
    - TestHelloEmpty llama a la función Hello con una cadena vacía. Esta prueba está diseñada para confirmar que el manejo de errores funciona correctamente. Si la llamada devuelve una cadena no vacía o no devuelve ningún error, utilizas el método Fatalf del parámetro t para imprimir un mensaje en la consola y finalizar la ejecución.

En la línea de comandos en el directorio greetings, ejecuta el comando go test para ejecutar la prueba.

El comando go test ejecuta las funciones de prueba (cuyos nombres comienzan con Test) en archivos de prueba (cuyos nombres terminan con _test.go). Puedes agregar la bandera -v para obtener una salida detallada que liste todas las pruebas y sus resultados.

Las pruebas deberían pasar.

~~~go
$ go test
PASS
ok      example.com/greetings   0.364s

$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example.com/greetings   0.372s
~~~

Interrumpe la función Hello de `greetings` para ver una prueba fallida.

La función de prueba TestHelloName verifica el valor de retorno para el nombre que especificaste como parámetro de la función Hello. Para ver un resultado de prueba fallido, cambia la función Hello en greetings/greetings.go para que ya no incluya el nombre.

En greetings/greetings.go, pega el siguiente código en lugar de la función Hello. Ten en cuenta que las líneas resaltadas cambian el valor que la función devuelve, como si se hubiera eliminado accidentalmente el argumento del nombre.

~~~go
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    // message := fmt.Sprintf(randomFormat(), name)
    message := fmt.Sprint(randomFormat())
    return message, nil
}
~~~

En la línea de comandos en el directorio greetings, ejecuta go test para ejecutar la prueba.

Esta vez, ejecuta go test sin la bandera -v. La salida incluirá resultados solo para las pruebas que fallaron, lo cual puede ser útil cuando tienes muchas pruebas. La prueba TestHelloName debería fallar, mientras que TestHelloEmpty aún pasa.

~~~go
$ go test
--- FAIL: TestHelloName (0.00s)
    greetings_test.go:15: Hello("Gladys") = "Hail, %v! Well met!", <nil>, want match for `\bGladys\b`, nil
FAIL
exit status 1
FAIL    example.com/greetings   0.182s

~~~

En el próximo tema, verás cómo compilar e instalar tu código para ejecutarlo localmente.

---
## Compilar e instalar la aplicación
Compilar e instalar la aplicación

En este último tema, aprenderás dos comandos nuevos de Go. Si bien el comando `go run` es útil para compilar y ejecutar un programa cuando estás realizando cambios frecuentes, no genera un ejecutable binario.

Este tema introduce dos comandos adicionales para compilar código:

El comando go build compila los paquetes, junto con sus dependencias, pero no instala los resultados.
El comando go install compila e instala los paquetes.

Desde la línea de comandos en el directorio hello, ejecuta el comando go build para compilar el código en un ejecutable.

~~~bash
$ go build
~~~

Desde la línea de comandos en el directorio hello, ejecuta el nuevo ejecutable hello para confirmar que el código funciona.

Ten en cuenta que tu resultado puede diferir dependiendo de si cambiaste tu código en greetings.go después de probarlo.

En Linux o Mac:
~~~bash
$ ./hello
~~~

En Windows:

~~~shell
$ hello.exe
~~~

Has compilado la aplicación en un ejecutable para poder ejecutarla. Sin embargo, para ejecutarla actualmente, tu indicador debe estar en el directorio del ejecutable o debes especificar la ruta del ejecutable.

A continuación, instalarás el ejecutable para poder ejecutarlo sin especificar su ruta.

Descubre la ruta de instalación de Go, donde el comando go instalará el paquete actual.

Puedes descubrir la ruta de instalación ejecutando el comando go list, como en el siguiente ejemplo:

~~~bash
go list -f '{{.Target}}'
~~~

Por ejemplo, la salida del comando podría decir /home/gopher/bin/hello, lo que significa que los binarios se instalan en /home/gopher/bin. Necesitarás este directorio de instalación en el siguiente paso.

Agrega la ruta de instalación de Go al path de la shell de tu sistema.

De esta manera, podrás ejecutar el ejecutable de tu programa sin especificar dónde se encuentra el ejecutable.

En Linux o Mac, ejecuta el siguiente comando:
~~~bash
$ export PATH=$PATH:/ruta/a/tu/directorio/de/instalación
~~~

En Windows, ejecuta el siguiente comando:

~~~bash
$ set PATH=%PATH%;C:\ruta\a\tu\directorio\de\instalación
~~~

Como alternativa, si ya tienes un directorio como $HOME/bin en tu path de la shell y deseas instalar tus programas Go allí, puedes cambiar el destino de instalación estableciendo la variable GOBIN usando el comando go env:

~~~bash
$ go env -w GOBIN=/ruta/a/tu/bin
~~~

o

~~~shell
$ go env -w GOBIN=C:\ruta\a\tu\bin
~~~

Una vez que hayas actualizado el path de la shell, ejecuta el comando `go install` para compilar e instalar el paquete.

~~~bash
$ go install
~~~


Ejecuta tu aplicación simplemente escribiendo su nombre. Para hacer esto interesante, abre una nueva línea de comandos y ejecuta el nombre del ejecutable hello en algún otro directorio.

~~~bash
$ hello
~~~

---
## Desplegar a GitHub
Para desplegar el módulo de Go "greetings" en GitHub, sigue los siguientes pasos:

Crea un repositorio en GitHub:

1. Ve a la página principal de GitHub (https://github.com) e inicia sesión en tu cuenta.
    - Haz clic en el botón "New" (Nuevo) para crear un nuevo repositorio.
    - Proporciona un nombre para tu repositorio, por ejemplo, "greetings".
    - Opcionalmente, proporciona una descripción y selecciona las opciones de visibilidad y licencia deseadas.
    - Haz clic en el botón "Create repository" (Crear repositorio) para crear el repositorio vacío en GitHub.
2. Configura el repositorio local:
    - Abre una terminal o línea de comandos.
    - Navega al directorio raíz de tu proyecto local de Go.
    - Inicializa el repositorio Git ejecutando el comando git init.
    - Conecta tu repositorio local con el repositorio remoto en GitHub utilizando el comando git remote add origin `[URL_DEL_REPOSITORIO]`. Reemplaza `[URL_DEL_REPOSITORIO]` con la URL de tu repositorio en GitHub. Por ejemplo, git remote add origin https://github.com/tu-usuario/greetings.git.

3. Crea una estructura de directorios:
    - En el directorio raíz de tu proyecto, crea un directorio llamado greetings.
    - Dentro del directorio greetings, crea un archivo llamado greetings.go.
    - Copia y pega el código del módulo greetings en el archivo greetings.go.

4. Prepara y sube los archivos al repositorio:
    - Crea un archivo .gitignore en el directorio raíz de tu proyecto y agrega cualquier archivo o directorio que deseas ignorar en el control de versiones. Puedes encontrar ejemplos de archivos .gitignore para proyectos de Go en https://github.com/github/gitignore.
    - Agrega todos los archivos al área de preparación ejecutando el comando git add ..
    - Realiza un commit de los cambios ejecutando el comando git commit -m "Primer commit".
    - Sube los archivos al repositorio remoto en GitHub ejecutando el comando git push -u origin master. Esto creará una rama master en tu repositorio remoto y subirá los archivos.

5. Verifica que el módulo esté funcionando correctamente:
    - Ejecuta los comandos go build y go test para asegurarte de que tu módulo compile y pase las pruebas correctamente.

6. Actualiza tu repositorio en GitHub:
    - Haz cambios en tu módulo de Go según sea necesario.
    - Repite los pasos 4 y 5 para agregar, confirmar y subir los cambios a tu repositorio en GitHub.

Ahora tu módulo de Go "greetings" está desplegado en GitHub y está listo para ser utilizado por otros desarrolladores.

### Archivo README.md

~~~md
# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/alexroel/greetings
```
## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
    "fmt"
    "github.com/alexroel/greetings"
)

func main() {
    message, err := greetings.Hello("Alex")

    if err != nil {
        fmt.Println("Ocurrió un error:", err)
        return
    }

    fmt.Println(message)
}

```
Este ejemplo importa el paquete github.com/alexroel/greetings y llama a la función Hello para obtener un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.
~~~

---
## Resumen
En la sección "Crear un módulo" del curso de Go, aprendimos cómo iniciar un módulo en Go para que otros desarrolladores pudieran utilizarlo en sus proyectos. También exploramos cómo llamar a nuestro código desde otro módulo y cómo manejar y devolver errores adecuadamente.

Además, descubrimos cómo agregar una divertida funcionalidad a nuestro módulo: devolver un saludo aleatorio. Esto le dio un toque especial a nuestras aplicaciones.

Por supuesto, no nos olvidamos de la importancia de las pruebas en el desarrollo de software. Aprendimos cómo agregar pruebas a nuestro módulo para asegurarnos de que funcionaba correctamente en diferentes escenarios.

Una vez que construimos nuestro módulo y probamos su funcionalidad, aprendimos cómo compilar e instalar nuestra aplicación. Esto nos permitió utilizar nuestro módulo en cualquier proyecto sin problemas.

Finalmente, descubrimos cómo desplegar nuestro módulo en GitHub, una plataforma popular para compartir y colaborar en proyectos de código abierto.

Exploramos el emocionante mundo de la creación de módulos en Go y compartimos nuestro código con otros desarrolladores.
