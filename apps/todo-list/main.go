package main

import (
	"bufio"
	"fmt"
	"os"
)

// Estructura de Tarea
type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

// Estructura para lista de teareas
type ListaTareas struct {
	tareas []Tarea
}

// Método para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Método para marcar como completado
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// Método para editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// Métdodo para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//Instancia de Lista de tareas
	lista := ListaTareas{}

	// Instancia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese el descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el índice de la tarea que desea actualizar:")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese el descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción inválida.")
		}

		// Listar todas las tareas
		fmt.Println("Lista de tareas:")
		fmt.Println("===================================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("===================================================")
	}
}
