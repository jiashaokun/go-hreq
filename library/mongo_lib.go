package library

import (
	"context"
	"fmt"
	"time"

	"go-hreq/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoLib struct {
	Connect *mongo.Client
	DB *mongo.Database
	TB *mongo.Collection
}


func (m *MongoLib) MongoClient() error {
	host := fmt.Sprintf("mongodb://%s:%s", config.MongoConfig["host"], config.MongoConfig["port"])
	opts := &options.ClientOptions{}
	opts.SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, err := mongo.Connect(ctx, options.Client().ApplyURI(host), opts)

	if err != nil {
		return nil
	}
	m.Connect = con

	return nil
}

func (m *MongoLib) SetDB(db string) bool {
	if m.Connect == nil {
		return false
	}

	m.DB = m.Connect.Database(db)

	return true
}

// set table
func (m *MongoLib) SetTable(table string) {
	m.TB = m.DB.Collection(table)
}

func (m *MongoLib) Delete(c bson.M) error {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	_, err := m.TB.DeleteOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoLib) Add(c bson.M) error {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	_, err := m.TB.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoLib) Find(c bson.M) ([]bson.M,error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	val, err := m.TB.Find(ctx, c)

	if err != nil{
		return nil, err
	}
	var result []bson.M
	for val.Next(ctx){
		var res bson.M
		ersErr := val.Decode(&res)
		if ersErr != nil {
			return nil, ersErr
		}
		result = append(result, res)
	}
	return result, nil
}

func (m *MongoLib) UpdateNumById(id string, num int32) error {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	w := bson.M{"id": id}
	u := bson.M{"$set": bson.M{"req_num":num}}
	_, err := m.TB.UpdateOne(ctx, w, u)
	if err != nil {
		return err
	}
	return nil
}
