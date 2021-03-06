package pkg

import (
	"fmt"
	"strings"

	"go-hreq/library"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
)

// 待使用,待测试
type Req struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Method string `json:"method"`
	Num    int32  `json:"num"`
	ReqNum int32  `json:"req_num"`
	Info   string `json:"info"`
	Resp   string `json:"resp"`
}

// 发起一个请求
func (r *Req) Request() (int, error) {
	method := strings.ToUpper(r.Method)
	idx := -1
	var err error
	// 区分请求方式，方便后续扩展head等特定属性
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
		url := fmt.Sprintf("%s?%s", r.Url, r.Info)
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		//set header
		req.Header.SetContentType("application/json")
		req.Header.SetMethod("POST")
		req.SetRequestURI(url)

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if err = fasthttp.Do(req, resp); err != nil {
			return idx, err
		}

		body := resp.Body()
		idx = strings.Index(string(body), r.Resp)
	}

	return idx, err
}

func (r *Req) Do(cn *library.MongoLib, idx int) {

	req := r.ReqNum + 1
	if idx > -1 || req >= r.Num {
		cn.Delete(bson.M{"id": r.Id})
	} else {
		//req_num +1
		cn.UpdateNumById(r.Id, req)
	}
}
