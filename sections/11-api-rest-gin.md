# API RESTful con Go y Gin
1. [Introducción](#introducción)
2. [Crear proyecto](#crear-proyecto)
3. [Instalar Gin](#instalar-gin)
4. [Crear los datos](#crear-los-datos)
5. [Controlador para devolver todos los elementos](#controlador-para-devolver-todos-los-elementos)
6. [Controlador para agregar un nuevo elemento](#controlador-para-agregar-un-nuevo-elemento)
7. [Controlador para devolver un elemento específico](#controlador-para-devolver-un-elemento-específico)
8. [Resumen](#resumen)
---
## Introducción 
En esta sección, aprenderemos a desarrollar una API RESTful utilizando el lenguaje de programación Go y el framework Gin. Crearemos un proyecto desde cero y exploraremos diferentes aspectos de la creación de una API, como la instalación de Gin, la creación de datos, y la implementación de controladores para manejar diversas operaciones.

Los temas que cubriremos son los siguientes:

Crear el proyecto: Comenzaremos creando un nuevo proyecto Go y configurando el entorno de desarrollo necesario. Instalaremos todas las dependencias requeridas, incluido el framework Gin, y configuraremos la estructura básica del proyecto.

Instalar Gin: Aprenderemos cómo instalar el framework Gin en nuestro proyecto. Gin es un framework ligero y de alto rendimiento que nos facilitará la construcción de nuestra API RESTful.

Crear los datos: Exploraremos cómo crear y almacenar los datos en memoria. Aunque en un entorno de producción utilizaríamos una base de datos, en este tutorial utilizaremos datos almacenados en memoria para simplificar las cosas.

Controlador para devolver todos los elementos: Implementaremos un controlador que maneje las solicitudes GET para obtener todos los elementos de nuestra API. Aprenderemos cómo estructurar y devolver los datos en formato JSON.

Controlador para agregar un nuevo elemento: Implementaremos un controlador que maneje las solicitudes POST para agregar un nuevo elemento a nuestra API. Aprenderemos cómo recibir y validar los datos enviados por el cliente, y cómo agregar el nuevo elemento a nuestra colección de datos.

Controlador para devolver un elemento específico: Implementaremos un controlador que maneje las solicitudes GET para obtener un elemento específico de nuestra API. Aprenderemos cómo utilizar parámetros en la URL para identificar y recuperar el elemento solicitado.

A medida que avancemos en cada uno de estos temas, iremos construyendo gradualmente nuestra API RESTful con Go y Gin. ¡Vamos a comenzar creando nuestro proyecto y explorando el framework Gin!

---
## Crear proyecto 
Construirás una API que proporciona acceso a una tienda que vende grabaciones antiguas en vinilo. Por lo tanto, deberás proporcionar puntos finales a través de los cuales un cliente pueda obtener y agregar álbumes para los usuarios.

### Diseñar los puntos finales de la API

Al desarrollar una API, normalmente comienzas por diseñar los puntos finales. Los usuarios de tu API tendrán más éxito si los puntos finales son fáciles de entender.

Aquí están los puntos finales que crearás en este tutorial:

- `/albums`
    - GET: Obtener una lista de todos los álbumes, devueltos como JSON.
    - POST: Agregar un nuevo álbum a partir de los datos de solicitud enviados como JSON.
- `/albums/:id`
    - GET: Obtener un álbum por su ID, devolviendo los datos del álbum como JSON.

### Crear una carpeta para tu código
Para empezar, crea un proyecto para el código que escribirás.

Abre una terminal y cambia al directorio de tu directorio principal.

En Linux o Mac:
~~~bash
$ cd
~~~

En Windows:
~~~shell
C:> cd %HOMEPATH%
~~~
Usando la terminal, crea un directorio para tu código llamado "vinyl-api".
~~~bash
$ mkdir vinyl-api
$ cd vinyl-api
~~~
Crea un módulo en el cual puedas gestionar las dependencias.

Ejecuta el comando go mod init, proporcionándole la ruta del módulo en el cual estará tu código.

~~~bash
$ go mod init example/vinyl-api
go: creando nuevo go.mod: módulo example/vinyl-api
~~~

Este comando crea un archivo go.mod en el cual se listarán las dependencias que agregues para su seguimiento. Para obtener más información sobre cómo nombrar un módulo con una ruta de módulo, consulta la documentación sobre gestión de dependencias.

---
## Instalar Gin
Para instalar Gin en tu proyecto Go, puedes utilizar el comando go get:
~~~bash
go get -u github.com/gin-gonic/gin
~~~

Una vez instalado, puedes crear un archivo llamado main.go y agregar el siguiente código para crear un simple "Hola Mundo" usando Gin:

~~~go
package main

import "github.com/gin-gonic/gin"

func main() {
    // Crea una nueva instancia de Gin
    router := gin.Default()

    // Define la ruta GET para el endpoint "/"
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "¡Hola Mundo!",
        })
    })

    // Inicia el servidor en el puerto 8080
    router.Run("localhost:8080")
}
~~~
Ahora, puedes acceder a `http://localhost:8080` en tu navegador o utilizar herramientas como cURL o Postman para hacer una solicitud GET a la ruta "/" y verás la respuesta "¡Hola Mundo!" en formato JSON.

Recuerda importar el paquete "github.com/gin-gonic/gin" al principio de tu archivo main.go para utilizar Gin en tu proyecto.

---
## Crear los datos
Para mantener las cosas simples en el tutorial, almacenarás los datos en memoria. Por lo general, una API más típica interactuaría con una base de datos.

Ten en cuenta que almacenar los datos en memoria significa que el conjunto de álbumes se perderá cada vez que detengas el servidor y luego se recreará cuando lo reinicies.

Debajo de la declaración de paquete, pega la siguiente declaración de una estructura llamada album. La utilizarás para almacenar los datos de los álbumes en memoria.

Las etiquetas de estructura como json:"artist" especifican cómo debería ser el nombre de un campo cuando se serializan los contenidos de la estructura en JSON. Sin ellas, el JSON utilizaría los nombres de campo en mayúscula de la estructura, un estilo menos común en JSON.

~~~go
// album representa datos sobre un álbum musical.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
~~~

Después de la declaración de la estructura, pega la siguiente lista de estructuras de álbum que contienen los datos que usarás para empezar.

~~~go
// slice de albums para almacenar datos iniciales de álbumes.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
~~~

A continuación, escribirás el código para implementar tu primer punto final.

---
## Controlador para devolver todos los elementos
Cuando el cliente hace una solicitud a `GET /albums`, quieres devolver todos los álbumes como JSON.

### Para hacer esto, escribirás lo siguiente:
- Lógica para preparar una respuesta
- Código para mapear la ruta de la solicitud a tu lógica
- Ten en cuenta que esto es lo contrario de cómo se ejecutarán en tiempo de ejecución, pero estás agregando dependencias primero y luego el código que depende de ellas.

Después del código de la estructura que agregaste en la sección anterior, pega el siguiente código para obtener la lista de álbumes.

Esta función getAlbums crea JSON a partir de la lista de estructuras de álbum y escribe el JSON en la respuesta.

~~~go
// getAlbums responde con la lista de todos los álbumes como JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
~~~

Escribes una función getAlbums que toma un parámetro gin.Context. Ten en cuenta que podrías haber dado a esta función cualquier nombre; ni Gin ni Go requieren un formato de nombre de función específico.

gin.Context es la parte más importante de Gin. Lleva los detalles de la solicitud, valida y serializa JSON, y más. (A pesar del nombre similar, esto es diferente del paquete de contexto incorporado de Go).

Llamas a Context.IndentedJSON para serializar la estructura en JSON y agregarla a la respuesta.

El primer argumento de la función es el código de estado HTTP que quieres enviar al cliente. Aquí estás pasando la constante StatusOK del paquete net/http para indicar 200 OK.

Ten en cuenta que puedes reemplazar Context.IndentedJSON con una llamada a Context.JSON para enviar JSON más compacto. En la práctica, la forma indentada es mucho más fácil de trabajar al depurar y la diferencia de tamaño suele ser pequeña.

Cerca de la parte superior de main.go, justo debajo de la declaración de la lista de álbumes, pega el siguiente código para asignar la función del controlador a una ruta del punto final.

Esto establece una asociación en la que getAlbums maneja las solicitudes a la ruta del punto final /albums.

~~~go
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
~~~

En este código:

Inicializas un enrutador de Gin utilizando Default.

Utilizas la función GET para asociar el método HTTP GET y la ruta /albums con una función controladora.

Ten en cuenta que estás pasando el nombre de la función getAlbums. Esto es diferente de pasar el resultado de la función, lo cual harías pasando getAlbums() (nota los paréntesis).

Utilizas la función Run para vincular el enrutador a un http.Server y comenzar el servidor.

Cerca de la parte superior de main.go, justo debajo de la declaración del paquete, importa los paquetes necesarios para admitir el código que acabas de escribir.

Las primeras líneas de código deberían lucir así:

~~~go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
~~~

Guarda main.go.

Ejecuta el código

Desde la línea de comandos en el directorio que contiene main.go, ejecuta el código. Usa el argumento de punto para indicar "ejecutar el código en el directorio actual".

~~~bash
$ go run .
~~~
Una vez que el código esté en ejecución, tendrás un servidor HTTP en funcionamiento al que puedes enviar solicitudes.

Desde una nueva ventana de línea de comandos, utiliza curl para hacer una solicitud a tu servicio web en ejecución.

~~~bash
$ curl http://localhost:8080/albums
~~~

El comando mostrará los datos con los que inicializaste el servicio.

~~~bash
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]
~~~
¡Has iniciado una API! En la siguiente sección, crearás otro punto final con código para manejar una solicitud POST para agregar un elemento.

---
## Controlador para agregar un nuevo elemento
Cuando el cliente realiza una solicitud POST a /albums, quieres agregar el álbum descrito en el cuerpo de la solicitud a los datos de álbumes existentes.

Para hacer esto, escribirás lo siguiente:

Agrega código para agregar los datos del álbum a la lista de álbumes.

En algún lugar después de las declaraciones de importación, pega el siguiente código. (El final del archivo es un buen lugar para este código, pero Go no impone el orden en el que declaras las funciones).

~~~go
// postAlbums agrega un álbum desde JSON recibido en el cuerpo de la solicitud.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Llama a BindJSON para vincular el JSON recibido a
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Agrega el nuevo álbum a la lista.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
~~~

En este código, haces lo siguiente:

Utilizas Context.BindJSON para vincular el cuerpo de la solicitud a newAlbum.
Agregas la estructura de álbum inicializada a partir del JSON a la lista de álbumes.
Agregas un código de estado 201 a la respuesta, junto con el JSON que representa el álbum que agregaste.

Cambia tu función main para que incluya la función router.POST, como se muestra a continuación.
~~~go
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
~~~
En este código, haces lo siguiente:

Asocias el método POST en la ruta /albums con la función postAlbums.

Con Gin, puedes asociar un controlador con una combinación de método y ruta HTTP. De esta manera, puedes enrutar por separado las solicitudes enviadas a una sola ruta en función del método que esté utilizando el cliente.

Ejecuta el código

Si el servidor sigue en ejecución desde la sección anterior, detenlo.

Desde la línea de comandos en el directorio que contiene main.go, ejecuta el código.
~~~bash
$ go run .
~~~

Desde una ventana de línea de comandos diferente, utiliza curl para hacer una solicitud a tu servicio web en ejecución.

~~~bash
$ curl http://localhost:8080/albums
--include
--header "Content-Type: application/json"
--request "POST"
--data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
~~~

El comando debería mostrar las cabeceras y el JSON del álbum agregado.

~~~bash
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Mié, 02 Jun 2021 00:34:12 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
~~~

Al igual que en la sección anterior, utiliza curl para obtener la lista completa de álbumes, lo cual puedes hacer para confirmar que se agregó el nuevo álbum.
~~~bash
$ curl http://localhost:8080/albums
--header "Content-Type: application/json"
--request "GET"
~~~

El comando debería mostrar la lista de álbumes.

~~~bash
[
  {
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  },
  {
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
  },
  {
    "id": "3",
    "title": "Sarah Vaughan and Clifford Brown",
    "artist": "Sarah Vaughan",
    "price": 39.99
  },
  {
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
  }
]
~~~

En la siguiente sección, agregarás código para manejar una solicitud GET para un elemento específico.

---
## Controlador para devolver un elemento específico
Cuando el cliente realiza una solicitud GET a `/albums/[id]`, quieres devolver el álbum cuyo ID coincide con el parámetro de ruta id.

Después de la función postAlbums que agregaste en la sección anterior, pega el siguiente código para recuperar un álbum específico.

Esta función getAlbumByID extraerá el ID de la ruta de la solicitud y luego buscará un álbum que coincida.

~~~go
// getAlbumByID encuentra el álbum cuyo valor de ID coincide con el parámetro id
// enviado por el cliente, y luego devuelve ese álbum como respuesta.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Recorre la lista de álbumes, buscando
	// un álbum cuyo valor de ID coincida con el parámetro.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}
~~~

En este código, haces lo siguiente:

Utilizas Context.Param para recuperar el parámetro de ruta id de la URL. Cuando mapees este controlador a una ruta, incluirás un marcador de posición para el parámetro en la ruta.

Recorres las estructuras de álbum en la lista, buscando una cuyo campo de ID coincida con el valor del parámetro id. Si se encuentra, serializas esa estructura de álbum a JSON y la devuelves como respuesta con un código HTTP 200 OK.

Como se mencionó anteriormente, un servicio del mundo real probablemente usaría una consulta a una base de datos para realizar esta búsqueda.

Devuelves un error HTTP 404 con http.StatusNotFound si no se encuentra el álbum.

Finalmente, cambia tu función main para que incluya una nueva llamada a router.GET, donde la ruta ahora es /albums/:id, como se muestra en el siguiente ejemplo.

~~~go
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
~~~

En este código, haces lo siguiente:

Asocias la ruta /albums/:id con la función getAlbumByID. En Gin, los dos puntos que preceden a un elemento en la ruta indican que el elemento es un parámetro de ruta.

Ejecuta el código

Si el servidor sigue en ejecución desde la sección anterior, detenlo.

Desde la línea de comandos en el directorio que contiene main.go, ejecuta el código para iniciar el servidor.

~~~bash
$ go run .
~~~

Desde una ventana de línea de comandos diferente, utiliza curl para hacer una solicitud a tu servicio web en ejecución.

~~~bash
$ curl http://localhost:8080/albums/2
~~~

El comando debería mostrar el JSON del álbum cuyo ID utilizaste. Si el álbum no se encuentra, recibirás JSON con un mensaje de error.

~~~bash
{
"id": "2",
"title": "Jeru",
"artist": "Gerry Mulligan",
"price": 17.99
}
~~~

---
## Resumen 
En esta sección, aprendimos a desarrollar una API RESTful utilizando el lenguaje de programación Go y el framework Gin. Comenzamos creando un nuevo proyecto Go y configuramos el entorno de desarrollo necesario. Instalamos el framework Gin y configuramos la estructura básica del proyecto.

A continuación, creamos datos de ejemplo y exploramos cómo almacenarlos en memoria. Implementamos controladores para manejar las solicitudes GET y POST. El controlador de GET nos permitió devolver todos los elementos de la API en formato JSON, mientras que el controlador de POST nos permitió agregar nuevos elementos a la colección de datos.

Además, implementamos un controlador para manejar solicitudes GET específicas, donde utilizamos parámetros en la URL para identificar y devolver un elemento específico de la API.

Durante este proceso, aprendimos cómo utilizar las funciones y características proporcionadas por el framework Gin, como la serialización de JSON y la gestión de rutas.

Con estos conocimientos, hemos sido capaces de construir una API RESTful básica y comprender los fundamentos de Go y Gin en el contexto de la creación de una API.