package main

import (
	"apirestGorm/handlers"
	"apirestGorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.MigrarUser()

	//rutas
	mux := mux.NewRouter()

	//endPoind
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")                 //OBTENER TODOS
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")       //OBTENER 1
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")              //GRABAR 1
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")    //EDITAR 1
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE") //ELIMINAR 1
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
