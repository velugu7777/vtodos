package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

type Todo struct {
	Id     int
	Name   string
	IsDone bool
}

var mytodos []Todo

func init() {
	t1 := Todo{
		Id:     1,
		Name:   "Reading books",
		IsDone: false,
	}
	t2 := Todo{
		Id:     2,
		Name:   "Playing Cricket",
		IsDone: false,
	}
	mytodos = append(mytodos, t1, t2)
}

func main() {
	goji.Get("/todos", todos)
	goji.Post("/todos/newTodo", newTodo)
	goji.Get("/todos/:id", getTodo)
	goji.Put("/todos/:id", putTodo)
	goji.Delete("/todos/:id", delTodo)
	goji.Serve()
}

func todos(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%+v", mytodos)
}

func newTodo(w http.ResponseWriter, req *http.Request) {

}

func getTodo(c web.C, w http.ResponseWriter, req *http.Request) {

}

func putTodo(c web.C, w http.ResponseWriter, req *http.Request) {

}

func delTodo(c web.C, w http.ResponseWriter, req *http.Request) {

}
