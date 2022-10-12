# Control de flujos

1. [Declaración IF](#Declaración-SI)
2. [Declaración Switch](#Declaración-Switch)
3. [Bucle For](#Bucle-For)
4. [Funciones](#Funciones)
5. [Uso de break, continue y defer](#Uso-de-break,-continue-y-defer)
   
---
## Declaración SI
Las condiciones ayundan con la lógica de programación, con las condiciones puede tomar deciones nuestras aplicaciones dependiendo de la expreciones que pongamos. 

~~~go
	//t := time.Now()

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Buenos dias!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Buenas tardes!")
	} else {
		fmt.Println("¡Buenas noches!")
	}
~~~

---
## Declaración Switch
- Utilice la instrucción `switch` para seleccionar uno de los muchos bloques de código que se ejecutarán.

Usando condiciones.
~~~go
    //t := time.Now()
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("¡Buenos dias!")
	case t.Hour() < 17:
		fmt.Println("¡Buenas tardes!")
	default:
		fmt.Println("¡Buenas noches!")
	}
~~~

- Tambien podemos hacer como una busquedad de datos y multibusqueda.

~~~go
	//os := runtime.GOOS;
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Go run -> Linux.")
	case "windows":
		fmt.Println("Go run -> Windows.")
	default:
		fmt.Println("Go run -> OS X.")
	}
~~~

---
## Bucle For
- El ciclo `for` recorre un bloque de código un número específico de veces.
- El bucle `for` es el único bucle disponible en Go.

Bucle `for` normal con: 
- Inicialización 
- Evaluación de una expresión 
- Imcremento 
~~~go

  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
~~~

### Bucle con una expresión y infinito
- En otros lenguajes se usan bucle `while` para trabajar con una expresión, pero en Go solo se trabaja con `for`. 
- Tambien para realizar un bucle infinito

~~~go

// While en go
  for codicion {
    fmt.Println("Hola")
  }

  //Bucle infinito
  for {
    fmt.Println("Hola")
  }
~~~


---
## Funciones
- Una función contiene un bloque de código que se ejecuta solo cuando llamas la función.
- Una función puede ejecutar el bloque de código que tenga y devolver un valor.
- Una función puede tomar cero o más argumentos.
- Observe que el tipo viene después del nombre de la variable.
 
~~~go
//Creando mis funciones
func play(name string) {
    //fmt.Println("Hola desde la función")
    fmt.Println("Hola ", name, "esta jugando")
}
 
func sum(a, b int) int {
    return a + b
}
 
func main() {
    //llmar una función
    play("Alex")
    fmt.Println("La suma es: ", sum(40, 50))
}
~~~

---
## Uso de break, continue y defer

### La declaración de ruptura
La declaración `break` se utiliza para interrumpir/terminar la ejecución del ciclo.

~~~go
for i:=0; i < 10; i++ {
	if i == 5 {
      break
    }
   fmt.Println(i)
}
~~~

### La declaración de continuación
La declaración `continue` se usa para omitir una o más iteraciones en el ciclo. Luego continúa con la siguiente iteración en el bucle.
~~~go
for i:=0; i < 10; i++ {
	if i == 5 {
      continue
    }
   fmt.Println(i)
}
~~~

### La declaración de aplazar 
La sentencia `defer` en Go mueve el código al final de la función, si se quiere ver así, aunque en realidad lo ejecuta antes de que la función se termine, sin importar el punto en donde se hace el return.

~~~go
	defer fmt.Println("Hola Alex")
	fmt.Println("Hola Mundo")
~~~