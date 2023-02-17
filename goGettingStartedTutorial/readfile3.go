package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 基于ioutil下的ReadFile
	content, err := ioutil.ReadFile("./test.txt")

	if err != nil {
		fmt.Println("打开文件或读文件出错了", err)
		return
	}
	fmt.Println(string(content))
}
