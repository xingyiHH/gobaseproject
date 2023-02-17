package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "基于ioutil.WriteFile写文件\n"
	err := ioutil.WriteFile("./text.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("打开文件或写入文件错误", err)
		return
	}
}
