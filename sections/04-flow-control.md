# Control de flujos

1. [Condición If](#Condición-If)
2. [Condición Switch](#Condición-Switch)
3. [Bucle For](#Bucle-For)


---
## Condición If
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
## Condición Switch 
Una forma mas corta de hacer multiples condiciones es con `switch`. 


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

Tambien podemos hacer como una busquedad de datos.

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

Bucle `for` normal con: 
- Inicialización 
- Condición 
- Imcremento 
~~~go
// Dibuja corazones
func drawHearts(lives int) string {
	var hearts string

	for i := 0; i < lives; i++ {
		hearts += "♥ "
	}
	return hearts
}
~~~

