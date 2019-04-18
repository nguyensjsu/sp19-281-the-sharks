package main

import (
	"encoding/json"
	"fmt"
	"log"
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
var mongo_user = os.Getenv("MONGO_USERNAME")
var mongo_pass = os.Getenv("MONGO_PASS")

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API ROUTES
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/order/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order", orderHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order/{orderId}", orderHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders/{userId}", orderInfo(formatter)).Methods("GET")
	mx.HandleFunc("/order", newOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/order", deleteOrderHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/orders", updateOrderHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/orders/updateorderstatus/{orderid}", updateOrderStatusHandler(formatter)).Methods("PUT")

}

// HELPER FUNCTIONS
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API PING HANDLER
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"BURGER ORDER API Version 1.0 is ALIVE!"})
	}
}

// API BURGER ORDER HANDLER
func orderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err := session.DB("admin").Login(mongo_user, mongo_pass)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var uuid string = params["orderId"]
		if uuid == "" {
			var orders_array []BurgerOrder
			err = c.Find(bson.M{}).All(&orders_array)
			fmt.Println("Burger Orders:", orders_array)
			formatter.JSON(w, http.StatusOK, orders_array)
		} else {
			fmt.Println("orderID: ", uuid)
			var result BurgerOrder
			err = c.Find(bson.M{"orderId": uuid}).One(&result)
			if err != nil {
				formatter.JSON(w, http.StatusNotFound, "Order Not Found")
				return
			}
			_ = json.NewDecoder(req.Body).Decode(&result)
			fmt.Println("Burger Order: ", result)
			formatter.JSON(w, http.StatusOK, result)
		}
	}
}

// API GET ORDER INFO BY USERID
func orderInfo(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		err := session.DB("admin").Login(mongo_user, mongo_pass)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		params := mux.Vars(req)
		var uuid string = params["userId"]
		fmt.Println("userId: ", uuid)
		var result []BurgerOrder
		err = c.Find(bson.M{"userId": uuid}).All(&result)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if len(result) == 0 {
			formatter.JSON(w, http.StatusNotFound, "User Doesn't have order")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Burger Order: ", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API PLACE NEW BURGER ORDER
func newOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var newOrder BurgerOrder
		if err := json.NewDecoder(req.Body).Decode(&newOrder); err != nil {
			formatter.JSON(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		fmt.Println(newOrder)

		newOrder.OrderId = bson.NewObjectId()
		newOrder.OrderStatus = "Order has been placed!!"
		newOrder.PaymentStatus = "Pending"

		if err := c.Insert(&newOrder); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		var result BurgerOrder
		// if err = c.FindId(bson.ObjectIdHex(order.OrderID)).One(&result); err != nil {
		// 	formatter.JSON(w, http.StatusInternalServerError, err.Error())
		// 	return
		// }

		if err = c.FindId(newOrder.OrderId).One(&result); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		formatter.JSON(w, http.StatusOK, result)
	}
}

// API DELETE BURGER
func deleteOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		orderid := vars["orderid"]

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		fmt.Println(orderid)

		if err := c.Remove(bson.M{"OrderID": orderid}); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		formatter.JSON(w, http.StatusOK, "Order has been deleted successfully!!")
	}
}

//UPDATE BURGER ORDER
func updateOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var newOrder BurgerOrder
		if err := json.NewDecoder(req.Body).Decode(&newOrder); err != nil {
			formatter.JSON(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := c.UpdateId(newOrder.OrderId, &newOrder); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		var updatedOrder BurgerOrder

		if err = c.FindId(newOrder.OrderId).One(&updatedOrder); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		c.Find(bson.M{"newOrder.OrderID": newOrder.OrderId}).One(&updatedOrder)
		formatter.JSON(w, http.StatusOK, updatedOrder)

	}
}

//UPDATE ORDER STATUS

func updateOrderStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		orderid := vars["orderid"]

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var result BurgerOrder
		if err = c.FindId(bson.ObjectIdHex(orderid)).One(&result); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		result.OrderStatus = "Order has been dispatched!"
		result.PaymentStatus = "Payment received"

		if err := c.UpdateId(result.OrderId, &result); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		var updatedOrder BurgerOrder

		if err = c.FindId(result.OrderId).One(&updatedOrder); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		formatter.JSON(w, http.StatusOK, updatedOrder)

	}
}
