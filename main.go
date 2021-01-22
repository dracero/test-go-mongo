//	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://brad123:brad123@cluster0.zf9fl.mongodb.net/udemy?retryWrites=true&w=majority"))
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Connection mongoDB with helper class
collection := helper.ConnectDB()
func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}
