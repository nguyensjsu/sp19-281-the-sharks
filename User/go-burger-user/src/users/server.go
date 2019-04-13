package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var mongodb_server = os.Getenv("MONGO_SERVER")
var mongodb_database = os.Getenv("MONGO_DATABASE")
var mongodb_collection = os.Getenv("MONGO_COLLECTION")
var mongo_admin_database = os.Getenv("MONGO_ADMIN_DATABASE")
var mongo_username = os.Getenv("MONGO_USERNAME")
var mongo_password = os.Getenv("MONGO_PASS")

func init() {
	fmt.Println("Mongodb ENV: ", mongodb_server)
	fmt.Println("Mongodb DB ENV: ", mongodb_database)
}

func MenuServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router))
	return n
}

func initRoutes(router *mux.Router, formatter *render.Render) {

	router.HandleFunc("/users/test/ping", checkPing(formatter)).Methods("GET")
}

func checkPing(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Burger Users API Server Working on machine: "
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}
