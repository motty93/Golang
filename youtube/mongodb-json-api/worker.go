package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatFactWorker struct {
	client *mongo.Client
}

func NewCatFactWorker(c *mongo.Client) *CatFactWorker {
	return &CatFactWorker{
		client: c,
	}
}

func (cfw *CatFactWorker) start() error {
	col := cfw.client.Database("catfacts").Collection("facts")
	fmt.Println(col)
	ticker := time.NewTicker(2 * time.Second)

	for {
		res, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}

		var catFact bson.M // map[string]any // map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&catFact); err != nil {
			return err
		}

		_, err = col.InsertOne(context.TODO(), catFact)
		if err != nil {
			return err
		}

		<-ticker.C
	}
}
