package data

/*
TodoJSON JSON for todo items
*/
type TodoJSON struct {
	ID        int64  `json:"id"`
	Desc      string `json:"desc"`
	Due       string `json:"due"`
	Completed bool   `json:"completed"`
}

/*
TaskIDDoc atomic counter for task id
*/
type TaskIDDoc struct {
	Name string `json:"name"`
	Seq  int64  `json:"seq"`
}
