package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Dni       string             `json:"dni,omitempty" bson:"dni,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Adress    *Adress            `json:"adress,omitempty" bson:"adress,omitempty"`
}

type Adress struct {
	Street string `json:"street,omitempty"`
	Number string `json:"number,omitempty"`
}

//MongoDB
var client *mongo.Client

func main() {
	fmt.Println("Running...")

	//mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//routes
	router := mux.NewRouter()
	router.HandleFunc("/people", createPerson).Methods("POST")
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{dni}", getPerson).Methods("GET")
	router.HandleFunc("/people/{dni}", updatePerson).Methods("PUT")
	router.HandleFunc("/people/{dni}", deletePerson).Methods("DELETE")

	//start server
	http.ListenAndServe(":8080", router)
}

func getPeople(response http.ResponseWriter, request *http.Request) {
	//localhost:8080/people
	response.Header().Add("content-type", "application/json")

	var people []Person

	collection := client.Database("golang").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(people)
}

func getPerson(response http.ResponseWriter, request *http.Request) {
	//localhost:8080/people/{dni}

	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	dni, _ := params["dni"]

	var person Person

	collection := client.Database("golang").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, Person{Dni: dni}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(person)
}

func createPerson(response http.ResponseWriter, request *http.Request) {
	//POST
	//localhost:8080/people
	//{"dni": "33333333", "firstname":"Nahuel", "lastname":"Avalos", "adress":{"street":"Avenida Siempreviva","number":"742"}}

	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)

	collection := client.Database("golang").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(result)
}

func updatePerson(response http.ResponseWriter, request *http.Request) {
	//PUT
	//localhost:8080/people/{dni}
	//{"dni": "33333333", "firstname":"Nahuel Ivan", "lastname":"Avalos", "adress":{"street":"Avenida Siempreviva","number":"742"}}

	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	dni, _ := params["dni"]

	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	update := bson.M{"$set": person}

	collection := client.Database("golang").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, Person{Dni: dni}, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(result)
}

func deletePerson(response http.ResponseWriter, request *http.Request) {
	//DELETE
	//localhost:8080/people/{dni}

	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	dni, _ := params["dni"]

	collection := client.Database("golang").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, Person{Dni: dni})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(result)
}
