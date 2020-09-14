package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BSON = Binary Encoded Json. Includes additional ty pes e.g. int, long, date, floating point..
// bson.D = ordered document   bson.D{{"hello", "world"}, {"foo", "bar"}}
// bson.M = unordered document/map   bson.M{"hello": "world", "foo": "bar"}
// bson.A = array   bson.A{"Eric", "kyle"}
// bson.E = usually userd as an element inside bson.D
type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount" bson:"discount"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

var iphone10 = Product{
	ID:          primitive.NewObjectID(),
	Name:        "iphone10",
	Price:       900,
	Currency:    "USD",
	Quantity:    40,
	Vendor:      "apple",
	Accessories: []string{"charger", "headset", "slotopener"},
	SkuID:       "1234",
}

var trimmer = Product{
	// IDがないと000000000のObjectIDが入る
	ID:          primitive.NewObjectID(),
	Name:        "easy philips trimmer",
	Price:       120,
	Currency:    "USD",
	Quantity:    400,
	Vendor:      "Philips",
	Accessories: []string{"charger", "comb", "bladeset", "cleaning oil"},
	SkuID:       "2345",
}

var speaker = Product{
	ID:          primitive.NewObjectID(),
	Name:        "speakers",
	Price:       300,
	Currency:    "USD",
	Quantity:    25,
	Vendor:      "Bosch",
	Discount:    4,
	Accessories: []string{"cables", "remote"},
	SkuID:       "3466",
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		fmt.Println(err)
	}
	// なければ作成する
	db := client.Database("tronics")
	collection := db.Collection("products")
	// res, err := collection.InsertOne(context.Background(), trimmer)
	// res, err := collection.InsertOne(context.Background(), bson.D{
	// 	{"name", "eric"},
	// 	{"surname", "cartman"},
	// 	{"hobbies", bson.A{"videogame", "alexa", "kfc"}},
	// })
	// bson.Mを使ってデータを入れる方法
	// res, err := collection.InsertOne(context.Background(), bson.M{
	// 	"name":    "eric",
	// 	"surname": "cartman",
	// 	"hobbies": bson.A{"videogame", "alexa", "kfc"},
	// })

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}
