package main

import (
	"context"
	"testing"

	"github.com/tj/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCollection struct {
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	c.InsertedID = "asdf"

	return c, nil
}

func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	return c, nil
}

func (m *mockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	c := &mongo.Cursor{}
	c.Current = bson.Raw([]byte(`
	[
		{
			"first_name": "krunal",
			"last_name": "shimpi",
		}
	]
	`))

	return c, nil
}

func TestInsertData(t *testing.T) {
	// c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }
	// db := c.Database("tronics")
	// col := db.Collection("products")
	mockCol := &mockCollection{} // 上の設定はなしでmockを使用する
	user := User{"krunal", "shimpi"}

	res, err := insertData(mockCol, user)
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
	assert.Equal(t, "asdf", res.InsertedID)

	res, err = insertData(mockCol, User{"eric", "cartman"})
	assert.NotNil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}

func TestFindData(t *testing.T) {
	mockCol := &mockCollection{}
	users, err := findData(mockCol)
	assert.Nil(t, err)
	for _, user := range users {
		assert.Equal(t, "krunal", user.FirstName)
		assert.Equal(t, "shimpi", user.LastName)
	}
}
