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
---

1. src : 签名key对应的来源，该项可在config中配置
2. method : 请求方式 get/post
3. url : 需要重试的url
4. info : 重试的参数 a=b&c=d
5. num : 重试最多次数
6. checkResp : 匹配正确的结果，匹配到则不再重试
7. sn : 签名(签名在 middlware 中，可关闭)
8. 下图有划线部分的注释

<img src="https://github.com/jiashaokun/doc/blob/master/txt/hreq.jpg?raw=true">

---
[不用支持,做个样子而已]
---
<img src="https://github.com/jiashaokun/doc/blob/master/txt/pay.jpg?raw=true" width="300" heigth="300">