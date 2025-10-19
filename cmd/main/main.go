package main

import (
	"log"
	"net/http"

	"github.com/Angelo-Castellon/Book_Management/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStore(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
