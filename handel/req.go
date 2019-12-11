package handel

import (
	"fmt"
	"strconv"

	"go-hreq/config"
	"go-hreq/library"
	"go-hreq/response"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func addReq(c echo.Context) error {
	method := c.FormValue("method")
	url := c.FormValue("url")
	num := c.FormValue("num")
	info := c.FormValue("info")
	ckRes := c.FormValue("checkResp")

	rn, _ := strconv.Atoi(num)

	// 写入mongo
	id := uuid.NewV4().String()
	res := bson.M{"id": id, "url": url, "method": method, "num": rn, "req_num": 0, "info": info, "resp": ckRes}

	con := library.MongoLib{}
	conErr := con.MongoClient()
	if conErr != nil {
		fmt.Println(conErr)
	}

	dbErr := con.SetDB(config.MongoConfig["databases"])
	if dbErr == false {
		return nil
	}
	con.SetTable(config.MongoConfig["tb"])
	addErr := con.Add(res)
	if addErr != nil {
		return addErr
	}

	response.Response(c, 200, "", "success")

	return nil
}
