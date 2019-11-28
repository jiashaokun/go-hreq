package pkg

import (
	"fmt"
	"sync"
	"strings"
	"net/http"
	"io/ioutil"

	"go-hreq/util"
	"go-hreq/library"

	"go.mongodb.org/mongo-driver/bson"
)

func Hreq(cn *library.MongoLib, wg *sync.WaitGroup, m map[string]interface{})  {
	method := strings.ToUpper(m["method"].(string))
	switch method {
	case "GET":
		domain := fmt.Sprintf("%s?%s", m["url"], m["info"])
		resp, _ := http.Get(domain)
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		str := string(body)

		respString := m["resp"].(string)
		idx := strings.Index(str, respString)

		do(cn, wg, idx, m)
	case "POST":
		domain := m["url"].(string)
		path := m["info"].(string)

		pms := util.GetPostUrlPath(path)
		resp, _ := http.PostForm(domain, pms)
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		str := string(body)

		respString := m["resp"].(string)
		idx := strings.Index(str, respString)
		do(cn, wg, idx, m)
	}
}

// mongo 数据操作
func do(cn *library.MongoLib, wg *sync.WaitGroup, idx int, m map[string]interface{}) {
	//失败 并且 req_num + 1 == m['num']
	req, _ := m["req_num"].(int32);
	num := m["num"].(int32)
	id := m["id"].(string)
	reqNum := req + 1

	//如果成功 or 访问次数已经到最大，则删除
	if idx > -1 || reqNum >= num {
		cn.Delete(bson.M{"id":id})
	} else {
		//req_num +1
		cn.UpdateNumById(id, reqNum)
	}
}
