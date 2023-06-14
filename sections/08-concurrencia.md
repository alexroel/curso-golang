# Concurrencia 
1. [Introducción](#introducción)
2. [Concurrencia en Go](#concurrencia-en-go)
3. [Uso de concurrencia](#uso-de-concurrencia)
4. [Uso de canales](#uso-de-canales)
5. [Resumen](#resumen)

---
## Introducción 
¡Bienvenidos a la sección de Concurrencia en Go! En esta parte del curso, exploraremos cómo Go, un lenguaje de programación moderno y poderoso, maneja la concurrencia de manera efectiva. La concurrencia es esencial para aprovechar al máximo los recursos de hardware y mejorar el rendimiento de nuestras aplicaciones.

En esta sección, nos centraremos en dos aspectos clave de la concurrencia en Go: el uso de goroutines y el manejo de canales. Las goroutines son funciones independientes que se ejecutan de forma concurrente y ligera dentro de un programa Go. Estas goroutines nos permiten realizar múltiples tareas simultáneamente y aprovechar al máximo la potencia de los procesadores modernos.

Por otro lado, los canales son la herramienta fundamental para la comunicación y sincronización entre goroutines. Nos permiten enviar y recibir datos de manera segura y coordinada, evitando condiciones de carrera y garantizando una ejecución coherente de nuestras goroutines.

En esta sección, aprenderemos cómo crear y utilizar goroutines, cómo comunicarnos a través de canales y cómo aprovechar estas poderosas herramientas para resolver problemas complejos de concurrencia. Además, exploraremos patrones comunes de concurrencia en Go y técnicas avanzadas para maximizar el rendimiento de nuestras aplicaciones.

Así que, prepárate para sumergirte en el mundo de la concurrencia en Go y descubrir cómo este lenguaje nos brinda las herramientas necesarias para escribir programas rápidos, eficientes y concurrentes. ¡Vamos a empezar!

---
## Concurrencia en Go 

En Go, la concurrencia es una característica fundamental del lenguaje que permite la ejecución concurrente de múltiples goroutines (tareas ligeras similares a hilos). Las goroutines son independientes entre sí y se ejecutan de manera concurrente, lo que significa que pueden ejecutarse de forma simultánea o en un orden no determinista.

La concurrencia en Go se basa en el modelo de concurrencia de comunicación secuencial por procesos (CSP, por sus siglas en inglés), que se centra en la comunicación entre procesos ligeros en lugar de compartir datos. Go proporciona primitivas de concurrencia, como canales, para facilitar la comunicación y sincronización entre goroutines.

La concurrencia en Go permite aprovechar eficientemente los recursos de hardware multi núcleo y construir programas concurrentes y paralelos de manera concisa. Al utilizar goroutines en lugar de hilos pesados, Go puede manejar un gran número de tareas concurrentes sin incurrir en el costo adicional de crear y administrar múltiples hilos.

La concurrencia en Go se considera una de las principales fortalezas del lenguaje y ha sido diseñada para ser fácil de usar y comprender. La sintaxis y las primitivas proporcionadas por Go hacen que sea relativamente sencillo escribir programas concurrentes de manera segura y eficiente.

---
## Uso de concurrencia 
La aplicación sin concurrencia en Go es un programa que verifica el estado de diferentes APIs mediante solicitudes HTTP. El objetivo es determinar si cada API está en funcionamiento o si ha experimentado algún error.

El programa comienza registrando el tiempo de inicio de la ejecución. Luego, se define una lista de URLs de las APIs a verificar. Cada una de estas URLs se almacena en un slice llamado apis.

A continuación, se inicia un bucle for para recorrer cada una de las URLs en apis. En cada iteración del bucle, se llama a la función checkAPI pasando como argumento la URL actual. Esta función se encarga de realizar una solicitud HTTP GET a la URL y verificar si se produce algún error.

Si se produce un error durante la solicitud, se imprime un mensaje de error indicando que la API está caída. En caso contrario, se imprime un mensaje de éxito indicando que la API está en funcionamiento.

Una vez que se han verificado todas las APIs, se calcula el tiempo transcurrido desde el inicio del programa y se imprime un mensaje indicando el tiempo total de ejecución.

Es importante destacar que esta aplicación no utiliza concurrencia, lo que significa que las solicitudes HTTP se realizan de forma secuencial, una después de la otra. Esto implica que el programa esperará a que cada solicitud se complete antes de pasar a la siguiente.

~~~go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	// Recorreer los apis
	for _, api := range apis {
		checkAPI(api)
	}

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())
}

// Función que verifica los APIS
func checkAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: ¡%s está caído!\n", api)
		return
	}

	fmt.Printf("SUCCESS: ¡%s está en funcionamiento!\n", api)
}
~~~

- Se importan los paquetes necesarios: "fmt" para imprimir mensajes en la consola, "net/http" para realizar solicitudes HTTP y "time" para medir el tiempo de ejecución.

- En la función main(), se registra el tiempo de inicio de la ejecución utilizando time.Now() y se guarda en la variable start.

- Se define un slice llamado apis que contiene una lista de URLs de las APIs a verificar.

- A continuación, se inicia un bucle for para recorrer cada una de las URLs en apis.

- En cada iteración del bucle, se llama a la función checkAPI(api) pasando como argumento la URL actual. Esta función se encarga de verificar el estado de la API correspondiente.

- La función checkAPI(api) realiza una solicitud HTTP GET a la URL especificada por api utilizando http.Get(api). La respuesta obtenida y un posible error se asignan a las variables _ y err, respectivamente. El uso del guion bajo _ indica que no estamos interesados en el valor de la respuesta en este caso.

- Si se produce un error durante la solicitud (por ejemplo, si la API está caída o no responde), se imprime un mensaje de error indicando que la API está caída.

- En caso contrario, si no se produce ningún error durante la solicitud, se imprime un mensaje de éxito indicando que la API está en funcionamiento.

- Una vez que se han verificado todas las APIs, se calcula el tiempo transcurrido desde el inicio del programa utilizando time.Since(start). El resultado se guarda en la variable elapsed.

- Finalmente, se imprime un mensaje indicando el tiempo total de ejecución en segundos utilizando fmt.Printf().

### Agregando concurrencia
Para crear una goroutine, es necesario usar la palabra clave go antes de llamar a una función. 

~~~go
	// Recorreer los apis
	for _, api := range apis {
		go checkAPI(api)
	}
~~~
Vuelva a ejecutar el programa y observe lo que sucede.
Parece que el programa ya no comprueba las API, ¿verdad? Es posible que vea algo parecido a la salida siguiente:

~~~bash
¡Listo! ¡Tomó 2.7371e-05 segundos!
~~~
¡Muy rápido! ¿Qué ha ocurrido? Verá el mensaje final que indica que el programa ha finalizado porque Go ha creado una goroutine para cada sitio dentro del bucle e inmediatamente a pasado a la siguiente línea.

Aunque no parece que la función checkAPI se esté ejecutando, realmente sí lo está haciendo. Simplemente no tuvo tiempo de finalizarse. Observe lo que ocurre si incluye un temporizador de suspensión justo después del bucle:

~~~go
	// Recorreer los apis
	for _, api := range apis {
		go checkAPI(api)
	}

	time.Sleep(5 * time.Second)
~~~

Ahora, cuando vuelva a ejecutar el programa, podría ver una salida similar a la siguiente:

~~~bash
ERROR: ¡https://api.somewhereintheinternet.com/ está caído!
SUCCESS: ¡https://api.github.com está en funcionamiento!
SUCCESS: ¡https://dev.azure.com está en funcionamiento!
SUCCESS: ¡https://management.azure.com está en funcionamiento!
SUCCESS: ¡https://outlook.office.com/ está en funcionamiento!
SUCCESS: ¡https://graph.microsoft.com está en funcionamiento!
¡Listo! ¡Tomó 5.002491318 segundos!
~~~

Parece que funciona, ¿verdad? En realidad, no exactamente. ¿Qué ocurre si desea agregar un nuevo sitio a la lista? Quizás tres segundos no son suficientes. ¿Cómo podría saberlo? No puede. Debe haber una manera mejor, y eso es lo que analizaremos en la sección siguiente cuando hablemos de los canales.

---
## Uso de canales 
En Go, los canales son una característica fundamental para la comunicación y sincronización entre goroutines (subprocesos ligeros) dentro de un programa concurrente. Un canal es una estructura que permite enviar y recibir valores entre goroutines, actuando como un conducto a través del cual fluye la información.

Para crear un canal en Go, se utiliza la función make() con la siguiente sintaxis:
~~~go
canal := make(chan tipoDato)
~~~

Donde tipoDato especifica el tipo de datos que se enviarán a través del canal. Puede ser cualquier tipo de datos válido en Go, como int, string, struct, etc.

Una vez creado el canal, se pueden enviar y recibir datos utilizando la notación de flecha <-. Por ejemplo:

~~~go
// Crear un canal de tipo entero
canal := make(chan int)

// Enviar un valor a través del canal
canal <- 10

// Recibir un valor del canal
valor := <-canal
~~~

La operación <- se utiliza para enviar un valor al canal (colocándolo a la izquierda de la flecha) o recibir un valor del canal (colocándolo a la derecha de la flecha).

### Canales y concurrencia
En el programa use canales para quitar la funcionalidad de suspensión. En primer lugar, vamos a crear un canal de cadena en la función main, como se indica a continuación:

~~~go
// Crear un canal de tipo string
ch := make(chan string)
~~~

Y quitaremos la línea de suspensión `time.Sleep(5 * time.Second)`.

Ahora, podemos usar canales para comunicarse entre goroutines. En lugar de imprimir el resultado en la función checkAPI, se refactorizará el código y ese mensaje se enviará por el canal. Para usar el canal desde esa función, debe agregar el canal como parámetro. La función checkAPI debe tener el siguiente aspecto:

~~~go
// Función que verifica los APIS
func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s está caído!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: ¡%s está en funcionamiento!\n", api)
}
~~~

Tenga en cuenta que es necesario usar la función fmt.Sprintf porque no quiere imprimir ningún texto, simplemente enviar texto con formato por el canal. Además, observe que usamos el operador `<-` después de la variable de canal para enviar datos.

Ahora debe cambiar la función main para enviar la variable de canal y recibir los datos para imprimirla, como se muestra a continuación:

~~~go
	// Recorreer los apis
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	// Leer datos de canal
	fmt.Println(<-ch)
~~~

Observe cómo usamos el operador <- antes de que el canal indique que queremos leer datos del canal.
Cuando vuelva a ejecutar el programa, verá una salida similar a la siguiente:

~~~bash
ERROR: ¡https://api.somewhereintheinternet.com/ está caído!

¡Listo! ¡Tomó 0.009662042 segundos!
~~~

Al menos funciona sin una llamada a una función de suspensión, ¿no? Pero todavía no hace lo que queremos. Vemos la salida solo de una de las goroutines, pero creamos cinco. En la siguiente clase descubriremos por qué este programa funciona de esta manera.

### Canales no almacenados en búfer
Cuando se crea un canal mediante la función make(), se crea un canal no almacenado en búfer, que es el comportamiento predeterminado. Los canales no almacenados en búfer bloquean la operación de envío hasta que algún componente esté listo para recibir los datos. Como se ha afirmado antes, el envío y la recepción son operaciones de bloqueo. Esta operación de bloqueo también es la razón por la que el programa de la sección anterior se ha detenido en cuanto ha recibido el primer mensaje.

Podemos empezar diciendo que fmt.Print(<-ch) bloquea el programa porque está leyendo de un canal y espera a que lleguen algunos datos. En cuanto hay algunos, continúa con la línea siguiente y el programa finaliza.

¿Qué ha ocurrido con el resto de las goroutines? Todavía se están ejecutando, pero ya no hay ninguna escuchando. Y dado que el programa terminó pronto, algunas goroutines no pudieron enviar datos. Para demostrar esto, vamos a agregar otra línea fmt.Print(<-ch), como se indica a continuación:

~~~go
	ch := make(chan string)

	// Recorreer los apis
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	fmt.Print(<-ch)
	fmt.Print(<-ch)
~~~
Cuando vuelva a ejecutar el programa, verá una salida similar a la siguiente:
~~~bash
ERROR: ¡https://api.somewhereintheinternet.com/ está caído!
SUCCESS: ¡https://api.github.com está en funcionamiento!
¡Listo! ¡Tomó 0.48367305 segundos!
~~~

Observe que ahora verá la salida de dos API. Si continúa agregando más líneas fmt.Print(<-ch), acabará leyendo todos los datos que se envían al canal. Pero ¿qué ocurre si intenta leer más datos y ya no hay ninguna goroutine que envíe datos? Por ejemplo:

~~~go
ch := make(chan string)

for _, api := range apis {
    go checkAPI(api, ch)
}

fmt.Print(<-ch)
fmt.Print(<-ch)
fmt.Print(<-ch)
fmt.Print(<-ch)
fmt.Print(<-ch)
fmt.Print(<-ch)

fmt.Print(<-ch)
~~~
Cuando vuelva a ejecutar el programa, verá una salida similar a la siguiente:

~~~bash
ERROR: ¡https://api.somewhereintheinternet.com/ está caído!
SUCCESS: ¡https://api.github.com está en funcionamiento!
SUCCESS: ¡https://management.azure.com está en funcionamiento!
SUCCESS: ¡https://dev.azure.com está en funcionamiento!
SUCCESS: ¡https://graph.microsoft.com está en funcionamiento!
SUCCESS: ¡https://outlook.office.com/ está en funcionamiento!

~~~

Funciona, pero el programa no finaliza. La última línea de impresión lo está bloqueando porque está esperando recibir datos. Tendrá que cerrar el programa con un comando como Ctrl+C.

El ejemplo anterior simplemente demuestra que la lectura y recepción de datos son operaciones de bloqueo. Para corregir este problema, podría cambiar el código a un bucle for y recibir solo los datos que sabe con certeza que va a enviar, como en este ejemplo:

~~~go
for i := 0; i < len(apis); i++ {
    fmt.Print(<-ch)
}
~~~

El programa está haciendo lo que se supone que debe hacer. Ya no usa una función de suspensión; usa canales. Observe también que ahora se tardan aproximadamente 1.357984 segundos en finalizar en lugar de los casi 5 segundos cuando no se usaba la simultaneidad.

---
## Resumen
En esta sección del curso de Go, exploramos cómo Go maneja la concurrencia de manera efectiva. Nos centramos en dos aspectos clave: el uso de goroutines y el manejo de canales.

Las goroutines son funciones independientes que se ejecutan de forma concurrente y ligera dentro de un programa Go. Aprendimos a crear y utilizar goroutines, aprovechando su capacidad para realizar múltiples tareas simultáneamente y maximizar el rendimiento del hardware.

Los canales, por otro lado, son herramientas fundamentales para la comunicación y sincronización entre goroutines. Aprendimos a enviar y recibir datos de manera segura y coordinada utilizando canales, evitando condiciones de carrera y garantizando una ejecución coherente de las goroutines.

Además, exploramos patrones comunes de concurrencia en Go y técnicas avanzadas para optimizar el rendimiento de nuestras aplicaciones.

En resumen, esta sección nos brindó una comprensión profunda de cómo aprovechar la concurrencia en Go, utilizando goroutines y canales para escribir programas rápidos, eficientes y concurrentes.
