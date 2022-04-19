// connect to mongodb
package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
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

func deleteStudent(collection *mongo.Collection, name string) {
	// delete one student
	filter := bson.M{"name": name}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func inputStudent(collection *mongo.Collection) {
	var input_student Student

	fmt.Println("Enter name: ")
	fmt.Scanln(&input_student.Name)
	fmt.Println("Enter age: ")
	fmt.Scanln(&input_student.Age)
	fmt.Println("Enter class: ")
	fmt.Scanln(&input_student.Class)

	filter := bson.M{"name": input_student.Name, "age": input_student.Age, "class": input_student.Class}
	var result Student
	err := collection.FindOne(context.TODO(), filter).Decode(&result) // ini buat find data
	if err != mongo.ErrNoDocuments {
		fmt.Println("Found a duplicate: ", result)
	} else {
		_, err := collection.InsertOne(context.TODO(), input_student)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	getAllStudents(collection)
}

func getAllStudents(collection *mongo.Collection) {
	var result Students
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//print all students
	for cursor.Next(context.TODO()) {
		var student Student
		err := cursor.Decode(&student)
		if err != nil {
			fmt.Println(err)
			return
		}
		result = append(result, student)
	}
	fmt.Println("Found a result: ", result)
}

func main() {
	mongoClient := connectMongo(mongoURL)
	defer mongoClient.Disconnect(context.TODO())
	fmt.Println("Connected to MongoDB")

	collection := mongoClient.Database("test").Collection("students")

	var student = Students{
		{
			Name:  "John",
			Age:   20,
			Class: "A",
		},
		{
			Name:  "Jane",
			Age:   21,
			Class: "B",
		},
		{
			Name:  "Jack",
			Age:   22,
			Class: "C",
		},
	}

	for _, student := range student {
		filter := bson.M{"name": student.Name}
		var result Student
		err := collection.FindOne(context.TODO(), filter).Decode(&result) // ini buat find data
		if err != mongo.ErrNoDocuments {
			fmt.Println("Found a result: ", result)
		} else {
			_, err := collection.InsertOne(context.TODO(), student) // ini buat insert data
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	for {
		inputStudent(collection)
		// stop input if control + c is pressed
		if err := mongoClient.Ping(context.TODO(), nil); err != nil {
			fmt.Println(err)
			return
		}
	}
}
