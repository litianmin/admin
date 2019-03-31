package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建一个文件夹
	// 首先判断该文件路径是否存在
	filePath := "./admin/image/"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("该文件夹不存在")
		// 创建文件夹
		err = os.MkdirAll(filePath, os.ModePerm) // 创建所有文件夹
		if err != nil {
			fmt.Println("创建失败，请仔细检查")
		} else {
			fmt.Println("创建成功")
		}
	}

	// 创建成功文件夹之后，现在来创建文件
	fileName := "test.txt"

	// 先判断该文件是否存在，请用IsNotExist 来判断，不要用IsExist,因为如果文件已经存在的话，是不会报错的
	if _, err := os.Stat(filePath + fileName); os.IsNotExist(err) {
		fmt.Println("文件不存在")

		file, err := os.Create(filePath + fileName) // 这个函数如果文件已经存在，那么会先删除然后创建
		if err != nil {
			fmt.Println("创建文件成功！")
		}
		defer file.Close()

		// 下面往文件来写东西了
		file.Write([]byte("I something happened!"))

	} else {
		fmt.Println("文件已经存在了")

		// os.Open(filePath + fileName)	// 这是Open， 只是只读，不能写的，所以用下面的
		file, _ := os.OpenFile(filePath+fileName, os.O_APPEND, os.ModePerm)
		file.Write([]byte("11--看看有没有删除后面的东西"))
	}

}
