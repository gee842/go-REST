# go-REST
Simple REST API with Persistent Storage, in Golang

Database Logs stored inside of todos.json
Use case modelled after skeleton of a global To-Do list REST API Application.

To Run: 
  Required: Place github.com/gorilla/mux into your User or System src folder- eg. (~/go/src/mux)
  Inside the Directory, excecute     "go run ."
  Server Will be available at localhost:8080
  


Routes:

GET:

/                       - Simply Prints "Welcome!"

/todos                  - Lists all the stored To-Dos

/todos/Destroy/{id}     - Removes {id} from the database

/todos/{id}             - Displays Todo with ID {id}


POST:

/todos - Adds Entry to Database. 
Usage Example: curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
