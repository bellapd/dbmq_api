package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Delete struct {
	logger      *log.Logger
	MongoClient *mongo.Client
}

type DeleteUser struct {
	ID int `json:"id"`
}

func NewDelete(logger *log.Logger, mongoClient *mongo.Client) *Delete {
	return &Delete{logger, mongoClient}
}

func (delete *Delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	delete.logger.Println("Received request for 'DELETE'")
	w.WriteHeader(http.StatusOK) // 200 OK

	// read the json file sent by users
	var received DeleteUser
	json.NewDecoder(r.Body).Decode(&received)
	if received == (DeleteUser{}) {
		delete.logger.Println("Received empty data")
		return
	} else {
		delete.logger.Println("Received data: ", received)
	}

	// Delete student from Mongodb
	collection := delete.MongoClient.Database("test").Collection("students")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": received.ID})
	if err != nil {
		delete.logger.Println("Error deleting student: ", err)
	} else {
		delete.logger.Println("Successfully deleted student: ", received)
	}

	w.Write([]byte("Received request for 'DELETE'"))
}
