package main

import (
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
	// defer file.Close() 方式二：defer语句，操作完毕之后，关闭文件
	tmp := make([]byte, 128) //每次最多读128个字节
	n, err := file.Read(tmp)
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Printf("读取文件夹出错了：%v", err)
		return
	}
	fmt.Printf("读了%d个字节\n", n)
	fmt.Println(string(tmp))
	file.Close() // 方式一：操作完毕之后，关闭文件

	/*
		循环读整个文件
	*/
	// var tmp = make([]byte, 128)
	// for{
	//     n, err := file.Read(tmp)
	//     if err == io.EOF {
	//         fmt.Println("文件读完了")
	//         return
	//     }
	//     if err != nil {
	//         fmt.Println("文件打开失败", err)
	//         return
	//     }
	//     fmt.Printf("读取了%d个字节\n", n)
	//     fmt.Println(string(tmp))
	// }
}
