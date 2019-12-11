package backands

type MongoData struct {
	Id     string "id"
	Url    string "url"
	Method string "method"
	Num    int32  "num"
	ReqNum int32  "req_num"
	Info   string "info"
	Resp   string "resp"
}
