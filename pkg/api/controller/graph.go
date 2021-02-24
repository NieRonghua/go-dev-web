package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-dev-web/pkg/api/utils"
	"strconv"
)

//type ResponseData struct {
//	Info interface{}
//}

type RequestData struct {
	Name string `form:"name" json:"name" uri:"name" binding:"required"`
	Age  int    `form:"age" json:"age" uri:"age" binding:"required"`
}

type ResJsonData struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 错误消息
	Data interface{} `json:"data"` // 数据
}

//GET请求：http://localhost:7788/?name=%27tom%27&age=18
func Index(ctx *gin.Context) {
	var (
		req = new(RequestData)
		rsp = new(ResJsonData)
		err error
	)
	fmt.Println("index")
	// 绑定参数至结构体
	if err = ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(utils.MakeReply(rsp.Data, err))
		return
	}

	// 返回数据
	Student := make(map[string]string)
	Student["name"] = req.Name
	Student["age"] = strconv.Itoa(req.Age)
	rsp.Data = Student
	ctx.JSON(utils.MakeReply(rsp.Data, err))
}

//POST请求 curl --location --request POST 'http://localhost:7788/hello' \
//--header 'Content-Type: application/json' --data-raw ' { "name": "Jack", "age": 21 }'
func HelloWorld(ctx *gin.Context) {
	var (
		req = new(RequestData)
		rsp = new(ResJsonData)
		err error
	)
	// 绑定参数至结构体
	if err = ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(utils.MakeReply(rsp.Data, err))
		return
	}
	// 返回数据
	Student := make(map[string]string)
	Student["name"] = req.Name
	Student["age"] = strconv.Itoa(req.Age)
	rsp.Data = Student
	ctx.JSON(utils.MakeReply(rsp.Data, err))
}
