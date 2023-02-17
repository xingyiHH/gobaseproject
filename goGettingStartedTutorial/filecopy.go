/*
实战：文件拷贝
需求：编写一个函数，实现文件拷贝的功能

	仅输入目标文件和源文件，就可以实现文件内容的拷贝

实现1：边读边写
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(dstName, srcName string) {
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

	// 基于bufio实现的读写
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读完了")
			break
		}
		if err != nil {
			fmt.Println("读文件出错了", err)
			return
		}
		_, err = writer.WriteString(line)
		fmt.Println(line)
		if err != nil {
			fmt.Println("写文件出错了", err)
			return
		}
	}

	writer.Flush()
}

func main() {
	copyFile("./dst.txt", "./src.txt")
}
