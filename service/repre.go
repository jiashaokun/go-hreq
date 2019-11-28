package service

import (
	"sync"

	"go-hreq/pkg"
	"go-hreq/config"
	"go-hreq/library"

	"go.mongodb.org/mongo-driver/bson"
)

// 重试数据库中访问没有超过次数的请求,超过次数则删除
func Repre(){
	// 获取数据
	connect := new(library.MongoLib)
	conErr := connect.MongoClient()
	if conErr != nil {
		panic(conErr)
	}

	dbErr := connect.SetDB(config.MongoConfig["databases"])
	if dbErr != true {
		panic("Mongo DB Set DB was wrong")
	}

	connect.SetTable(config.MongoConfig["tb"])

	// 检查新增
	fd := bson.M{"req_num": bson.M{"$lt": 3}}
	fdVal, fdErr := connect.Find(fd)

	if fdErr != nil {
		panic("Mongo DB Find Was Wrong !!!")
	}

	// 操作数据发送请求 todo
	wg := sync.WaitGroup{}
	for _, v := range fdVal {
		pkg.Hreq(connect, &wg, v)
		/*
		wg.Add(1)
		go func(cn *library.MongoLib, wg *sync.WaitGroup, v bson.M) {
			defer wg.Done()
			pkg.Hreq(cn, wg, v)
		}(connect, &wg, v)

		 */
	}
	wg.Wait()
}
