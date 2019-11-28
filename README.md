# hreq
http repeat query
该项目辅助 API 开发时的错误和失败的请求提供再重试，异步失败的 http 请求，通过接口添加至 go-hreq ，启动 con ，系统将会把 Mongo 中未完成的 和 失败的继续请求，直到执行相应的次数后，删除该记录。


[项目依赖]
---
    1. MongoDB
    2. Redis (后续添加,用于缓存)
    
[目录结构]
---
    1. cmd 执行文件入口
    2. config 数据库等配置
    3. handel 路由
    4. library 签名及DB封装
    5. middle 中间件
    6. pkg 最终执行
    7. response 结果返回
    8. service 复杂业务逻辑
    9. util 工具

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
---
<img src="https://github.com/jiashaokun/doc/blob/master/txt/pay.jpg?raw=true" width="300" heigth="300">