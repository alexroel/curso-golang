# Programación orientada a objetos en Go 
1. [POO en Go](#poo-en-go) 
2. [Estructuras y tipos](#estructuras-y-tipos) 
3. [Métodos y cosntructores](#métodos-y-cosntructores) 
4. [Encapsulamientos](#encapsulamientos)
5. [Composicion](#composicion)
7. [Polimorfismo](#polimorfismo)

---
## POO en Go 
La programación orientada a objetos (POO) en Go es un enfoque que utiliza estructuras y métodos para modelar objetos del mundo real en el código. Aunque Go es un lenguaje de programación que no se considera totalmente orientado a objetos, ofrece varias características de POO, como estructuras, métodos, interfaces y composición.

A continuación, se describen algunas características de POO en Go:

- **Estructuras**: las estructuras en Go son similares a las clases en otros lenguajes de programación orientados a objetos. Las estructuras se utilizan para representar objetos y tienen campos para almacenar datos.

- **Métodos**: los métodos en Go son funciones que operan en una estructura en particular. Se utilizan para agregar comportamiento a las estructuras y para definir operaciones específicas para un objeto.

- **Interfaces**: las interfaces en Go son tipos abstractos que definen un conjunto de métodos. Permiten que diferentes tipos implementen el mismo conjunto de métodos, lo que permite la polimorfismo y la abstracción.

- **Composición**: Go utiliza la composición para modelar relaciones entre objetos. Esto implica la creación de una estructura que contiene otras estructuras.

- **Encapsulamiento**: aunque Go no tiene modificadores de acceso públicos o privados, se puede lograr el encapsulamiento utilizando convenciones de nomenclatura de variables y estructuras.

En resumen, aunque Go no es un lenguaje orientado a objetos puro, puede utilizarse para escribir código orientado a objetos utilizando estructuras, métodos, interfaces, composición y encapsulamiento. La programación orientada a objetos en Go es útil para modelar objetos del mundo real y para crear código modular y reutilizable.

---
## Estructuras y tipos 
En Go, en lugar de utilizar el concepto de clases y objetos como en otros lenguajes orientados a objetos, se utilizan estructuras y tipos para representar entidades. Las estructuras en Go son similares a las clases en otros lenguajes de programación orientados a objetos, pero con algunas diferencias.

Una estructura en Go es un tipo de datos que contiene un conjunto de campos, cada uno con su propio tipo de datos. Los campos pueden ser de cualquier tipo de datos, incluyendo otros tipos de estructuras. Las estructuras en Go se definen utilizando la palabra clave `type`, seguida del nombre de la estructura y la definición de los campos que contendrá.

Por ejemplo, la siguiente estructura representa un libro:

~~~go
type Book struct {
    Title  string
    Author string
    Pages  int
}
~~~

Esta estructura define un tipo de datos llamado `Book` con tres campos: `Title`, `Author` y `Pages`. Cada campo es de un tipo de datos diferente: `Title` y `Author` son cadenas de caracteres (string), y `Pages` es un entero (int).

Para crear una instancia de una estructura en Go, se utiliza la palabra clave `var` seguida del nombre de la variable, seguida del operador `= `y la definición de la estructura utilizando la sintaxis de llaves.

Por ejemplo, para crear un libro con el título "Moby Dick", el autor "Herman Melville" y 300 páginas, se puede escribir:

~~~go
var myBook Book = Book{
    Title: "Moby Dick",
    Author: "Herman Melville",
    Pages: 300,
}
~~~

En este ejemplo, se crea una variable llamada `myBook` del tipo `Book`, y se inicializa con los valores de título, autor y número de páginas utilizando la sintaxis de llaves y los nombres de los campos.

---
## Métodos y cosntructores 
### Métodos
En Go, los métodos son funciones que operan en un tipo específico, como una estructura. Los métodos se definen utilizando la palabra clave `func`, seguida del nombre del método, los argumentos y el tipo al que se aplica el método.

Por ejemplo, podemos agregar un método `PrintInfo` a la estructura `Book` que imprima en pantalla el título, el autor y el número de páginas del libro:

~~~go
func (b *Book) PrintInfo() {
    fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}
~~~

En este ejemplo, el método `PrintInfo` tiene como receptor (receiver) la estructura `Book`, lo que significa que se puede llamar en cualquier instancia de la estructura `Book`. La sintaxis func (b *Book) antes del nombre del método indica que b es una variable de tipo Book que se utiliza dentro del método.

Para llamar a un método en una estructura, se utiliza la notación de punto (.) seguida del nombre del método y los argumentos necesarios. Por ejemplo, para llamar al método PrintInfo en el libro `myBook`, se escribe:

~~~go
myBook.PrintInfo()
~~~
### Cosntructores
En cuanto a los constructores en Go, no hay un concepto de constructor tradicional como en otros lenguajes de programación orientados a objetos. En lugar de eso, se suelen crear funciones que devuelven instancias de una estructura, inicializadas con los valores necesarios.

Por ejemplo, podemos crear una función `NewBook` que devuelva una instancia de la estructura `Book`, inicializada con un título, autor y número de páginas dados:

~~~go
func NewBook(title string, author string, pages int) *Book {
    return &Book{
        Title:  title,
        Author: author,
        Pages:  pages,
    }
}
~~~

Esta función toma tres argumentos, un título, un autor y un número de páginas, y devuelve una instancia de la estructura `Book` inicializada con estos valores. Para utilizar esta función, simplemente la llamamos como una función normal:

~~~go
myBook := NewBook("Moby Dick", "Herman Melville", 300)
~~~

En este ejemplo, se llama a la función `NewBook` con los argumentos necesarios, y la función devuelve una instancia de la estructura `Book`, que se asigna a la variable `myBook`.

---
## Encapsulamientos
En Go, la encapsulación se logra utilizando la primera letra del nombre de un identificador para indicar su visibilidad. Si un identificador comienza con una letra mayúscula, se considera exportado y puede ser accedido desde fuera del paquete. Si un identificador comienza con una letra minúscula, es privado y solo puede ser accedido dentro del paquete.

En el siguiente ejemplo, la estructura Book tiene tres campos: Title, Author y Pages, todos ellos exportados porque comienzan con una letra mayúscula. Esto significa que cualquier paquete que importe este paquete puede acceder y modificar estos campos.

~~~go
type Book struct {
    Title  string
    Author string
    Pages  int
}
~~~

Para hacer que los campos de la estructura Book sean privados y no puedan ser accedidos desde fuera del paquete, podemos hacer uso de la convención de nomenclatura de Go y cambiar la primera letra del nombre de los campos a minúscula.

~~~go
type Book struct {
    title  string
    author string
    pages  int
}
~~~

En este caso, los campos title, author y pages son privados y solo pueden ser accedidos desde dentro del paquete que los define. Si otro paquete intenta acceder a estos campos, Go devolverá un error de compilación. Para proporcionar acceso controlado a estos campos, podemos definir funciones que devuelvan o modifiquen los valores de los campos de la estructura Book. Por ejemplo:

~~~go
func (b *Book) SetTitle(title string) {
    b.title = title
}

func (b *Book) GetTitle() string {
    return b.title
}
~~~

Con estas funciones, podemos cambiar o leer el título del libro desde fuera del paquete de la siguiente manera:

~~~go
book := Book{title: "El Gran Gatsby", author: "F. Scott Fitzgerald", pages: 180}
book.SetTitle("El Gran Gatsby (Edición Especial)")
fmt.Println(book.GetTitle()) // Imprime: El Gran Gatsby (Edición Especial)
~~~
De esta manera, hemos logrado encapsular los campos de la estructura Book y proporcionar acceso controlado a ellos desde fuera del paquete.

---
## Composicion
La composición es un mecanismo que permite crear nuevas estructuras en Go que combinan los campos y métodos de una o más estructuras existentes. Esto se logra mediante la incrustación de una estructura dentro de otra.

La composición es una alternativa a la herencia, que es un concepto utilizado en muchos otros lenguajes de programación orientados a objetos. A diferencia de la herencia, que permite que una clase herede campos y métodos de una clase base, la composición en Go permite crear nuevas estructuras que contienen campos y métodos de otras estructuras.


~~~go
type Textbook struct {
	Book
	editorial string
	level     string
}

func NewTextbook(title string, author string, pages int, editorial string, level string) *Textbook {
	return &Textbook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

// Método para imprimer información
func (b *Textbook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}
~~~

El código define una estructura `Textbook` que es una composición de la estructura `Book`, a la que se le agregan los atributos editorial y level.

La función `NewTextbook` es un constructor que recibe los parámetros necesarios para crear una instancia de Textbook, crea un objeto `Book` con los parámetros de título, autor y páginas, y lo agrega a la estructura `Textbook`. Luego, asigna los valores de los atributos editorial y level y retorna un puntero a la estructura `Textbook` creada.

Finalmente, el método PrintInfo se define en la estructura `Textbook` y permite imprimir la información del libro, incluyendo el título, autor, páginas, editorial y nivel.

---
## Polimorfismo

El polimorfismo es la capacidad de los objetos de diferentes clases para responder al mismo mensaje o método. En Go, se puede implementar el polimorfismo a través de interfaces.

Una interfaz es un conjunto de métodos que una estructura debe implementar para satisfacer la interfaz. Esto permite que una estructura implemente múltiples interfaces y, por lo tanto, tenga diferentes comportamientos en diferentes contextos.

En Go, el polimorfismo se puede lograr mediante interfaces. Podemos definir una interfaz común para que tanto la estructura Book como la estructura Textbook implementen. Luego, podemos crear una función que acepte un parámetro de esta interfaz y pueda manejar ambos tipos de estructuras.

Aquí hay un ejemplo de cómo se puede lograr el polimorfismo para este código:

~~~go
type Printable interface {
    PrintInfo()
}

func Print(p Printable) {
    p.PrintInfo()
}

func main() {

    var myBook = NewBook("Pedro Páramo", "Juan Rulfo", 128)
	//myBook.PrintInfo()

    var myTextBook = NewTextbook("Comunicación", "Jaime Gamarra", 261, "Santillana SAC", "Secundaria")
	//myTextBook.PrintInfo()
    
    Print(bomyBookok)
    Print(myTextBook)
}
~~~

En este ejemplo, hemos definido una interfaz Printable que tiene un solo método PrintInfo(). Ambas estructuras Book y Textbook implementan esta interfaz, ya que ambas tienen el método PrintInfo(). Luego hemos creado una función Print() que toma un parámetro de tipo Printable y llama al método PrintInfo() en el objeto proporcionado. Finalmente, hemos creado un objeto book de tipo Book y un objeto textbook de tipo Textbook, y los hemos pasado como argumentos a la función Print().

Al compilar y ejecutar este código, podemos ver que ambos objetos son manejados correctamente por la función Print(), lo que demuestra el polimorfismo.
