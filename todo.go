package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gotodo/db"
	"gotodo/model"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	db.Init()
	r := mux.NewRouter()
	r.HandleFunc("/{id:[0-9]+}", getOneHandler).Methods("GET")
	r.HandleFunc("/list", getListHandler).Methods("GET")
	r.HandleFunc("/task/delete/{id:[0-9]+}", deleteHandler).Methods("DELETE")
	r.HandleFunc("/task/{id:[0-9]+}", putHandler).Methods("PUT")
	r.HandleFunc("/task/add", postHandler).Methods("POST")
	http.Handle("/task", r)
	http.ListenAndServe(":"+os.Args[1], nil)
	db.Close()
}

func getToDo(r *http.Request) *model.TodoJSON {
	decoder := json.NewDecoder(r.Body)
	var t model.TodoJSON
	err := decoder.Decode(&t)
	if err != nil {
		panic("bad json in create")
	}
	log.Println(t)
	return &t
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	t := getToDo(r)
	db.Add(t)
	fmt.Fprintf(w, strconv.FormatUint(uint64(t.ID), 10))
}

func getOneHandler(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		panic("bad taskid in getOneHandler: " + mux.Vars(r)["id"])
	}
	todo := db.Find(i)
	respSeg, err := json.Marshal(todo)
	fmt.Fprintf(w, string(respSeg))
}

func getListHandler(w http.ResponseWriter, r *http.Request) {
	var todolist []model.TodoJSON
	db.List(&todolist)
	respSeg, err := json.Marshal(todolist)
	if err != nil {
		panic("bad json in get list")
	}
	fmt.Fprintf(w, string(respSeg))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	t := getToDo(r)
	i, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		panic("bad taskid in putHandler: " + mux.Vars(r)["id"])
	}
	t.ID = i
	db.Update(t)
	fmt.Fprintf(w, strconv.FormatUint(uint64(t.ID), 10))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		panic("bad taskid in deleteHandler: " + mux.Vars(r)["id"])
	}
	db.Delete(i)
	fmt.Fprintf(w, "ok")
}
