package met

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
		fmt.Fprintln(w, "error")
	}
	json.Unmarshal(reqbody, &newTask) // asigno los valores  recibido al struct
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tasks)
}

//------------------------------------------*
func Indexrouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola")
}

//------------------------------------------*
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

//------------------------------------------*
func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["ID"])
	if err != nil {
		fmt.Fprintln(w, "error id")
		return
	}
	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(tasks)
		}
	}
}

//------------------------------------------*
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["ID"])
	if err != nil {
		fmt.Fprintln(w, "error id")
		return
	}
	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
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
	json.Unmarshal(reqbody, &updateTask)
	for i, t := range tasks {
		if t.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updateTask.ID = taskId
			tasks = append(tasks, updateTask)
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(tasks)
		}
	}

}
