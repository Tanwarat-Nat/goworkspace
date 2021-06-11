package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func readSingleDoc(collection *mongo.Collection, year int) {
	fmt.Println(collection.Name())
	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{"year": 1894}
	singleResult := collection.FindOne(ctx, filter)
	cancelFindOne()

	raw, err := singleResult.DecodeBytes()
	if err != nil {
		log.Fatal("main: findOne: ", err)
	}
	fmt.Println(raw)

	result := Movie{}
	singleResult.Decode(&result)
	fmt.Println("---------------------")
	resultByte, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("main: unable to maeshal json: ", result)
	}
	fmt.Println(string(resultByte))
}

func insertDoc(collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//res, err := collection.InsertOne(ctx, Movie{ID: primitive.NewObjectID(), ImdbID: 1234, Title: "Avengers EndGame"})
	res, err := collection.InsertOne(ctx, bson.M{"_id": primitive.NewObjectID(), "imdbID": "1234", "Ttitle": "Avengers EndGame"})
	if err != nil {
		return nil, err
	}
	return res, nil

}

/*
ถ้าส่งไปแบบ Movie อันไนไม่เอา ค่าที่แสดงผลก็จะเป็น defult ของตัว field ต่างๆ
ควรระวัง!! ถ้าเราไม่ส่งไปเป็นแบบ Movie ถ้าเราส่งไปเป็นตาม official document
res, err := collection.InsertOne(ctx, bson.M{_id: primitive.NewObjectID(), imdbID: 1234, title: "Avengers EndGame"})
แบบนี้มันก็จะ flexible มาก มันจะแสดงผลเท่าที่เราใส่ไป อันไหนไม่ใส่ก็ไม่แสดงผลเลย

BSON (Binary JSON): เอา JSON เข้ามาครอบ สร้างอะไรอย่างนึงมาครอบ JSON เพื่อสามารถเข้าไปเช็คเข้าไปอ่าน
เข้าไปทำ operation ต่างๆได้รวดเร็วมากขึ้น
*/
