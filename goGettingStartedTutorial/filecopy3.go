/*
实战：文件拷贝
需求：编写一个函数，实现文件拷贝的功能

	仅输入目标文件和源文件，就可以实现文件内容的拷贝

实现3：io.Copy
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile2(dstName, srcName string) {
	// 打开文件读数据
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}
	defer src.Close()

	// 打开文件写数据
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}
	defer dst.Close()

	// 基于io.Copy实现文件拷贝
	io.Copy(dst, src)
}

func main() {
	copyFile2("./dst.txt", "./src.txt")
}
