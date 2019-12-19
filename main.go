/*
@Time : 2019/12/12 10:48 上午
@Author : 伏火
@File : main.go
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"go_project/api"
)


func main (){
	r := gin.Default()
	//go-gin 基本用法
	r.GET("test",func(c *gin.Context){
		c.JSON(200,gin.H{"message":"获取成功"})
	})
	//各种http方式测试
	api.HttpTest(r.Group("/test"))
	r.Run("7777")
}
