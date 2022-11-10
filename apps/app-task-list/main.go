package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) appendTask(tl *task) {

	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	name string
	desc string //Descripción
	comp bool   //Completado
}

//Marcar tarea completada
func (t *task) markComplete() {
	t.comp = true
}

//Editar descripción
func (t *task) editDescription(desc string) {
	t.desc = desc
}

//Editar el nombre
func (t *task) editName(name string) {
	t.name = name
}

func main() {
	fmt.Println("Mi Lista de Taerea")

	t := task{
		name: "Curso de Go",
		desc: "Completar mi tarea en esta semana",
		comp: false,
	}

	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %v\n", t.name, t.desc, t.comp)

	t.editName("Curso de Golang")
	t.editDescription("Terminar en una semana")
	t.markComplete()

	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %v\n", t.name, t.desc, t.comp)

}

/*
type taskList struct {
	tasks []*task
}

func (t *taskList) appendTask(tl *task) {

	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markAsCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func main() {
	t1 := &task{
		name:        "Terminar Curso de Go",
		description: "Terminar el Curso de Go en Platzi en las proximas dos semanas",
	}
	t2 := &task{
		name:        "Terminar Curso de JavaScript",
		description: "Terminar mi curso de Async/Await en JavaScript",
	}
	list := &taskList{
		tasks: []*task{
			t1, t2,
		}Println()
	}
	fmt.Printf("%+v\n", *list.tasks[0])
	fmt.Printf("%+v\n", *list.tasks[1])
	list.tasks[1].markAsCompleted()
	fmt.Printf("%+v\n", *list.tasks[1])
	t3 := &task{
		name:        "Construir mi propio servidor web",
		description: "Construir mi propio web server utilizando Go",
	}
	list.appendTask(t3)
	fmt.Printf("%+v\n", *list.tasks[2])
	list.removeTask(1)
	fmt.Printf("%+v\n", *list.tasks[1])
}
*/
