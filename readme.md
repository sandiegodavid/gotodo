Welcome to GoTodo.

# Project Setup
1. start MongoDB
2. install mgo package: `go get gopkg.in/mgo.v2`
3. install gorilla/mux package: `go get github.com/gorilla/mux`
4. checkout src repository: `cd $GOPATH/src;git clone https://github.com/sandiegodavid/gotodo.git;cd gotodo`
5. start server: `go install;$GOPATH/bin/gotodo 8080`

Afterwards, you can sending REST commands(e.g. curl, Postman, Advanced Rest Client) e.g.

* `curl http://localhost:8080/task/add -X POST -H "ContentType: application/json" -d '{"desc":"huge task", "due": "yesterday}'`
* `curl http://localhost:8080/task/1`
* `curl http://localhost:8080/task/list`
* `curl http://localhost:8080/task/1 -X PUT -H "ContentType: application/json" -d '{"desc":"big task", "due": "today"}'`
* `curl http://localhost:8080/task/delete/1 -X DELETE`

# Status

* All normal commands work
* Test missing
* Edge cases need better logging/response
* response header needs enhancement
