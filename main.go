package main

import (
	"api-go/met"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hola heroku")
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	PORT := os.Getenv("PORT")
	router.HandleFunc("/", met.Indexrouter).Methods("GET")
	router.HandleFunc("/tasks/{ID}", met.GetTask).Methods("GET")
	router.HandleFunc("/tasks", met.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{ID}", met.DeleteTasks).Methods("DELETE")
	router.HandleFunc("/tasks/{ID}", met.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks", met.CreateTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origin)(router)))
}
