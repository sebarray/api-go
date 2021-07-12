package met

import (
	"api-go/sqlg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}
type AllTask []Task

var tasks = AllTask{
	{
		ID:      1,
		Name:    "TAREA1",
		Content: "prueba de tarea",
	},
}

//------------------------------------------*
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	reqbody, err := ioutil.ReadAll(r.Body) //recibo datos que envia el cliente
	if err != nil {
		fmt.Fprintln(w, "error al recibir datos del cliente")
	}

	json.Unmarshal(reqbody, &newTask)                       // asigno los valores  recibido al struct
	newTask.ID = sqlg.Insert(newTask.Name, newTask.Content) //********************************CREAR TAREAS
	//
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

//------------------------------------------*
func Indexrouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola")
	sqlg.Selecte("js")
}

//------------------------------------------*
func GetTasks(w http.ResponseWriter, r *http.Request) {
	alltasksS := sqlg.SelectAll("") //*******************************TRAE TAREAS
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(alltasksS)
}

//------------------------------------------*
func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tasks := sqlg.Delete(vars["ID"])
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

//------------------------------------------*
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	task := sqlg.SelectAll(vars["ID"]) //********************************TRAE UNA TAREA
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(task)

}

//----------------------------------------------------*

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["ID"])
	var updateTask Task

	if err != nil {
		fmt.Fprintln(w, "error de id")
	}
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "error de dato")
	}
	updateTask.ID = taskId
	json.Unmarshal(reqbody, &updateTask)

	sqlg.Update(vars["ID"], updateTask.Name, updateTask.Content) //******************************ACTUALIZA UNA TAREA

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(updateTask)

}

//----------------------------------------------------------------
