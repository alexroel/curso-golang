# Persistencia de datos con MySQL

1. [Introducción](#introducción)
2. [Instalación de MySQL](#instalación-de-mysql)
3. [Creación de base de datos](#creación-de-base-de-datos)
4. [Conexión a base de datos](#conexión-a-base-de-datos)
6. [Uso de godotenv](#uso-de-godotenv)
7. [Listar contactos](#listar-contactos)
8. [Obtener un contacto](#obtener-un-contacto)
7. [Registrar contacto](#registrar-contacto)
8. [Editar contacto](#editar-contacto)
9. [Eliminar contacto](#eliminar-contacto)
10. [Terminar el proyecto](#terminar-el-proyecto)
11. [Resumen](#resumen)

---
## Introducción
En esta sección del curso de Golang, exploraremos cómo trabajar con una base de datos MySQL para la persistencia de datos en nuestras aplicaciones. A lo largo de esta sección, aprenderemos los fundamentos necesarios para interactuar con una base de datos MySQL desde una aplicación Go.

#### Temas cubiertos en esta sección:

1. **Instalación de MySQL:** Aprenderemos cómo instalar y configurar MySQL en nuestro sistema para comenzar a trabajar con bases de datos.

2. **Creación de base de datos:** Descubriremos cómo crear una base de datos en MySQL utilizando comandos SQL básicos.

3. **Conexión a base de datos:** Veremos cómo establecer una conexión entre nuestra aplicación Go y la base de datos MySQL para poder realizar operaciones de lectura y escritura.

4. **Uso de godotenv:** Aprenderemos a utilizar la biblioteca `godotenv` para cargar variables de entorno desde un archivo `.env`, lo que nos permitirá gestionar de manera segura y conveniente la configuración de la conexión a la base de datos.

5. **Listar contactos:** Implementaremos la funcionalidad para listar todos los contactos almacenados en la base de datos.

6. **Obtener un contacto:** Aprenderemos a recuperar un contacto específico de la base de datos utilizando su identificador único.

7. **Registrar contacto:** Implementaremos la funcionalidad para registrar un nuevo contacto en la base de datos.

8. **Editar contacto:** Exploraremos cómo actualizar la información de un contacto existente en la base de datos.

9. **Eliminar contacto:** Aprenderemos a eliminar un contacto de la base de datos cuando sea necesario.

10. **Terminar el proyecto: con menú de navegación usando bufio:** Concluiremos la sección construyendo un menú de navegación interactivo utilizando el paquete `bufio`, que permitirá al usuario realizar todas las operaciones CRUD mencionadas anteriormente de manera intuitiva.

¡Al final de esta sección, estarás bien equipado para trabajar con bases de datos MySQL en tus proyectos Go y gestionar eficazmente la persistencia de datos!

---
## Instalación de MySQL

MySQL es un sistema de gestión de bases de datos relacional de código abierto ampliamente utilizado en el desarrollo de aplicaciones web y empresariales. En esta guía, aprenderás cómo instalar MySQL en Ubuntu dentro de Windows Subsystem for Linux (WSL) y cómo crear un nuevo usuario con privilegios en MySQL.

#### Instalación de MySQL en WSL Ubuntu

Para comenzar, primero asegúrate de tener acceso a tu terminal de WSL Ubuntu. Puedes abrirlo desde tu menú de inicio de Windows o desde tu aplicación WSL preferida. A continuación, sigue estos pasos:

1. **Verifica la versión de MySQL:**
   ```bash
   mysql --version
   ```
   Este comando te mostrará la versión de MySQL instalada en tu sistema, si ya está instalada.

2. **Instala MySQL:**
   ```bash
   sudo apt install mysql-server
   ```
   Esto instalará el servidor de MySQL en tu sistema. Durante la instalación, se te pedirá que configures una contraseña para el usuario "root" de MySQL.

3. **Consulta el manual de MySQL (opcional):**
   ```bash
   man mysql
   ```
   Esta es una buena práctica para familiarizarte con los comandos y opciones disponibles en MySQL.

4. **Inicia sesión como usuario root en MySQL:**
   ```bash
   sudo su
   mysql -u root -p
   ```
   Ingresa la contraseña de root que configuraste durante la instalación.

#### Creación de un Nuevo Usuario en MySQL

Una vez dentro de MySQL, puedes crear un nuevo usuario y otorgarle privilegios según sea necesario:

1. **Listar bases de datos disponibles:**
   ```sql
   show databases;
   ```
   Este comando te mostrará una lista de las bases de datos disponibles en tu servidor MySQL.

2. **Crear un nuevo usuario:**
   ```sql
   CREATE USER 'alexroel'@'localhost' IDENTIFIED BY '123456';
   ```
   Esto creará un nuevo usuario llamado "alexroel" con la contraseña "123456".

3. **Otorgar privilegios al nuevo usuario:**
   ```sql
   GRANT ALL PRIVILEGES ON *.* TO 'alexroel'@'localhost' WITH GRANT OPTION;
   ```
   Esto otorgará todos los privilegios al usuario "alexroel" en todas las bases de datos.

4. **Salir de MySQL:**
   ```sql
   \q;
   ```

#### Acceder al Servidor MySQL como el Nuevo Usuario

Una vez que hayas creado el nuevo usuario, puedes iniciar sesión en MySQL con ese usuario:

```bash
mysql -u alexroel -p
```

Ingresa la contraseña cuando se te solicite y estarás conectado como el usuario "alexroel" en el servidor MySQL.

#### Buenas Prácticas
- **Actualización regular**: Mantén tu sistema y software actualizado regularmente para evitar vulnerabilidades de seguridad.
- **Uso de contraseñas seguras**: Utiliza contraseñas fuertes y seguras para los usuarios de MySQL para proteger tus datos.
- **Principio de privilegios mínimos**: Asigna solo los privilegios necesarios a los usuarios de MySQL para limitar el acceso no autorizado.
- **Respaldo regular de datos**: Realiza copias de seguridad periódicas de tus bases de datos para proteger contra la pérdida de datos.

Siguiendo estos pasos y buenas prácticas, podrás instalar MySQL en tu entorno de WSL Ubuntu y administrar usuarios y bases de datos de manera segura y eficiente. ¡Explora más allá y sigue aprendiendo sobre bases de datos y MySQL!

### Cambiar contraseña de root para que sea mas segura

Si el usuario root de MySQL sigue aceptando cualquier contraseña, podría haber varias razones posibles para este comportamiento inesperado. Aquí hay algunos pasos adicionales que puedes seguir para abordar este problema:

### Paso 1: Verificar la configuración de autenticación

1. Accede a MySQL como usuario root:
   ```bash
   mysql -u root -p
   ```

2. Ejecuta el siguiente comando SQL para verificar cómo se está llevando a cabo la autenticación:
   ```sql
   SELECT host, user, authentication_string FROM mysql.user WHERE user='root';
   ```

3. Observa la columna `authentication_string`. Si ves que la columna está vacía o contiene algo diferente de una cadena de caracteres (como un asterisco '*'), eso puede indicar que el usuario root no tiene una contraseña establecida o que se está utilizando un método de autenticación diferente.

### Paso 2: Establecer la contraseña manualmente

Si la columna `authentication_string` está vacía, puedes intentar establecer manualmente una contraseña para el usuario root:

```sql
ALTER USER 'root'@'localhost' IDENTIFIED BY 'tu_nueva_contraseña';
```

Reemplaza `'tu_nueva_contraseña'` con la contraseña deseada.

### Paso 3: Verificar la autenticación basada en plugins

Asegúrate de que el usuario root no esté configurado para la autenticación basada en plugins. Puedes verificar esto ejecutando el siguiente comando SQL:

```sql
SELECT host, user, plugin FROM mysql.user WHERE user='root';
```

Si ves que el usuario root tiene un plugin asociado que no sea `mysql_native_password`, eso podría ser la causa del problema. Puedes cambiar el plugin a `mysql_native_password` con el siguiente comando:

```sql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'tu_nueva_contraseña';
```

### Paso 4: Reiniciar el servidor MySQL

Después de realizar cambios en la configuración de MySQL, es una buena práctica reiniciar el servidor para asegurarte de que los cambios surtan efecto:

```bash
sudo service mysql restart
```

---
## Creación de base de datos
Para crear una base de datos y una tabla en MySQL con el nombre `db_contacts` y una tabla llamada `contact`, así como insertar registros en esa tabla, puedes seguir estos pasos:

1. **Creación de la base de datos y la tabla:**
   Puedes usar el cliente de MySQL, como MySQL Workbench o HeidiSQL, o simplemente ejecutar comandos SQL desde la consola de MySQL. Aquí tienes un ejemplo de cómo hacerlo:

```sql
CREATE DATABASE IF NOT EXISTS db_contacts;

USE db_contacts;

CREATE TABLE IF NOT EXISTS contact (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(20)
);

SHOW TABLES;
```

Este código crea una base de datos llamada `db_contacts` y una tabla llamada `contact` con cuatro columnas: `id`, `name`, `email` y `phone`.

2. **Inserción de registros:**
   Después de crear la tabla, puedes insertar algunos registros de ejemplo. Aquí tienes un ejemplo de cómo hacerlo:

```sql
INSERT INTO contact (name, email, phone) VALUES
('Juan Pérez', 'juan@example.com', '123-456-7890'),
('María Gómez', 'maria@example.com', '987-654-3210'),
('Carlos López', NULL, '555-123-4567');

SELECT * FROM contact;
```

Esto inserta tres registros en la tabla `contact`. Ten en cuenta que he dejado el campo `email` de `Carlos López` como `NULL` para mostrar que la inserción maneja valores nulos.

Una vez que hayas ejecutado estos comandos SQL, tendrás una base de datos llamada `db_contacts` con una tabla `contact` y algunos registros insertados en ella. Ahora puedes interactuar con esta base de datos utilizando Go y MySQL para realizar operaciones como consultas, actualizaciones o eliminaciones de registros.

---
## Conexión a base de datos

La conexión a una base de datos es una parte fundamental en el desarrollo de aplicaciones que requieren almacenar y recuperar datos de manera eficiente y segura. En este artículo, exploraremos cómo establecer una conexión a una base de datos MySQL utilizando el lenguaje de programación Go (también conocido como Golang). A lo largo de este proceso, explicaremos cada paso con claridad para facilitar su comprensión.

### Paso 1: Instalación del controlador MySQL

Antes de comenzar, necesitamos asegurarnos de tener el controlador MySQL instalado en nuestro entorno de desarrollo de Go. Podemos hacer esto utilizando el comando `go get` seguido de la ruta del controlador:

```bash
go get -u github.com/go-sql-driver/mysql
```

Este comando descargará e instalará el controlador MySQL necesario para interactuar con la base de datos MySQL desde Go.

### Paso 2: Escribir el Código de Conexión

A continuación, escribiremos el código necesario para establecer una conexión a la base de datos MySQL en Go. Utilizaremos el paquete `database/sql` y el controlador MySQL que acabamos de instalar.

```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Cadena de conexión a la base de datos MySQL
	dsn := "usuario:contraseña@tcp(localhost:3306)/nombre_base_de_datos"

	// Abrir una conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar la conexión a la base de datos
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Conexión a la base de datos MySQL exitosa")
}
```

En este código:

- Importamos los paquetes necesarios: `database/sql` para la funcionalidad de la base de datos y `_ "github.com/go-sql-driver/mysql"` para el controlador MySQL.
- Definimos la cadena de conexión `dsn`, que incluye el nombre de usuario, contraseña, dirección del servidor y nombre de la base de datos.
- Abrimos una conexión a la base de datos utilizando `sql.Open()`. Es importante manejar cualquier error que ocurra en este paso.
- Verificamos la conexión a la base de datos utilizando `db.Ping()`. Si hay algún error, terminamos la aplicación con `log.Fatal()`.

En Go, cuando importas un paquete usando `_`, estás indicando al compilador que importe el paquete, pero no lo use directamente en tu código. Esto significa que el paquete se importará para que esté disponible para otras partes del código que sí lo necesiten, pero no se hará referencia directa a él en el propio código.

En el caso específico de `_ "github.com/go-sql-driver/mysql"`, el paquete `github.com/go-sql-driver/mysql` es el controlador MySQL para Go. Al importarlo con `_`, estás diciendo al compilador que importe el controlador para que esté disponible, pero no necesariamente lo estás utilizando explícitamente en tu código. El uso de `_` es útil cuando necesitas importar un paquete solo por su inicialización, como es común con los controladores de bases de datos en Go.

En resumen, `_ "github.com/go-sql-driver/mysql"` simplemente importa el paquete `github.com/go-sql-driver/mysql` para que el controlador MySQL esté disponible en tu programa, pero no necesitas hacer referencia explícita a él en tu código.

---
## Uso de godotenv
En este artículo, exploraremos cómo usar `godotenv`, una biblioteca de Go, para gestionar variables de entorno en una aplicación Go que se conecta a una base de datos MySQL. `godotenv` nos permite cargar variables de entorno desde un archivo `.env`, lo que facilita la configuración de nuestra aplicación y la ocultación de información sensible como credenciales de base de datos.

### Instalación de godotenv

El primer paso es instalar `godotenv` en nuestro entorno Go. Podemos hacerlo fácilmente utilizando el comando `go get`:

```bash
go get github.com/joho/godotenv
```

### Creación del archivo .env

A continuación, creamos un archivo `.env` en el directorio raíz de nuestro proyecto. Este archivo contendrá nuestras variables de entorno, como las credenciales de la base de datos:

```plaintext
DB_NAME=db_contacts
DB_USER=alexroel
DB_PASSWORD=123456
DB_HOST=localhost
DB_PORT=3306
```

En este ejemplo, hemos definido variables como el nombre de la base de datos, el usuario y la contraseña de MySQL, así como la dirección y el puerto del servidor MySQL.

### Conexión a la Base de Datos en Otro Paquete

Ahora, vamos a crear una función en un paquete `database` que maneje la conexión a la base de datos utilizando las variables de entorno cargadas desde el archivo `.env`. Aquí está el código:

```go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Cadena de conexión a la base de datos MySQL
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Abrir una conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar la conexión a la base de datos
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión a la base de datos MySQL exitosa")

	return db, nil
}
```

En este código, utilizamos `godotenv.Load()` para cargar las variables de entorno desde el archivo `.env`. Luego, construimos la cadena de conexión a la base de datos MySQL utilizando estas variables. Finalmente, abrimos una conexión a la base de datos y verificamos si la conexión fue exitosa.

### Prueba de Conexión

Para probar nuestra conexión a la base de datos, podemos escribir un programa principal que llame a nuestra función `Connect` y maneje cualquier error que ocurra. Aquí está el código de ejemplo:

```go
package main

import (
	"log"

	"tu_paquete/database"
	"tu_paquete/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
```

En este programa principal, llamamos a la función `Connect` del paquete `database` para establecer la conexión a la base de datos. Si hay algún error, imprimimos el error y finalizamos la ejecución del programa. 

Con `godotenv`, hemos simplificado la gestión de variables de entorno en nuestra aplicación Go, lo que la hace más segura y portátil.

---
## Listar contactos
En esta guía, aprenderemos a listar contactos almacenados en una base de datos MySQL utilizando modelos y controladores en un proyecto Go.

#### Paso 1: Crear el modelo de contacto

Dentro del paquete `models`, crea un archivo llamado `contact.go` con el siguiente contenido:

```go
package models

// Estructura de contacto
type Contact struct {
	Id    int
	Name  string
	Email string
	Phone string
}
```

Este modelo define la estructura de un contacto con campos como `Id`, `Name`, `Email` y `Phone`.

#### Paso 2: Crear el controlador para listar contactos

Dentro del paquete `handlers`, crea un archivo llamado `handlers.go` con el siguiente contenido:

```go
package handlers

import (
	"database/sql"
	"db-mysql/models"
	"fmt"
	"log"
)

// ListContacts lista todos los contactos desde la base de datos
func ListContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact;"

	// Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("---------------------------------------------------------------------------")
	for rows.Next() {
		// Instancia de modelo contact
		contact := models.Contact{}

		var valueEmail sql.NullString // Para manejar valor null
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		// Verificar si el valor es null o no
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo electrónico"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("---------------------------------------------------------------------------")
	}
}
```

Este controlador `ListContacts` toma una conexión de base de datos como argumento y ejecuta una consulta SQL para seleccionar todos los contactos. Luego, itera sobre los resultados y muestra cada contacto en la consola.

#### Paso 3: Probar desde el archivo `main.go`

Finalmente, podemos probar nuestro código desde el archivo `main.go`:

```go
package main

import (
	"db-mysql/database"
	"db-mysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Listar contactos
	handlers.ListContacts(db)
}
```

Este código crea una conexión a la base de datos utilizando la función `Connect` del paquete `database`, y luego llama a la función `ListContacts` del paquete `handlers` para listar todos los contactos desde la base de datos.

Con estos pasos, has creado un programa en Go que lista contactos desde una base de datos MySQL utilizando modelos y controladores. ¡Felicidades!

---
## Obtener un contacto
En esta guía, aprenderemos cómo obtener un contacto de una base de datos MySQL utilizando un controlador y un modelo en un proyecto Go.

#### Paso 1: Crear el controlador para obtener un contacto por ID

Dentro del paquete `handlers`, crea un archivo llamado `handlers.go` con el siguiente contenido:

```go
package handlers

import (
	"database/sql"
	"db-mysql/models"
	"fmt"
	"log"
)

...

// GetContactByID obtiene un contacto de la base de datos mediante su ID
func GetContactByID(db *sql.DB, contactID int) {
	// Consulta SQL para seleccionar un contacto por su ID
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	// Instancia de modelo contact
	contact := models.Contact{}
	var valueEmail sql.NullString // Para manejar valor null

	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("no se encontró ningún contacto con el ID %d", contactID)
		}
	}

	// Verificar si el valor es null o no
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo electrónico"
	}

	fmt.Println("\nLISTA DE UN CONTACTO:")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
		contact.ID, contact.Name, contact.Email, contact.Phone)
	fmt.Println("---------------------------------------------------------------------------")
}
```

En este código, hemos creado una función `GetContactByID` que acepta un parámetro adicional `contactID`, que es el ID del contacto que deseamos obtener de la base de datos. Esta función ejecuta una consulta SQL para seleccionar un contacto por su ID, y luego escanea el resultado en una estructura de contacto del paquete `models`. Si no se encuentra ningún contacto con el ID proporcionado, se imprime un mensaje de error.

#### Paso 2: Probar desde el archivo `main.go`

Ahora podemos probar nuestra función `GetContactByID` desde el archivo `main.go` o cualquier otro lugar donde deseemos obtener un contacto por su ID:

```go
package main

import (
	"db-mysql/database"
	"db-mysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Obtener un contacto por ID
	contactID := 3 // ID del contacto que deseamos obtener
	handlers.GetContactByID(db, contactID)
}
```

Este código crea una conexión a la base de datos utilizando la función `Connect` del paquete `database`, y luego llama a la función `GetContactByID` del paquete `handlers` para obtener un contacto por su ID. En este ejemplo, estamos obteniendo el contacto con ID `3`, pero puedes cambiar este valor según tus necesidades.

Con estos pasos, has creado un programa en Go que puede obtener un contacto de una base de datos MySQL por su ID. ¡Felicidades!

---
## Registrar contacto
En esta guía, aprenderemos cómo registrar un nuevo contacto en una base de datos MySQL utilizando un controlador en un proyecto Go.

#### Paso 1: Crear el controlador para registrar un nuevo contacto

Dentro del paquete `handlers`, crea una función llamada `CreateContact` en el archivo `handlers.go` con el siguiente contenido:

```go
package handlers

import (
	"database/sql"
	"log"

	"tu_paquete/models" // Asegúrate de importar el paquete models adecuadamente
)

// CreateContact registra un nuevo contacto en la base de datos
func CreateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQL para insertar un nuevo contacto
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}
```

En esta función `CreateContact`, aceptamos dos parámetros: `db`, que es la conexión a la base de datos, y `contact`, que es un objeto de tipo `models.Contact` que contiene los detalles del nuevo contacto. Luego, ejecutamos una sentencia SQL de inserción para agregar el nuevo contacto a la base de datos.

#### Paso 2: Probar la función desde el archivo `main.go`

Ahora puedes probar la función `CreateContact` desde tu archivo `main.go` o desde cualquier otro lugar donde desees registrar un nuevo contacto:

```go
package main

import (
	"db-mysql/database"
	"db-mysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear una instancia de Contact
	newContact := models.Contact{
		Name:  "Nuevo Usuario",
		Email: "nuevo@example.com",
		Phone: "123456789",
	}

	// Registrar un nuevo contacto
	handlers.CreateContact(db, newContact)
}
```

Este código crea una conexión a la base de datos utilizando la función `Connect` del paquete `database`, y luego crea una nueva instancia de `models.Contact` con los detalles del nuevo contacto. Luego, llama a la función `CreateContact` del paquete `handlers` para registrar el nuevo contacto en la base de datos.

Con estos pasos, has creado una función en Go que registra un nuevo contacto en una base de datos MySQL. ¡Felicidades!

---
## Editar contacto
En esta guía, aprenderemos cómo editar un contacto existente en una base de datos MySQL utilizando un controlador en un proyecto Go.

#### Paso 1: Crear el controlador para editar un contacto

Dentro del paquete `handlers`, crea una función llamada `UpdateContact` en el archivo `handlers.go` con el siguiente contenido:

```go
package handlers

// UpdateContact actualiza un contacto existente en la base de datos
func UpdateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQL para actualizar un contacto
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto actualizado con éxito")
}
```

En esta función `UpdateContact`, aceptamos dos parámetros: `db`, que es la conexión a la base de datos, y `contact`, que es un objeto de tipo `models.Contact` que contiene los detalles actualizados del contacto. Luego, ejecutamos una sentencia SQL de actualización para modificar el contacto existente en la base de datos.

#### Paso 2: Probar la función desde el archivo `main.go`

Ahora puedes probar la función `UpdateContact` desde tu archivo `main.go` o desde cualquier otro lugar donde desees editar un contacto existente:

```go
package main

import (
	"db-mysql/database"
	"db-mysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear una instancia de Contact con los detalles actualizados
	updatedContact := models.Contact{
		ID:    4, // ID del contacto que deseas actualizar
		Name:  "Alex",
		Email: "alex@example.com",
		Phone: "987654321",
	}

	// Actualizar el contacto en la base de datos
	handlers.UpdateContact(db, updatedContact)
}
```

Este código crea una conexión a la base de datos utilizando la función `Connect` del paquete `database`, y luego crea una nueva instancia de `models.Contact` con los detalles actualizados del contacto. Luego, llama a la función `UpdateContact` del paquete `handlers` para editar el contacto existente en la base de datos.

Con estos pasos, has creado una función en Go que edita un contacto existente en una base de datos MySQL. ¡Felicidades!

---
## Eliminar contacto
En esta guía, aprenderemos cómo eliminar un contacto de una base de datos MySQL utilizando un controlador en un proyecto Go.

#### Paso 1: Crear el controlador para eliminar un contacto

Dentro del paquete `handlers`, crea una función llamada `DeleteContact` en el archivo `handlers.go` con el siguiente contenido:

```go
package handlers

// DeleteContact elimina un contacto de la base de datos
func DeleteContact(db *sql.DB, contactID int) {
	// Sentencia SQL para eliminar un contacto por su ID
	query := "DELETE FROM contact WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto eliminado con éxito")
}
```

En esta función `DeleteContact`, aceptamos dos parámetros: `db`, que es la conexión a la base de datos, y `contactID`, que es el ID del contacto que deseamos eliminar. Luego, ejecutamos una sentencia SQL de eliminación para eliminar el contacto de la base de datos.

#### Paso 2: Probar la función desde el archivo `main.go`

Ahora puedes probar la función `DeleteContact` desde tu archivo `main.go` o desde cualquier otro lugar donde desees eliminar un contacto:

```go
package main

import (
	"db-mysql/database"
	"db-mysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ID del contacto que deseas eliminar
	contactID := 5

	// Eliminar el contacto de la base de datos
	handlers.DeleteContact(db, contactID)

	// Listar contactos (opcional, solo para verificar que el contacto fue eliminado)
	handlers.ListContacts(db)
}
```

Este código crea una conexión a la base de datos utilizando la función `Connect` del paquete `database`, y luego llama a la función `DeleteContact` del paquete `handlers` para eliminar el contacto con el ID proporcionado de la base de datos. Luego, opcionalmente, lista los contactos restantes para verificar que el contacto fue eliminado correctamente.

Con estos pasos, has creado una función en Go que elimina un contacto de una base de datos MySQL. ¡Felicidades! Si tienes alguna pregunta adicional o necesitas más ayuda, no dudes en preguntar.

---
## Terminar el proyecto
En esta guía, aprenderemos cómo interactuar con una base de datos MySQL utilizando un menú interactivo implementado en Go. Utilizaremos el paquete `bufio` para manejar la entrada del usuario y los controladores definidos en el proyecto para realizar operaciones CRUD en la base de datos.

#### Paso 1: Configurar la estructura del proyecto

Asegúrate de tener la estructura de tu proyecto configurada correctamente con los paquetes `database`, `handlers` y `models`. Los controladores CRUD (`ListContacts`, `GetContactByID`, `CreateContact`, `UpdateContact`, `DeleteContact`) deben estar implementados en el paquete `handlers`.

#### Paso 2: Implementar el menú interactivo

```go
package main

import (
	"bufio"
	"db-mysql/database"
	"db-mysql/handlers"
	"db-mysql/models"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opción seleccionada por el usuario
		var option int
		fmt.Scanln(&option)

		// Ejecutar la opción seleccionada
		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese el ID del contacto: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactByID(db, idContact)
		case 3:
			newContact := inputContactDetails()
			handlers.CreateContact(db, newContact)
		case 4:
			updatedContact := inputContactDetails()
			handlers.UpdateContact(db, updatedContact)
		case 5:
			fmt.Print("Ingrese el ID del contacto que quiere eliminar: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}

// Función para ingresar los detalles del contacto desde la entrada estándar
func inputContactDetails() models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electrónico del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el número de teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
```

En este código, hemos creado un menú interactivo que muestra diferentes opciones al usuario y espera la entrada del usuario para realizar las operaciones correspondientes en la base de datos. Utilizamos el paquete `bufio` para leer la entrada del usuario de manera más eficiente.

#### Paso 3: Explicación del paquete bufio

El paquete `bufio` proporciona funcionalidades para leer la entrada del usuario de manera eficiente. Algunos de los métodos más comunes proporcionados por `bufio` son:

- `NewReader`: Este método crea un nuevo lector que lee desde la entrada estándar.
- `Reader.ReadString`: Este método lee la entrada hasta que encuentra un delimitador especificado y devuelve la cadena leída.
- `Reader.ReadBytes`: Este método lee la entrada hasta que encuentra un byte especificado y devuelve una secuencia de bytes leída.
- `Reader.Scan`: Este método se utiliza junto con un escáner para leer la entrada del usuario en palabras o líneas.
- `Reader.ReadLine`: Este método lee una sola línea de la entrada del usuario y la devuelve como una secuencia de bytes junto con un indicador de final de línea.

En nuestro código, utilizamos `NewReader` para crear un lector que leerá la entrada del usuario desde `os.Stdin`. Luego, utilizamos `ReadString` para leer la entrada del usuario hasta que presione la tecla Enter y obtener la entrada como una cadena. También podríamos haber usado otros métodos como `ReadBytes`, `Scan`, o `ReadLine`, dependiendo de los requisitos específicos de nuestro programa.

Con esto, has implementado un menú interactivo en Go utilizando el paquete `bufio` y has explicado los métodos proporcionados por este paquete. ¡Felicidades! Ahora puedes interactuar con una base de datos MySQL de manera eficiente desde tu programa en Go.

---
## Resumen
En esta sección del curso de Golang, exploramos cómo trabajar con una base de datos MySQL para la persistencia de datos en nuestras aplicaciones. A lo largo de esta sección, aprendimos los fundamentos necesarios para interactuar con una base de datos MySQL desde una aplicación Go.

Comenzamos instalando y configurando MySQL en nuestro sistema para prepararnos para trabajar con bases de datos. Luego, creamos una base de datos utilizando comandos SQL básicos y establecimos una conexión entre nuestra aplicación Go y la base de datos MySQL para poder realizar operaciones de lectura y escritura.

Utilizamos la biblioteca `godotenv` para cargar variables de entorno desde un archivo `.env`, lo que nos permitió gestionar de manera segura y conveniente la configuración de la conexión a la base de datos.

Implementamos la funcionalidad para listar todos los contactos almacenados en la base de datos y aprendimos a recuperar un contacto específico utilizando su identificador único. También exploramos cómo registrar un nuevo contacto, actualizar la información de un contacto existente y eliminar un contacto de la base de datos cuando fuera necesario.

Finalmente, concluimos la sección construyendo un menú de navegación interactivo utilizando el paquete `bufio`, que permitió al usuario realizar todas las operaciones CRUD de manera intuitiva.

Al finalizar esta sección, estamos bien equipados para trabajar con bases de datos MySQL en nuestros proyectos Go y gestionar eficazmente la persistencia de datos.