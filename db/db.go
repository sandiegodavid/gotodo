package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gotodo/model"
	"log"
)

var session *mgo.Session
var db *mgo.Database
var todos *mgo.Collection
var counters *mgo.Collection
var err error
var _taskid = "taskid"

/*
Init init function
*/
func Init() {
	lsession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session = lsession
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("gotodo")
	todos = db.C("todos")
	initTaskIDCounter()
}

/*
Add add function
*/
func Add(t *model.TodoJSON) {
	t.ID = getNextTaskID()
	err = todos.Insert(t)
	if err != nil {
		panic(err)
	}
}

/*
Update update function
*/
func Update(t *model.TodoJSON) {
	err = todos.Update(bson.M{}, t)
	if err != nil {
		panic(err)
	}
}

func initTaskIDCounter() {
	counters = db.C("counters")
	idDocCount, err := counters.Find(bson.M{"name": _taskid}).Count()
	if err == nil && idDocCount == 0 {
		counters.Insert(bson.M{"name": _taskid, "seq": 0})
	}
}

func getNextTaskID() int64 {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}
	idDoc := model.TaskIDDoc{}
	info, err := counters.Find(bson.M{"name": _taskid}).Apply(change, &idDoc)
	if err != nil {
		log.Println(info)
		panic(err)
	}
	return idDoc.Seq
}

/*
Find find function
*/
func Find(id int64) *model.TodoJSON {
	todo := model.TodoJSON{}
	err = todos.Find(bson.M{"id": id}).One(&todo)
	return &todo
}

/*
List list function
*/
func List(todolist *[]model.TodoJSON) {
	err = todos.Find(bson.M{}).All(todolist)
}

/*
Delete delete function
*/
func Delete(id int64) {
	log.Println(id)
	err = todos.Remove(bson.M{"id": id})
}

/*
Close close function
*/
func Close() {
	session.Close()
}
