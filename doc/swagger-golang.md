## goswagger 
- [文档](https://goswagger.io/)
- go-swagger对标openapi 2.0. 一种RESTful API的表示形式
- openapi 一种规范，[文档地址](https://github.com/OAI/OpenAPI-Specification)

### swagger 格式
- meta: 元数据 常见有如下
	- consumes： 字符串数组，可识别的客户端 mime类型，具体支持的可见[地址](https://goswagger.io/use/spec/meta.html#supported-mime-types)
	- produces：字符串数组，可返回的mime类型
	- basepath: 设置的url前缀
	- license：
	- schemes: 字符串数组，两种类型http和https
- paths: url设置
	- [url]
		- [method]：方法，参数，返回值, 是否弃用等
- definitions: 数据格式定义
	- object对应 interface{}
	- boolean 对应 bool
	- number 
		- format double  对应double
		- format float 对应 float
	- interger 
		- format int64/int32/uint32/uint64 对应 ...
	- file  对应 io.ReadCloser
	- allOf 对应 interface{}	(但是会对数据一一转换，看是否符合其中一种类型数据)

### golang 框架
- [gin godoc](https://godoc.org/github.com/gin-gonic/gin) 
- [echo](https://github.com/labstack/echo)
- [beego](https://github.com/astaxie/beego )
- [Buffalo](https://github.com/gobuffalo/buffalo )
- [iris](https://github.com/kataras/iris )
- [revel](https://github.com/revel/revel )

- [常见的middleware](https://github.com/labstack/echo/tree/v2/middleware)

```
gin/echo: 封装net/http库,提供更简单的调用方式
iris/revel/beego: 提供除封装外，还有module,controller, session,template等功能
```

### 服务开发
- 并发控制
- 认证授权
- 超时控制（TODO） 
- 熔断 (TODO)
- 降载 （TODO）


### golang 一些
- static analysis 
	- vet (go tool vet help)
	- [revive](https://github.com/mgechev/revive)




