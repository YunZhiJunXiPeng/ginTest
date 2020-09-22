package main

/**
** @ Author       : SmallSuperMan
** @ Date         : 2020-09-07 23:39:56
** @ LastEditTime : 2020-09-22 22:53:28
** @ LastEditors  : SmallSuperMan
** @ FilePath     : \ginTest\main.go
** @ Description  :
**/

import (
	"fmt"
	// "log"
	// "net/http"
	// "path"
	// "path/filepath"
	// "github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("你好,世界！\n")
	_ = test()
	// router := gin.Default()
	// router.POST("/upload", func(c *gin.Context) {
	// 	// 从表单取文件
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		c.String(http.StatusBadRequest, "Bad request")
	// 		return
	// 	}
	// 	files := file.Filename
	// 	paths, fileName := filepath.Split(files)
	// 	fmt.Println(paths, fileName)      //获取路径中的目录及文件名
	// 	fmt.Println(filepath.Base(files)) //获取路径中的文件名
	// 	fmt.Println(path.Ext(files))      //获取路径中的文件的后缀
	// 	// 打印文件名称包括路径
	// 	log.Println(file.Filename)
	// 	// 上传到项目根目录  名字不变
	// 	c.SaveUploadedFile(file, fileName)
	// 	// 打印信息
	// 	c.String(http.StatusCreated, fmt.Sprintf("'%s' upload!!!!!!!!!", file.Filename))
	// })
	// router.Run(":8000")
}
