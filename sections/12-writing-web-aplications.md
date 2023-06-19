# Escribir aplicaciones web

1. [Introducción](#introducción)
2. [Empezando](#empezando)
3. [El paquete net/http](#el-paquete-nethttp)
4. [Usar net/http](#usar-nethttp)
5. [El paquete html/template](#el-paquete-htmltemplate)
6. [Manejo de páginas inexistentes](#manejo-de-páginas-inexistentes)
7. [Guardar páginas](#guardar-páginas)
8. [Caché de plantillas](#caché-de-plantillas)
9. [Validación](#validación)
10. [Presentación de literales de funciones y clausuras](#presentación-de-literales-de-funciones-y-clausuras)
11. [Resumen](#resumen)

---
## Introducción 
En esta sección, vamos a explorar cómo escribir aplicaciones web en Go con una serie de temas interesantes. Comenzaremos por presentar el paquete net/http, que es la base para construir aplicaciones web en Go.

Usaremos el paquete net/http para servir páginas wiki. Aprenderemos cómo configurar un servidor HTTP básico y cómo manejar solicitudes y respuestas utilizando los manejadores (handlers) proporcionados por este paquete.

Luego, nos adentraremos en el paquete html/template, que nos permitirá generar y mostrar contenido dinámico en nuestras páginas web. Aprenderemos cómo cargar y utilizar plantillas HTML para renderizar contenido de manera flexible.

Un problema común que enfrentamos en el desarrollo web es el manejo de páginas inexistentes. Aprenderemos cómo redirigir al usuario a una página de edición cuando intenta acceder a una página que no existe.

Además, exploraremos cómo guardar las páginas editadas en el servidor. Implementaremos un formulario de edición y aprenderemos cómo guardar los cambios realizados por el usuario en un archivo en el servidor.

También abordaremos el tema del manejo de errores de manera adecuada. Aprenderemos a validar y manejar los errores que puedan ocurrir durante la ejecución de nuestra aplicación web.

Optimizaremos nuestra aplicación mediante la caché de plantillas. Aprenderemos cómo compilar y almacenar en memoria las plantillas HTML para mejorar el rendimiento de nuestra aplicación.

Por último, introduciremos los literales de funciones y clausuras. Veremos cómo utilizar estas características de Go para abstraer la funcionalidad repetitiva y simplificar nuestros controladores de manejo de solicitudes.

---
## Empezando
Crea un directorio nuevo para este tutorial dentro de tu GOPATH y cambia al directorio:
~~~bash
$ mkdir gowiki
$ cd gowiki
~~~
Crea un archivo llamado wiki.go, ábrelo en tu editor favorito y agrega las siguientes líneas:

~~~go
package main

import (
    "fmt"
    "os"
)
~~~
Importamos los paquetes fmt y os de la biblioteca estándar de Go. Más adelante, a medida que implementemos funcionalidades adicionales, agregaremos más paquetes a esta declaración de importación.

### Estructuras de datos
Comencemos definiendo las estructuras de datos. Una wiki consiste en una serie de páginas interconectadas, cada una de las cuales tiene un título y un cuerpo (el contenido de la página). Aquí definimos Page como una estructura con dos campos que representan el título y el cuerpo.
~~~go
type Page struct {
    Title string
    Body  []byte
}
~~~

El tipo []byte significa "una lista de bytes" (consulte Slices: uso e internos para obtener más información sobre las listas). El elemento Body es un []byte en lugar de una cadena (string) porque ese es el tipo que esperan las bibliotecas io que utilizaremos, como verás a continuación.

La estructura Page describe cómo se almacenarán los datos de la página en memoria. Pero, ¿qué hay de almacenamiento persistente? Podemos abordar eso creando un método save en Page:

~~~go
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}
~~~

La firma de este método dice: "Este es un método llamado save que toma como receptor p, un puntero a Page. No toma parámetros y devuelve un valor de tipo error".

Este método guardará el cuerpo (Body) de la página en un archivo de texto. Por simplicidad, utilizaremos el título como nombre de archivo.

El método save devuelve un valor de error porque ese es el tipo de retorno de WriteFile (una función de la biblioteca estándar que escribe una lista de bytes en un archivo). El método save devuelve el valor de error para permitir que la aplicación lo maneje en caso de que algo salga mal al escribir el archivo. Si todo va bien, Page.save() devolverá nil (el valor cero para punteros, interfaces y algunos otros tipos).

El número entero octal 0600, que se pasa como tercer parámetro a WriteFile, indica que el archivo debe crearse con permisos de lectura y escritura solo para el usuario actual. (Consulta la página de manual Unix open(2) para obtener más detalles).

Además de guardar páginas, también querremos cargar páginas:

~~~go
func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := os.ReadFile(filename)
    return &Page{Title: title, Body: body}
}
~~~

La función loadPage construye el nombre de archivo a partir del parámetro de título, lee el contenido del archivo en una nueva variable llamada body y devuelve un puntero a un Page literal construido con los valores de título y cuerpo adecuados.

Las funciones pueden devolver múltiples valores. La función de la biblioteca estándar os.ReadFile devuelve []byte y error. En loadPage, aún no se está gestionando el error; se utiliza el "identificador en blanco" representado por el guion bajo (_) para descartar el valor de retorno del error (en esencia, asignando el valor a nada).

Pero, ¿qué sucede si ReadFile encuentra un error? Por ejemplo, el archivo podría no existir. No debemos ignorar tales errores. Modifiquemos la función para que devuelva *Page y error.

~~~go
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
~~~

Los llamadores de esta función ahora pueden comprobar el segundo parámetro; si es nil, significa que se ha cargado correctamente una Page. Si no es nil, será un error que puede ser manejado por el llamador (consulta la especificación del lenguaje para obtener más detalles).

En este punto, tenemos una estructura de datos simple y la capacidad de guardar y cargar desde un archivo. Escribamos una función main para probar lo que hemos escrito:

~~~go
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("Esta es una página de muestra.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
~~~
Después de compilar y ejecutar este código, se creará un archivo llamado TestPage.txt que contendrá el contenido de p1. Luego, se leerá el archivo en la estructura p2 y se imprimirá el elemento Body en la pantalla.

Puedes compilar y ejecutar el programa de la siguiente manera:
~~~bash
$ go build wiki.go
$ ./wiki
Esta es una página de muestra.
~~~
(Si estás utilizando Windows, debes escribir "wiki" sin el "./" para ejecutar el programa).

---
## El paquete net/http
Aquí tienes un ejemplo completo y funcional de un servidor web simple:
~~~go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola! Me encantan los %s.", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
~~~

La función main comienza con una llamada a http.HandleFunc, que le indica al paquete http que maneje todas las solicitudes a la raíz del sitio web ("/") con el manejador (handler) definido.

Luego llama a http.ListenAndServe, especificando que debe escuchar en el puerto 8080 en cualquier interfaz (":8080"). (Por ahora, no te preocupes por su segundo parámetro, nil). Esta función se bloqueará hasta que se termine el programa.

ListenAndServe siempre devuelve un error, ya que solo retorna cuando ocurre un error inesperado. Para registrar ese error, envolvemos la llamada a la función con log.Fatal.

La función handler es del tipo http.HandlerFunc. Toma un http.ResponseWriter y un http.Request como argumentos.

Un valor http.ResponseWriter ensambla la respuesta del servidor HTTP; al escribir en él, enviamos datos al cliente HTTP.

Un http.Request es una estructura de datos que representa la solicitud HTTP del cliente. `r.URL.Path` es el componente de la ruta de la URL de la solicitud. El `[1:]` al final significa "crear una subsección de Path desde el primer carácter hasta el final". Esto elimina el "/" inicial del nombre de la ruta.

Si ejecutas este programa y accedes a la siguiente URL:

http://localhost:8080/monos
el programa mostrará una página que contiene:

¡Hola! Me encantan los monos.

---
## Usar net/http
Para utilizar el paquete net/http, debe importarse:
~~~go
import (
    "fmt"
    "os"
    "log"
    "net/http"
)
~~~
Creemos un manejador, viewHandler, que permitirá a los usuarios ver una página del wiki. Manejará las URL que comiencen con "/view/".
~~~go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
~~~

Nuevamente, observa el uso de _ para ignorar el valor de retorno del error de loadPage. Esto se hace aquí por simplicidad y generalmente se considera una mala práctica. Lo solucionaremos más adelante.

Primero, esta función extrae el título de la página de r.URL.Path, el componente de ruta de la URL de la solicitud. La ruta se redefine con [len("/view/"):] para eliminar el componente inicial "/view/" de la ruta de la solicitud. Esto se debe a que la ruta invariablemente comenzará con "/view/", que no forma parte del título de la página.

La función luego carga los datos de la página, formatea la página con una cadena de HTML simple y la escribe en w, el http.ResponseWriter.

Para utilizar este manejador, reescribamos nuestra función main para inicializar http utilizando viewHandler para manejar cualquier solicitud bajo la ruta /view/.

~~~go
func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
~~~

Creemos algunos datos de página (como test.txt), compilaremos nuestro código y probaremos servir una página del wiki.

Abre el archivo test.txt en tu editor y guarda la cadena "Hola mundo" (sin comillas) en él.
~~~bash
$ go build wiki.go
$ ./wiki
~~~
(Si estás utilizando Windows, debes escribir "wiki" sin el "./" para ejecutar el programa).

Con este servidor web en ejecución, al visitar http://localhost:8080/view/test se mostrará una página titulada "test" que contiene las palabras "Hola mundo".

---
## El paquete html/template
Un wiki no es un wiki sin la capacidad de editar páginas. Creemos dos nuevos manejadores: uno llamado editHandler para mostrar un formulario de "editar página" y otro llamado saveHandler para guardar los datos ingresados a través del formulario.

Primero, agreguémoslos a main():

~~~go
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    //http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
~~~

La función editHandler carga la página (o, si no existe, crea una estructura Page vacía) y muestra un formulario HTML.

~~~go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Editando %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Guardar\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}
~~~

Esta función funcionará correctamente, pero todo ese HTML codificado de forma rígida es feo. Por supuesto, hay una mejor manera.

### Uso del paquete html/template
El paquete html/template es parte de la biblioteca estándar de Go. Podemos usar html/template para mantener el HTML en un archivo separado, lo que nos permite cambiar el diseño de nuestra página de edición sin modificar el código Go subyacente.

Primero, debemos agregar html/template a la lista de importaciones. Tampoco usaremos fmt, así que debemos eliminar eso.
~~~go
import (
    "html/template"
    "os"
    "net/http"
)
~~~

Creemos un archivo de plantilla que contenga el formulario HTML. Abre un archivo nuevo llamado `edit.html` y agrega las siguientes líneas:

~~~html
<h1>Editando {{.Title}}</h1>
<form action="/save/{{.Title}}" method="POST">
    <div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
    <div><input type="submit" value="Guardar"></div>
</form>
~~~

Modifica editHandler para usar la plantilla en lugar del HTML codificado:

~~~go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}
~~~

La función template.ParseFiles leerá el contenido de edit.html y devolverá un *template.Template.

El método t.Execute ejecuta la plantilla, escribiendo el HTML generado en http.ResponseWriter. Los identificadores punteados .Title y .Body se refieren a p.Title y p.Body.

Las directivas de plantilla están encerradas en dobles llaves. La instrucción printf "%s" .Body es una llamada a función que imprime .Body como una cadena en lugar de una secuencia de bytes, al igual que una llamada a fmt.Printf. El paquete html/template ayuda a garantizar que solo se genere HTML seguro y correcto mediante acciones de plantilla. Por ejemplo, automáticamente escapa cualquier signo mayor que (>), reemplazándolo por >, para asegurarse de que los datos del usuario no corrompan el HTML del formulario.

Como ahora estamos trabajando con plantillas, creemos una plantilla para nuestro viewHandler llamada `view.html`:

~~~html
<h1>{{.Title}}</h1>
<p>[<a href="/edit/{{.Title}}">editar</a>]</p>
<div>{{printf "%s" .Body}}</div>
~~~
Modifica viewHandler en consecuencia:

~~~go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}
~~~

Observa que hemos utilizado casi exactamente el mismo código de plantilla en ambos manejadores. Eliminemos esta duplicación moviendo el código de plantilla a su propia función:

~~~go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}
~~~

Y modifiquemos los manejadores para usar esa función:

~~~go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
~~~

Si comentamos el registro de nuestro manejador de guardado no implementado en main, podemos compilar y probar nuestro programa nuevamente.

---
## Manejo de páginas inexistentes

¿Qué sucede si visitas /view/UnaPáginaQueNoExiste? Verás una página que contiene HTML. Esto se debe a que ignora el valor de retorno de error de loadPage y continúa intentando completar la plantilla sin datos. En cambio, si la página solicitada no existe, debería redirigir al cliente a la página de edición para que se pueda crear el contenido:

~~~go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
~~~
La función http.Redirect agrega un código de estado HTTP de http.StatusFound (302) y una cabecera de ubicación (Location) a la respuesta HTTP.

---
## Guardar páginas
La función saveHandler se encargará de procesar el envío de formularios ubicados en las páginas de edición. Después de descomentar la línea relacionada en main, implementemos el controlador:
~~~go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
~~~
El título de la página (proporcionado en la URL) y el único campo del formulario, Body, se almacenan en una nueva estructura Page. Luego se llama al método save() para escribir los datos en un archivo y se redirige al cliente a la página /view/.

El valor devuelto por FormValue es de tipo string. Debemos convertir ese valor a []byte antes de que encaje en la estructura Page. Utilizamos []byte(body) para realizar la conversión.

---
### Manejo de errores
Hay varios lugares en nuestro programa donde se ignoran los errores. Esto es una mala práctica, principalmente porque cuando ocurre un error, el programa tendrá un comportamiento inesperado. Una solución mejor es manejar los errores y devolver un mensaje de error al usuario. De esta manera, si algo sale mal, el servidor funcionará exactamente como queremos y se le notificará al usuario.

Primero, vamos a manejar los errores en renderTemplate:

~~~go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
~~~

La función http.Error envía un código de respuesta HTTP específico (en este caso, "Error interno del servidor") y un mensaje de error. Ya podemos ver que la decisión de poner esto en una función separada está dando sus frutos.

Ahora, arreglemos saveHandler:
~~~go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
~~~
Cualquier error que ocurra durante p.save() se informará al usuario.

---
## Caché de plantillas
Hay una ineficiencia en este código: renderTemplate llama a ParseFiles cada vez que se renderiza una página. Un enfoque mejor sería llamar a ParseFiles una vez en la inicialización del programa, para analizar todas las plantillas en un solo *Template. Luego podemos usar el método ExecuteTemplate para renderizar una plantilla específica.

Primero creamos una variable global llamada templates y la inicializamos con ParseFiles.

~~~go
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
~~~

La función template.Must es un envoltorio conveniente que genera un panic cuando se le pasa un valor de error que no es nulo, y de lo contrario devuelve el *Template sin cambios. Un panic es apropiado aquí; si las plantillas no se pueden cargar, lo único sensato que podemos hacer es salir del programa.

La función ParseFiles toma cualquier número de argumentos de tipo string que identifican nuestros archivos de plantilla, y analiza esos archivos en plantillas con nombres basados en el nombre base del archivo. Si agregáramos más plantillas a nuestro programa, agregaríamos sus nombres a los argumentos de la llamada a ParseFiles.

Luego modificamos la función renderTemplate para llamar al método templates.ExecuteTemplate con el nombre de la plantilla adecuada:
~~~go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
~~~

Ten en cuenta que el nombre de la plantilla es el nombre del archivo de la plantilla, por lo que debemos concatenar ".html" al argumento tmpl.

---
## Validación
Como habrás observado, este programa tiene una grave falla de seguridad: un usuario puede proporcionar una ruta arbitraria para ser leída/escrita en el servidor. Para mitigar esto, podemos escribir una función para validar el título con una expresión regular.

Primero, agrega "regexp" a la lista de importaciones. Luego podemos crear una variable global para almacenar nuestra expresión de validación:

~~~go
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
~~~
La función regexp.MustCompile analizará y compilará la expresión regular, y devolverá un regexp.Regexp. MustCompile es diferente de Compile en que generará un panic si falla la compilación de la expresión, mientras que Compile devuelve un error como segundo parámetro.

Ahora, escribamos una función que utilice la expresión validPath para validar la ruta y extraer el título de la página:
~~~go
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Título de página inválido")
    }
    return m[2], nil // El título es la segunda subexpresión.
}
~~~

Si el título es válido, se devolverá junto con un valor de error nulo. Si el título es inválido, la función escribirá un error "404 Not Found" en la conexión HTTP y devolverá un error al controlador. Para crear un nuevo error, tenemos que importar el paquete errors.

Coloquemos una llamada a getTitle en cada uno de los controladores:

---
## Presentación de literales de funciones y clausuras
La captura de la condición de error en cada controlador introduce mucho código repetido. ¿Qué pasaría si pudiéramos envolver cada uno de los controladores en una función que haga esta validación y verificación de errores? Los literales de función de Go proporcionan un medio poderoso de abstraer la funcionalidad que puede ayudarnos aquí.

Primero, reescribimos la definición de función de cada uno de los controladores para que acepten una cadena de título:

~~~go
func viewHandler(w http.ResponseWriter, r *http.Request, title string)
func editHandler(w http.ResponseWriter, r *http.Request, title string)
func saveHandler(w http.ResponseWriter, r *http.Request, title string)
~~~
Ahora, definamos una función envolvente que tome una función del tipo mencionado anteriormente y devuelva una función del tipo http.HandlerFunc (adecuada para ser pasada a la función http.HandleFunc):

~~~go
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
    // Aquí extraeremos el título de la página de la solicitud
    // y llamaremos al controlador proporcionado 'fn'
    }
}
~~~
La función devuelta se llama clausura porque encierra valores definidos fuera de ella. En este caso, la variable fn (el único argumento de makeHandler) está encerrada por la clausura. La variable fn será uno de nuestros controladores save, edit o view.

Ahora podemos tomar el código de getTitle y usarlo aquí (con algunas modificaciones menores):

~~~go
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}
~~~
La clausura devuelta por makeHandler es una función que toma un http.ResponseWriter y un http.Request (en otras palabras, un http.HandlerFunc). La clausura extrae el título de la ruta de la solicitud y lo valida con la expresión regular validPath. Si el título no es válido, se escribirá un error en el ResponseWriter utilizando la función http.NotFound. Si el título es válido, se llamará a la función controladora fn encerrada con el ResponseWriter, la Request y el título como argumentos.

Ahora podemos envolver las funciones controladoras con makeHandler en main, antes de que se registren con el paquete http:
~~~go
func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
~~~
Finalmente, eliminamos las llamadas a getTitle de las funciones controladoras, lo que las simplifica considerablemente:

---
## Resumen 
En esta sección, exploramos cómo escribir aplicaciones web en Go con una serie de temas interesantes. Comenzamos presentando el paquete net/http, que es la base para construir aplicaciones web en Go.

Utilizamos el paquete net/http para servir páginas wiki y aprendimos a configurar un servidor HTTP básico. Aprendimos cómo manejar solicitudes y respuestas utilizando los manejadores proporcionados por este paquete.

Luego nos adentramos en el paquete html/template, que nos permitió generar y mostrar contenido dinámico en nuestras páginas web. Aprendimos a cargar y utilizar plantillas HTML para renderizar contenido de manera flexible.

Abordamos el problema del manejo de páginas inexistentes y aprendimos a redirigir al usuario a una página de edición cuando intenta acceder a una página que no existe.

Además, exploramos cómo guardar las páginas editadas en el servidor. Implementamos un formulario de edición y aprendimos a guardar los cambios realizados por el usuario en un archivo en el servidor.

También abordamos el tema del manejo de errores de manera adecuada. Aprendimos a validar y manejar los errores que pueden ocurrir durante la ejecución de nuestra aplicación web.

Optimizamos nuestra aplicación mediante la caché de plantillas. Aprendimos a compilar y almacenar en memoria las plantillas HTML para mejorar el rendimiento de nuestra aplicación.

Por último, introdujimos los literales de funciones y clausuras. Aprendimos a utilizar estas características de Go para abstraer la funcionalidad repetitiva y simplificar nuestros controladores de manejo de solicitudes.