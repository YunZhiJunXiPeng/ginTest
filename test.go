package main

/**
** @ Author       : SmallSuperMan
** @ Date         : 2020-09-16 00:28:31
** @ LastEditTime : 2022-05-16 00:25:40
** @ LastEditors  : SmallSuperMan
** @ FilePath     : \000_Web\ginTest\test.go
** @ Description  :
**/

import (
	"fmt"
	"path"
	"path/filepath"
)

func test() int {
	pathFileString := "C:/目录/文件名.后缀名"
	paths, fileName := filepath.Split(pathFileString)
	fmt.Println(paths, fileName)               // 获取路径中的目录及文件名
	fmt.Println(filepath.Base(pathFileString)) // 获取路径中的文件名
	fmt.Println(path.Ext(pathFileString))      // 获取路径中的文件的后缀
	return 0
}
