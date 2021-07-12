package sqlg

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

//----------------------------------------------------------------
func Insert(name, content string) int {

	conndb := Connect()
	insert, err := conndb.Prepare(" INSERT INTO tasks (name, content) VALUES (\"" + name + "\", \"" + content + "\")")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec()

	id, err := strconv.Atoi(Selecte(name))
	if err != nil {
		fmt.Println(err.Error())
	}
	return id
}

//------------------------------
func Connect() (conn *sql.DB) {
	driver := "mysql"
	cadena := os.Getenv("cadena")
	conn, err := sql.Open(driver, cadena)
	if err != nil {
		panic(err.Error())
	}
	return conn

}

//------------------------------
func Update(id, name, content string) {
	conndb := Connect()
	insert, err := conndb.Prepare("UPDATE tasks SET content =\"" + content + "\", name =\"" + name + "\" WHERE id =" + id)
	if err != nil {
		panic(err.Error())
	}
	insert.Exec()
}

//------------------------------
func SelectAll(id string) []Task {
	where := ""
	if id != "" {
		where = "where id=" + id
	}
	registry, err := Connect().Query("SELECT * FROM tasks " + where)

	var tasks []Task
	var task Task
	if err != nil {
		fmt.Println(err.Error())
	}
	for registry.Next() {
		var ids int
		var name, content string
		err := registry.Scan(&ids, &name, &content)
		if err != nil {
			fmt.Println(err.Error())
		}
		task.ID = ids
		task.Content = content
		task.Name = name
		tasks = append(tasks, task)
	}
	return tasks
}

//------------------------------
func Selecte(name string) string {
	registry, err := Connect().Query("SELECT id FROM tasks where name=\"" + name + "\"")
	var id string
	if err != nil {
		fmt.Println(err.Error())
	}
	for registry.Next() {
		err := registry.Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

	return id
}

//------------------------------
func Delete(id string) []Task {
	query, err := Connect().Prepare("DELETE FROM tasks WHERE id=" + id)
	if err != nil {
		panic(err.Error())
	}
	query.Exec()
	alltasksS := SelectAll("")
	return alltasksS
}
