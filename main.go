package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  int32              `bson:"age"`
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Karma",
		Age:  21,
	}

	fmt.Print(user)

	collection := client.Database("user").Collection("userData")

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}

	fmt.Print(result.InsertedID)

	//filter := bson.D{{Key: "name", Value: user.Name}}
	//userResult := User{}
	//errFind := collection.FindOne(context.Background(), filter).Decode(&userResult)
	//if errFind != nil {
	//	panic(errFind)
	//}
	//
	//fmt.Print(userResult)

	//cur, err := collection.Find(context.Background(), bson.D{})
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer cur.Close(context.Background())
	//for cur.Next(context.Background()) {
	//	userResult := User{}
	//
	//	err := cur.Decode(&userResult)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(userResult)
	//}
}
