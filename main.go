package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// HttpResponse struct for example
type HttpResponse struct {
	ID      string   `json:"_id" bson:"_id"`
	DevTeam []string `json:"devTeam" bson:"devTeam"`
}

// Constructor for HttpResponse, where the default value for DevTeam is an empty slice
func NewHttpResponse() *HttpResponse {
	r := new(HttpResponse)
	r.DevTeam = []string{}
	return r
}

func main() {
	// Configure and connect to MongoDB instance
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping connection to check liveness and connectivity
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Defer client disconnection
	defer client.Disconnect(context.Background())

	// Retrieve a single document from the collection
	nowRowsFound := client.Database("databaseNameTest").Collection("collectionNameTest")
	filter := bson.D{{"_id", "663c385af3dad2d0510a7216"}}

	// First approach: setting the default value for DevTeam as an empty slice
	firstSolution := new(HttpResponse)
	firstSolution.DevTeam = []string{}
	err = nowRowsFound.FindOne(context.TODO(), filter).Decode(&firstSolution)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No documents found")
			return
		} else {
			log.Fatal(err)
		}
	}
	jsonFirstSolution, _ := json.Marshal(firstSolution)
	fmt.Println(string(jsonFirstSolution))
	// {"_id":"663c385af3dad2d0510a7216","devTeam":[]} if no document is found
	// {"_id":"663c385af3dad2d0510a7216","devTeam":["devTeam1"]} if document is found

	// Second approach: calling Constructor, where the default value for DevTeam is an empty slice
	secondApproach := NewHttpResponse()
	err = nowRowsFound.FindOne(context.TODO(), filter).Decode(&secondApproach)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No documents found")
			return
		} else {
			log.Fatal(err)
		}
	}
	jsonSecondSolution, _ := json.Marshal(secondApproach)
	fmt.Println(string(jsonSecondSolution))
	// {"_id":"663c385af3dad2d0510a7216","devTeam":[]} if no document is found
	// {"_id":"663c385af3dad2d0510a7216","devTeam":["devTeam1"]} if document is found

	// Third approach: validating if the DevTeam field is nil, then set it as an empty slice
	thirdApproach := NewHttpResponse()
	err = nowRowsFound.FindOne(context.TODO(), filter).Decode(&thirdApproach)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No documents found")
			return
		} else {
			log.Fatal(err)
		}
	}
	if thirdApproach.DevTeam == nil {
		thirdApproach.DevTeam = []string{}
	}
	jsonThirdSolution, _ := json.Marshal(thirdApproach)
	fmt.Println(string(jsonThirdSolution))
	// {"_id":"663c385af3dad2d0510a7216","devTeam":[]} if no document is found
	// {"_id":"663c385af3dad2d0510a7216","devTeam":["devTeam1"]} if document is found

}
