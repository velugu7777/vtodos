package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

type Todo struct {
	id     int
	name   string
	IsDone bool
}

var mytodos []Todo

func init() {
	t1 := Todo{
		id:     1,
		name:   "Reading books",
		IsDone: false,
	}
	t2 := Todo{
		id:     2,
		name:   "playing cricket",
		IsDone: false,
	}
	mytodos = append(mytodos, t1, t2)
}
func main() {
	goji.Get("/todo", todos)
	goji.Post("/todo/newTodo", newTodo)
	goji.Get("/todo:id", getTodo)
	goji.Put("/todo:id", putTodo)
	goji.Delete("/todo:id", delTodo)
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
