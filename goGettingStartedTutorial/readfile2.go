package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("打开文件出错了：%v", err)
		return
	}
	defer file.Close()

	//基于bufio读文件，按行读文件
	reader := bufio.NewReader(file)
	for {
		// 告诉程序这样问价结束的标志，是字符'\n'
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Printf("读文件出错了：%v", err)
			return
		}
		fmt.Print(line)
	}
}
