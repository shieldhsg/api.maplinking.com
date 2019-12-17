/*
@Time : 2019/12/12 11:33 上午
@Author : 伏火
@File : httpTest.go
@Software: GoLand
*/
package api

import (
	"github.com/gin-gonic/gin"
	"go_project/common"

)

func HttpTest(route *gin.RouterGroup){
	route.GET("/get",testGet)
	route.POST("/post",testPost)
	route.PUT("/put",testPut)
	route.DELETE("/delete",testDelete)
	route.PATCH("/patch",testPatch)
	route.HEAD("/head",testHead)
	route.OPTIONS("/options",testOptions)
}

func testGet (c *gin.Context) {
	c.JSON(common.HTTP_OK,gin.H{"message":common.HTTP_OK_MESSAGE})
}

func testPost(c *gin.Context){
	c.JSON(common.HTTP_OK,gin.H{"message":common.HTTP_OK_MESSAGE})
}

func testPut(c *gin.Context){
	c.JSON(common.HTTP_OK,gin.H{"message":common.HTTP_OK_MESSAGE})
}

func testDelete(c *gin.Context){
	c.JSON(common.HTTP_OK,gin.H{"message":common.HTTP_OK_MESSAGE})
}

func testPatch(c *gin.Context){
	c.JSON(common.HTTP_OK,gin.H{"message":common.HTTP_OK_MESSAGE})
}

func testHead(c *gin.Context){
	c.JSON(common.HTTP_OK,gin.H{"message":common.HTTP_OK_MESSAGE})
}

func testOptions(c *gin.Context){
	c.JSON(common.HTTP_OK,gin.H{"message":common.HTTP_OK_MESSAGE})
}
