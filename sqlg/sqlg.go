package sqlg

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Insert(name, content string) {

	conndb := Connect()
	insert, err := conndb.Prepare(" INSERT INTO tasks (name, content) VALUES (\"" + name + "\", \"" + content + "\")")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec()
}

func Connect() (conn *sql.DB) {
	driver := "mysql"
	user := "root"
	passwd := " "
	db := "taskdb"
	host := "@tcp(127.0.0.1)/"
	conn, err := sql.Open(driver, user+":"+passwd+host+db)
	if err != nil {
		panic(err.Error())
	}
	return conn

}

func Update(id, name, content string) {
	conndb := Connect()
	insert, err := conndb.Prepare("UPDATE tasks SET content =\"" + content + "\", name =\"" + name + "\" WHERE id =" + id)
	if err != nil {
		panic(err.Error())
	}
	insert.Exec()

}

func Selecte() {
	registry, err := Connect().Query("SELECT * FROM clientes")

	if err != nil {
		fmt.Println(err.Error())
	}
	for registry.Next() {
		var nombre, correo string
		err := registry.Scan(&nombre, &correo)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("el nombre es " + nombre)
		fmt.Println("el correo es " + correo)
	}

}
