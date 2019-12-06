package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"go-hreq/library"
	"go-hreq/util"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
)

func Hreq(cn *library.MongoLib, m map[string]interface{}) {
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

		do(cn, idx, m)
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
		do(cn, idx, m)
	}
}

// mongo 数据操作
func do(cn *library.MongoLib, idx int, m map[string]interface{}) {
	//失败 并且 req_num + 1 == m['num']
	req, _ := m["req_num"].(int32)
	num := m["num"].(int32)
	id := m["id"].(string)
	reqNum := req + 1

	//如果成功 or 访问次数已经到最大，则删除
	if idx > -1 || reqNum >= num {
		cn.Delete(bson.M{"id": id})
	} else {
		//req_num +1
		cn.UpdateNumById(id, reqNum)
	}
}

// 待使用,待测试
type Req struct {
	Id     string
	Url    string
	Info   string
	Method string
	Num    int32
	Resp   string
	ReqNum int32
}

func (r *Req) Request() (int, error) {
	method := strings.ToUpper(r.Method)
	idx := -1
	var err error
	switch method {
	case "GET":
		url := fmt.Sprintf("%s?%s", r.Url, r.Info)
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		//set header
		req.Header.SetContentType("application/json")
		req.Header.SetMethod("GET")
		req.SetRequestURI(url)

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if err = fasthttp.Do(req, resp); err != nil {
			return idx, err
		}

		body := resp.Body()

		idx = strings.Index(string(body), r.Resp)

		return idx, nil
	case "POST":
		//todo
		return idx, nil
	}

	return idx, err
}

func (r *Req) Do(cn *library.MongoLib, idx int) {

	id := r.Id
	req := r.ReqNum + 1
	if idx > -1 || req >= r.Num {
		cn.Delete(bson.M{"id": id})
	} else {
		//req_num +1
		cn.UpdateNumById(id, req)
	}
}
