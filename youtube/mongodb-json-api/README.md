## Installing mongodb
### Installing mongodb with Docker
```bash
docker run --name mongo-test -p 27017:27017 -d mongo
```

### Go dependencies
```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
```

### Mongo Golang quickstart
```go
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
if err != nil {
    panic(err)
}
```
