package library

import (
	"fmt"
	"testing"

	"go-hreq/config"

	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongoLib_MongoClient(t *testing.T) {

	connect := new(MongoLib)
	conErr := connect.MongoClient()
	if conErr != nil {
		fmt.Println(conErr)
	}

	dbErr := connect.SetDB(config.MongoConfig["databases"])
	if dbErr != true {
		t.Fatal("Mongo DB Set DB was wrong")
	}

	connect.SetTable(config.MongoConfig["tb"])

	u2 := uuid.NewV4().String()
	fmt.Println(u2)

	ash := bson.M{"id": u2, "url":"http://cyapi.dev.xin.com", "methon":"post", "num":2, "req_num":0, "info": "a=1&b=2", "resp": "{\"code\":1"}
	addErr := connect.Add(ash)
	if addErr != nil {
		t.Fatal("insert mongo was wrong")
	}

	// 检查新增
	fd := bson.M{"num": bson.M{"$lt": 3}}
	fdVal, fdErr := connect.Find(fd)
	if fdErr != nil {
		t.Fatal("Mongo DB Find Was Wrong !!!")
	}
	fmt.Println(fdVal)


	/*
	del := bson.M{"id": u2}
	delErr := connect.Delete(del)
	if delErr != nil {
		t.Fatal("Mongo DB Delete Was Wrong !!!")
	}

	 */
}
