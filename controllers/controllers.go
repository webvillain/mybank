package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/webvillain/vikashbank8/db"
)

func ListAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := db.ListUsers()
	if err != nil {
		log.Fatal(err)
	}
	data := json.NewEncoder(w).Encode(users)
	fmt.Fprintln(w, data)

}
func SingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	opts := mux.Vars(r)
	id := opts["id"]
	newid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	user, err := db.SingleUser(newid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, user)

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// /var user1 *model.User
	opts := mux.Vars(r)
	name := opts["name"]
	email := opts["email"]
	_, err := db.CreateNewUser(name, email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "New User Is Created Successfully.")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	opts := mux.Vars(r)
	id := opts["id"]
	email := opts["email"]
	newid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	err = db.UpdateUser(int(newid), email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "User Is Updated .")

	w.Header().Set("Content-Type", "application/json")

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	opts := mux.Vars(r)
	id := opts["id"]
	newid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	err = db.DeleteUser(int(newid))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "User Id Deleted .")
}
