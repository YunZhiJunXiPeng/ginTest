/**
** @ Author       : SmallSuperMan
** @ Date         : 2020-09-07 23:39:56
** @ LastEditTime : 2022-11-23 00:50:41
** @ LastEditors  : SmallSuperMan
** @ FilePath     : \ginTest\main.go
** @ Description  :
**/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("你好,世界！\n")
	// 创建服务
	ginServer := gin.Default()
	// 访问地址
	ginServer.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "你好世界"})
	})
	// 服务端口
	ginServer.Run(":8080")
}
