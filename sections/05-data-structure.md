# Estructura de datos

1. [Array](#Array)
2. [Slices](#Slices)

---
## Array 
Las matrices se utilizan para almacenar múltiples valores del mismo tipo en una sola variable, en lugar de declarar variables separadas para cada valor.

Este ejemplo declara dos matrices con longitudes definidas:

~~~go
	//Declarar un array
	var array1 [3]int
	array1[0] = 10
	array1[1] = 20
	array1[2] = 30

	array2 := [5]int{40, 50, 60, 70, 80}

	fmt.Println(array1)
	fmt.Println(array2)
~~~

Este ejemplo declara una matriz con longitudes inferidas:

~~~go
	//Declarar un array
	array2 := [...]int{40, 50, 60, 70, 80}
	fmt.Println(array2)
~~~

Este ejemplo declara una matriz de cadenas:

~~~go
	//Declarar un array
	users := [3]string{"Juan", "Alex", "Pedro"}
	fmt.Println(users)
~~~

### Acceder y cambiar un elementos de una matriz
Puede acceder a un elemento de matriz específico consultando el número de índice.

~~~go
	//Declarar un array
	users := [3]string{"Juan", "Alex", "Pedro"}
	fmt.Println(users[1])
~~~

También puede cambiar el valor de un elemento de matriz específico haciendo referencia al número de índice.

~~~go
	//Declarar un array
	users := [3]string{"Juan", "Alex", "Pedro"}
	fmt.Println(users)

    //Modificando un elemento
    users[2] = "Roel"
    fmt.Println(users)
~~~


### Inicializar solo elementos específicos
Es posible inicializar solo elementos específicos en una matriz.

~~~go
	numbers := [5]int{2: 30, 4: 50}
	fmt.Println(numbers)
~~~

### Encuentra la longitud de una matriz
La función `len()` se utiliza para encontrar la longitud de una matriz:

~~~go
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(len(numbers))
~~~

---
## Crear Slice
- Los segmentos son similares a las matrices, pero son más potentes y flexibles.
- Al igual que las matrices, los sectores también se utilizan para almacenar múltiples valores del mismo tipo en una sola variable.
- Sin embargo, a diferencia de las matrices, la longitud de un segmento puede crecer y reducirse como mejor le parezca.


~~~go
	myslice := []int{100, 200, 300}
	fmt.Println(myslice)
~~~
- Función `len()`: devuelve la longitud del segmento (el número de elementos en el segmento)
- Función `cap()`: devuelve la capacidad de la porción (la cantidad de elementos a la que la porción puede crecer o reducirse)

~~~go
	myslice := []int{100, 200, 300}
	fmt.Println(myslice)
	fmt.Println(len(myslice)) //Logitud
	fmt.Println(cap(myslice)) //Capacidas
~~~

### Crear Slice de un Array 
Puede crear un segmento cortando una matriz:

~~~go
	//Un array
	arr1 := [6]int{10, 20, 30, 40, 50}

	myslice := arr1[1:4] //Crear slice

	fmt.Println("Myslice:", myslice)
	fmt.Println("Longitud:", len(myslice))
	fmt.Println("Capacidad:", cap(myslice))
~~~

### Crea una rebanada con la función make()
La función `make()` también se puede utilizar para crear un segmento.

- Sintaxis: `nombre_slice := make([]tipo, longitud, capacidad)`

~~~go
	myslice1 := make([]int, 5, 10)

	fmt.Println("Myslice:", myslice1)
	fmt.Println("Longitud:", len(myslice1))
	fmt.Println("Capacidad:", cap(myslice1))

	myslice2 := make([]int, 5)

	fmt.Println("Myslice:", myslice2)
	fmt.Println("Longitud:", len(myslice2))
	fmt.Println("Capacidad:", cap(myslice2))
~~~

---
## Modificar Slices


### Agregar mas elementos 
Para agregar mas elementos a un `slice` usaremos función `append()`. 

~~~go
	myslice := []int{100, 200, 300}

	myslice = append(myslice, 400, 500)
	fmt.Println(myslice)
    
	fmt.Println(len(myslice)) //Logitud
	fmt.Println(cap(myslice)) //Capacidas
~~~






