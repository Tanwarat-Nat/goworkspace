package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dbGoUser:dbGoUser@gocourse.kdrqc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal("main: cannot create client.: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("main: cannot connect to database.")
	}

	collection := client.Database("uncleBobDvd").Collection("movie_initial")

	//Read data years.
	//readSingleDoc(collection, 1894)

	// res, err := insertDoc(collection)
	// if err != nil {
	// 	log.Fatal("main: insert failed. ", err)
	// }
	// fmt.Println("Insert success ", res.InsertedID)

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//60b4ed54caa890ab7df1c0a1
	_, err = primitive.ObjectIDFromHex("60b4ed54caa890ab7df1c0a1")
	if err != nil {
		log.Fatal("lmain: cannot create object id")
	}

	err = collection.Drop(ctx)
	if err != nil {
		log.Fatal("main: cannot drop collection")
	}
	fmt.Println("collection dropped")

	//Update
	// update := bson.D{{Key: "$rename", Value: bson.D{{Key: "Ttitle", Value: "title"}}}}

	// uresult, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	// if err != nil {
	// 	log.Fatal("main: cannot update document")
	// }
	// fmt.Println("Update successfully: ", uresult.UpsertedID)
}

// format ให้อ่านง่าย: มำการ marshel structure ตรงนี้ไปเป็น string jason
type Movie struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	ImdbID      int                `json:"imdbID"`
	Title       string             `json:"title"`
	Year        int                `json:"year"`
	Rating      string             `json:"rating"`
	Runtime     string             `json:"runtime"`
	Genre       string             `json:"genre"`
	Released    string             `json:"released"`
	Director    string             `json:"director"`
	Writer      string             `json:"writer"`
	Cast        string             `json:"cast"`
	Metacritic  string             `json:"metacritic"`
	ImdbRating  float64            `json:"imdbRating"`
	ImdbVotes   int                `json:"imdbVotes"`
	Poster      string             `json:"poster"`
	Plot        string             `json:"plot"`
	Fullplot    string             `json:"fullplot"`
	Language    string             `json:"language"`
	Country     string             `json:"country"`
	Awards      string             `json:"awards"`
	Lastupdated string             `json:"lastupdated"`
	Type        string             `json:"type"`
}

/*encode : แปลงค่าปกติใน Go ให้ไปเป็น bson type
decode : decode จาก internal type ของ mongodb เองกลับมาเป็น gplang
*/
