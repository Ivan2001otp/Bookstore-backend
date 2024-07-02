package main

import (
	router "bookstore_backend/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	router.RegisteredRouters(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
