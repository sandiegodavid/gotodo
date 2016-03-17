package main

import (
	"encoding/json"
	"fmt"
	"gotodo/db"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	//db.Init()
	http.HandleFunc("/task/", taskHandler)
	log.Println(os.Args[1])
	http.ListenAndServe(os.Args[1], nil)
	//db.Close()
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	urlEdge := r.URL.Path[len("/task/"):]
	log.Println("r.URL.Path: " + r.URL.Path + ", urlEdge: " + urlEdge + "\n")
	switch r.Method {
	case "GET":
		handleGet(w, urlEdge)
	case "POST":
		handlePost(w, getToDo(r))
	case "PUT":
		handlePut(w, urlEdge, getToDo(r))
	case "DELETE":
		handleDelete(w, urlEdge[len("/delete/"):])
	default:
		// Give an error message.
	}
}

func getToDo(r *http.Request) *db.TodoJson {
	decoder := json.NewDecoder(r.Body)
	var t db.TodoJson
	err := decoder.Decode(&t)
	if err != nil {
		panic("bad json in create")
	}
	return &t
}

func handlePost(w http.ResponseWriter, t *db.TodoJson) {
	//	db.Add(t)
	fmt.Fprintf(w, strconv.FormatUint(uint64(t.Id), 10))
}

func handleGet(w http.ResponseWriter, urlEdge string) {
	if strings.Compare(urlEdge, "list") == 0 {
		var todolist []db.TodoJson
		db.List(todolist)
		respSeg, err := json.Marshal(todolist)
		if err != nil {
			panic("bad json in get list")
		}
		fmt.Fprintf(w, string(respSeg))
		return
	}
	i, err := strconv.ParseInt(urlEdge, 10, 64)
	if err != nil {
		panic("bad json in get")
	}
	todo := db.Find(i)
	respSeg, err := json.Marshal(todo)
	fmt.Fprintf(w, string(respSeg))
}

func handlePut(w http.ResponseWriter, urlEdge string, t *db.TodoJson) {
	i, err := strconv.ParseInt(urlEdge, 10, 64)
	if err != nil {
		panic("bad urlEdge in update: " + urlEdge)
	}
	t.Id = i
	db.Update(t)
	fmt.Fprintf(w, strconv.FormatUint(uint64(t.Id), 10))
}

func handleDelete(w http.ResponseWriter, urlEdge string) {
	i, err := strconv.ParseInt(urlEdge, 10, 64)
	if err != nil {
		panic("bad urlEdge in delete: " + urlEdge)
	}
	db.Delete(i)
	fmt.Fprintf(w, "ok")
}
