// connect to mongodb
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dbmq/handlers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Class string
}

type Students []Student

const mongoURL string = "mongodb://localhost:27017"

func connectMongo(url string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return client
}

// func getAllStudents(collection *mongo.Collection) {
// 	var result Students
// 	cursor, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//print all students
// 	for cursor.Next(context.TODO()) {
// 		var student Student
// 		err := cursor.Decode(&student)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		result = append(result, student)
// 	}
// 	fmt.Println("Found a result: ", result)
// }

func main() {
	mongoClient := connectMongo(mongoURL)
	defer mongoClient.Disconnect(context.TODO())
	fmt.Println("Connected to MongoDB")

	logger := log.New(os.Stdout, "student_api", log.LstdFlags) // tiap kali ada event di API dia simpan di log

	postHandler := handlers.NewPost(logger, mongoClient)
	// getHandler := handlers.NewGet(logger, mongoClient)
	// deleteHandler := handlers.NewDelete(logger, mongoClient)

	serveMux := http.NewServeMux()
	serveMux.Handle("/post", postHandler)
	// serveMux.Handle("/get", getHandler)
	// serveMux.Handle("/delete", deleteHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}() // run in background

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	signal := <-signalChannel
	logger.Println("Received termination, gracefully shutdown", signal)

	tc, _ := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	server.Shutdown(tc)
}
