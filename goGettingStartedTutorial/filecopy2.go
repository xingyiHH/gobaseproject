/*
实战：文件拷贝
需求：编写一个函数，实现文件拷贝的功能

	仅输入目标文件和源文件，就可以实现文件内容的拷贝

实现2：整体读整体写的方式
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func copyFile1(dstName, srcName string) {
	content, err := ioutil.ReadFile(srcName)
	if err != nil {
		fmt.Println("读文件出错", err)
		return
	}
	err = ioutil.WriteFile(dstName, content, 0666)
	if err != nil {
		fmt.Println("写文件出错", err)
		return
	}
}

func main() {
	copyFile1("./dst.txt", "./src.txt")
}
