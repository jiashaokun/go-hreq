# hreq
http repeat query
该项目辅助 API 开发时的错误和失败的请求提供再重试


[项目依赖]
---
    1. MongoDB
    2. Redis (后续添加,用于缓存)
    

[如何使用]
---
```shell
cd cmd
go build con.go con
go build main.go api
```

[API 调用]
```shell
横线部分是正确的sn，该信息可以在 middle encrypt.go 里面删除
```

<img src="https://github.com/jiashaokun/doc/blob/master/txt/go-hrep-api.jpg?raw=true">


---


[不用支持,做个样子而已]

<img src="https://github.com/jiashaokun/doc/blob/master/txt/pay.jpg?raw=true" width="300" heigth="300">