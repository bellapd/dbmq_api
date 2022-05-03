package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	logger      *log.Logger
	MongoClient *mongo.Client
}

type YAML struct {
	APIVersion string `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata struct {
		Name string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
}

func NewPost(logger *log.Logger, mongoClient *mongo.Client) *Post {
	return &Post{logger, mongoClient}
}

func (post *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	post.logger.Println("Received request for 'POST'")
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Received request for 'POST'"))

	data, err := io.ReadAll(r.Body)
	if err != nil {
		post.logger.Println("Failed to read body:", err)
	}
	// read file sent by users
	fmt.Println(string(data))
	// fmt.Println(data)

	w.Write([]byte("Received request for 'POST'"))
}
