package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	str := "hello world"
	file.Write([]byte(str))         //写入字节切片数据
	file.WriteString("hello world") //直接写入字符串数据

}
