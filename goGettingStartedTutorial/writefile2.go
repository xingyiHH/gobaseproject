package main

import (
	"bufio"
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
	//基于bufio.NewWriter写文件
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("基于bufio.NewWriter写文件\n")
	}
	writer.Flush() //将缓存中的内容刷新到文件
}
