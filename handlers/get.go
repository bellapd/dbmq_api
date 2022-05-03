package handlers

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"gopkg.in/mgo.v2/bson"
// )

// type Get struct {
// 	logger  *log.Logger
// 	mongodb *mongo.Client
// }

// func NewGet(logger *log.Logger, mongoClient *mongo.Client) *Get {
// 	return &Get{logger, mongoClient}
// }

// func (get *Get) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	get.logger.Println("Received request for 'POST'")
// 	w.WriteHeader(http.StatusOK) // 200 OK
// 	w.Write([]byte("Received request for 'POST'"))

// 	// print out all the users in the MongoDB
// 	collection := get.mongodb.Database("test").Collection("students")
// 	var result Student
// 	cursor, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		get.logger.Println("Failed to query MongoDB:", err)
// 		return
// 	}
// 	defer cursor.Close(context.TODO())
// 	for cursor.Next(context.TODO()) {
// 		err := cursor.Decode(&result)
// 		if err != nil {
// 			get.logger.Println("Failed to decode cursor:", err)
// 			return
// 		}
// 		get.logger.Println("Student:", result)
// 	}

// 	get.logger.Println("The end of 'GET' request")
// 	w.Write([]byte("The end of 'GET' request"))
// }
