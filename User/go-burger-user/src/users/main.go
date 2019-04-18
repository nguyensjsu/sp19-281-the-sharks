package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = os.Getenv("SERVER")
var mongodb_database = os.Getenv("DATABASE")
var mongodb_collection = os.Getenv("COLLECTION")
var mongo_admin_database = os.Getenv("ADMIN_DATABASE")
var mongo_username = os.Getenv("USERNAME")
var mongo_password = os.Getenv("PASSWORD")

func testUserAPI(w http.ResponseWriter, req *http.Request) {
	respBody, err := json.MarshalIndent("API working", "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	ResponseWithJSON(w, respBody, http.StatusOK)
	return
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func GetAllUser(w http.ResponseWriter, req *http.Request) {

	fmt.Println("in get all user")
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	var users []User
	err1 := c.Find(bson.M{}).All(&users)
	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed get all users: ", err1)
		return
	}

	respBody, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func GetUserById(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	var userData User
	_ = json.NewDecoder(req.Body).Decode(&userData)
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)

	var user User
	err1 := c.Find(bson.M{"id": params["id"]}).One(&user)
	if err1 != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed find book: ", err1)
		return
	}

	if user.Id == "" {
		ErrorWithJSON(w, "User not found", http.StatusNotFound)
		return
	}

	respBody, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	unqueId := uuid.Must(uuid.NewV4())
	user.Id = unqueId.String()
	session, err := mgo.Dial(mongodb_server)

	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()

	c := session.DB(mongodb_database).C(mongodb_collection)

	err = c.Insert(user)
	if err != nil {
		if mgo.IsDup(err) {
			ErrorWithJSON(w, "User with this ID already exists", http.StatusBadRequest)
			return
		}

		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Rgistration failed ", err)
		return
	}

	respBody, err := json.MarshalIndent(user, "", "  ")

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"id": params["id"]}
	err = c.Remove(query)

	if err != nil {
		switch err {
		default:
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed delete user: ", err)
			return
		case mgo.ErrNotFound:
			ErrorWithJSON(w, "User not found", http.StatusNotFound)
			return
		}
	}

	respBody, err := json.MarshalIndent("User deleted", "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func UserSignIn(w http.ResponseWriter, req *http.Request) {
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	query := bson.M{"email": person.Email,
		"password": person.Password}
	var user User

	err = c.Find(query).One(&user)
	if err == mgo.ErrNotFound {

		ErrorWithJSON(w, "Login Failed", http.StatusUnauthorized)
		return
	}
	userData := bson.M{
		"email":   user.Email,
		"name":    user.Name,
		"address": user.Address,
		"id":      user.Id}

	respBody, err := json.MarshalIndent(userData, "", "  ")
	ResponseWithJSON(w, respBody, http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/ping", testUserAPI).Methods("GET")
	router.HandleFunc("/users", GetAllUser).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users/signup", RegisterUser).Methods("POST")
	router.HandleFunc("/users/signin", UserSignIn).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
