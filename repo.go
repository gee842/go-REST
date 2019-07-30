package main

import "fmt"
import "log"

var todos Todos

func init() {
	//
	// RepoCreateTodo(Todo{Name: "Host meetup"})

	if err := Load("./todos.json", &todos); err != nil {
		RepoCreateTodo(Todo{Name: "My First Todo"})
	}

}

//RepoFindTodo Find Todo in db
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//RepoCreateTodo creates Todo
func RepoCreateTodo(t Todo) Todo {
	totalTodos := len(todos)
	if totalTodos == 0 {

		t.ID = 0
	} else {
		t.ID = todos[(len(todos)-1)].ID + 1
	}
	todos = append(todos, t)
	if err := Save("./todos.json", todos); err != nil {
		log.Fatalln(err)
	}
	return t
}

//RepoDestroyTodo removes Todo
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {

			todos = append(todos[:i], todos[i+1:]...)
			if err := Save("./todos.json", todos); err != nil {
				log.Fatalln(err)
			}
			return fmt.Errorf("Deleted Todo with id of %d", id)
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
