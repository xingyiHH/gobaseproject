package main

import (
	"fmt"
)

func f1() {
	fmt.Println("111")
}

func f2() {
	defer func() {
		// if判断的特殊写法
		if err := recover(); err != nil {
			// 做异常处理
			fmt.Println(err)
		}
		// 不论有没有出现异常，需要的操作写在下面
		fmt.Println("不论有没有出现异常，需要的操作写在下面")

	}()
	fmt.Println("222")
	// panic("出错误了")
}

func f3() {
	fmt.Println("333")
}

func main() {
	f1()
	f2()
	f3()
}
