/*
	Author: Albert Mensah-Ansah
	Title: GO Fundamentals (REST API)
	This project is supposed to help me get up to speed
	with basic CRUD in GO
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Todo Struct (Model)
type Todo struct {
	ID     int     `json:"id"`
	Task   string  `json:"task"`
	Status *Status `json:"status"`
}

type Status struct {
	ID         int    `json:"id"`
	StatusName string `json:"status_name"`
}

var todos []Todo

//Get all Todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

//Get a single Todo {id}
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params here ..
	for _, item := range todos {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			json.NewEncoder(w).Encode((item))
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

//Create a new todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = cap(todos)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

//update a book {id}
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range todos {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = id
			todos = append(todos, todo)
			return
		}
	}
	json.NewEncoder(w).Encode(todos)

}

//delete {id}
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func main() {
	//Initialize Router
	r := mux.NewRouter()

	//in memory data .. .
	// initialize todos with the slice todo struct
	todos = append(todos, Todo{ID: 1, Task: "Clean House", Status: &Status{ID: 1, StatusName: "Not Started"}})
	todos = append(todos, Todo{ID: 2, Task: "Make Bed Pen", Status: &Status{ID: 2, StatusName: "In Progress"}})
	todos = append(todos, Todo{ID: 3, Task: "Write that code", Status: &Status{ID: 3, StatusName: "Complete"}})

	// Router Handler / Endpooints
	r.HandleFunc("/api/todos", getTodos).Methods("GET")
	r.HandleFunc("/api/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/api/todos", createTodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/api/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
