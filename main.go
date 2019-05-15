package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

var dbName = "postgres"
var dbConnect = "host=localhost port=5432 user=rodney dbname=dlab password=bob sslmode=disable"

// InitialMigrations sets up the database objects
func InitialMigrations() {

	MigrateManuscript()

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/manuscripts", AllManuscripts).Methods("GET")
	//myRouter.HandleFunc("/manscript/new", NewUser).Methods("POST")
	//myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	//myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fmt.Println("Starting dlabcloud")
	InitialMigrations()
	handleRequests()
}
