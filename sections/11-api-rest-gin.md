# Desarrollando una API RESTful con Go y Gin

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

### Instalar Gin
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
    r := gin.Default()

    // Define la ruta GET para el endpoint "/"
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "¡Hola Mundo!",
        })
    })

    // Inicia el servidor en el puerto 8080
    r.Run(":8080")
}
~~~
Ahora, puedes acceder a `http://localhost:8080` en tu navegador o utilizar herramientas como cURL o Postman para hacer una solicitud GET a la ruta "/" y verás la respuesta "¡Hola Mundo!" en formato JSON.

Recuerda importar el paquete "github.com/gin-gonic/gin" al principio de tu archivo main.go para utilizar Gin en tu proyecto.

---
## Conexion a MongoDB Atlas