package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/webvillain/vikashbank8/model"
)

var Db = DbConn()

func DbConn() *sql.DB {
	db, err := sql.Open("sqlite3", "./bank.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (Id INTEGER PRIMARY KEY NOT NULL , Name TEXT , Email TEXT);")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
	//	defer db.Close()
	return db

}

func ListUsers() ([]*model.User, error) {
	var user *model.User
	var users []*model.User
	rows, err := Db.Query("SELECT * FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email)
	}
	users = append(users, user)
	return users, nil
}

func SingleUser(id int64) (*model.User, error) {
	var user *model.User
	row, err := Db.Query("SELECT * FROM users WHERE Id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email)
	}

	return user, nil
}

// create user works fine
func CreateNewUser(name string, email string) (*model.User, error) {
	var newuser *model.User
	stmt, err := Db.Prepare("INSERT INTO users (Name , Email)VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(name, email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())

	return newuser, nil
}

func UpdateUser(id int, email string) error {
	updateuser := `UPDATE users
	SET Email = ?
	WHERE Id = ?;`

	stmt, err := Db.Prepare(updateuser)
	if err != nil {
		log.Fatal(err)
	}
	res, _ := stmt.Exec(id, email)
	fmt.Println("Rows Affected : ", res)
	return nil
}

func DeleteUser(Id int) error {
	deleteuser := `DELETE FROM users
	WHERE Id = ?;`
	stmt, err := Db.Prepare(deleteuser)
	if err != nil {
		log.Fatal(err)
	}
	res, _ := stmt.Exec(Id)
	fmt.Println("Rows Affected :", res)
	return nil
}
