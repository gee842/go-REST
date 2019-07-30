package main

import "time"

//Todo Structure
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos list of To Do(s)
type Todos []Todo
