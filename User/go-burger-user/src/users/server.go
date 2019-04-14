package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	fmt.Println("MongoDB collection : ", mongodb_collection)
}

func MenuServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(router *mux.Router, formatter *render.Render) {

	router.HandleFunc("/users/ping", testUserAPI(formatter)).Methods("GET")
	router.HandleFunc("/users", GetAllUser).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
}

func testUserAPI(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Burger Users API working on go server"})
	}
}

func GetAllUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := struct{ Message string }{"Can't connect to Database."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
	err = session.DB(mongo_admin_database).Login(mongo_username, mongo_password)
	if err != nil {
		message := struct{ Message string }{"Can't login into database"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{}
	var result []bson.M
	err = c.Find(query).All(&result)
	if err != nil {
		message := struct{ Message string }{"No users were found!!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(message)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetUserById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var m User
	_ = json.NewDecoder(req.Body).Decode(&m)
	fmt.Println("Get data of user: ", params["id"])
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		message := struct{ Message string }{"Can't connect to database"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
	err = session.DB(mongo_admin_database).Login(mongo_username, mongo_password)
	if err != nil {
		message := struct{ Message string }{"Can't login into database"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"id": params["id"]}
	var result bson.M
	err = c.Find(query).One(&result)
	if err != nil && err != mgo.ErrNotFound {
		message := struct{ Message string }{"Exception while querying database"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(message)
		return
	} else if err == mgo.ErrNotFound {
		message := struct{ Message string }{"User not found for ID"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(message)
		return
	}
	json.NewEncoder(w).Encode(result)
}
