package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/marcelozilio/golang-rest-api/tree/master/config"
	. "github.com/marcelozilio/golang-rest-api/tree/master/config/dao"
	personrouter "github.com/marcelozilio/golang-rest-api/tree/master/router"
)

var dao = PersonDAO{}
var config = Config{}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/persons", personrouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/persons/{id}", personrouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/persons", personrouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/persons/{id}", personrouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/persons/{id}", personrouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}