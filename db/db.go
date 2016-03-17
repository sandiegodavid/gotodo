package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var session *mgo.Session
var db *mgo.Database
var todos *mgo.Collection
var counters *mgo.Collection
var err error
var _taskid = "taskid"

type TodoJson struct {
	Id        int64  `json:"id"`
	Desc      string `json:"desc"`
	Due       string `json:"due"`
	Completed bool   `json:"completed"`
}

type TaskIdDoc struct {
	Name string `json:"name"`
	Seq  int64  `json:"seq"`
}

func Init() {
	lsession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session = lsession
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("gotodo")
	todos = db.C("todos")
	initTaskIdCounter()
	// todo := TodoJson{23, "go get it", "a 3 hours", false}
	// Add(&todo)
	// find(23)
}

func Add(t *TodoJson) {
	t.Id = getNextTaskId()
	err = todos.Insert(t)
	if err != nil {
		panic(err)
	}
}

func Update(t *TodoJson) {
	err = todos.Update(bson.M{}, t)
	if err != nil {
		panic(err)
	}
}

func initTaskIdCounter() {
	counters = db.C("counters")
	idDocCount, err := counters.Find(bson.M{"name": _taskid}).Count()
	if err == nil && idDocCount == 0 {
		counters.Insert(bson.M{"name": _taskid, "seq": 0})
	}
}

func getNextTaskId() int64 {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}
	idDoc := TaskIdDoc{}
	info, err := counters.Find(bson.M{"name": _taskid}).Apply(change, &idDoc)
	if err != nil {
		log.Println(info)
		panic(err)
	}
	return idDoc.Seq
}

func Find(id int64) *TodoJson {
	todo := TodoJson{}
	err = todos.Find(bson.M{"id": id}).One(&todo)
	log.Println(todo)
	return &todo
}

func List(todolist []TodoJson) {
}

func Delete(id int64) {
}

func Close() {
	session.Close()
}
