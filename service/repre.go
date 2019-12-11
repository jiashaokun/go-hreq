package service

import (
	"sync"

	"go-hreq/config"
	"go-hreq/library"
	"go-hreq/pkg"
)

// 重试数据库中访问没有超过次数的请求,超过次数则删除
func Repre() {
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

	if err := connect.FindAll(); err != nil {
		panic("Mongo DB Find ALL was wrong")
	}

	if len(connect.Value) > 0 {
		wg := sync.WaitGroup{}

		for _, v := range connect.Value {
			req := pkg.Req{
				Id:     v.Id,
				Url:    v.Url,
				Info:   v.Info,
				Method: v.Method,
				Num:    v.Num,
				Resp:   v.Resp,
				ReqNum: v.ReqNum,
			}
			wg.Add(1)
			go func(req *pkg.Req, cn *library.MongoLib) {
				defer wg.Done()
				idx, _ := req.Request()
				req.Do(connect, idx)
			}(&req, connect)

		}
		wg.Wait()
	}
}
