package main

import (
	"api-go/met"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", met.Indexrouter)
	router.HandleFunc("/tasks/{ID}", met.GetTask).Methods("GET")
	router.HandleFunc("/tasks", met.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{ID}", met.DeleteTasks).Methods("DELETE")
	router.HandleFunc("/tasks/{ID}", met.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks", met.CreateTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":803", router))
}
