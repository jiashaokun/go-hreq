package handel

import (
	"fmt"
	"net/http"
	"strconv"

	"go-hreq/config"
	"go-hreq/library"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func addReq(c echo.Context) error {

	src := c.FormValue("src")

	method := c.FormValue("method")
	url := c.FormValue("url")
	num := c.FormValue("num")
	info := c.FormValue("info")

	rn, _ := strconv.Atoi(num)

	// 写入mongo
	id := uuid.NewV4().String()
	res := bson.M{"id": id, "url": url, "methon":method, "num": rn, "info": info}

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

	return c.String(http.StatusOK, src)
}
