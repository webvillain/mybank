package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/webvillain/vikashbank8/controllers"
	"github.com/webvillain/vikashbank8/db"
)

func main() {
	db.DbConn()
	r := mux.NewRouter()
	r.HandleFunc("/bank", controllers.ListAll).Methods("GET")
	r.HandleFunc("/bank/{id}", controllers.SingleUser).Methods("GET")
	r.HandleFunc("/bank/{name}/{email}", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/bank/{id}/{email}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/bank/{id}", controllers.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
