package main

import (
	"bookkeep/services"
	"fmt"
	"os"
)

// 存放所有注册的账号
//var allAccounts = make(map[string]*Account, 100)

func Register() {
	var (
		name     string
		pwd      string
		pwdAgain string
	)
	fmt.Printf("请输入用户名： ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("请输入密码： ")
	fmt.Scanln(&pwd)
	fmt.Printf("请再次输入密码： ")
	fmt.Scanln(&pwdAgain)
	if pwd != pwdAgain {
		fmt.Println("两次密码输入不一致")
		return
	}

	message, _, _ := services.RegisterService(name, pwd)
	fmt.Println(message)

}

func Login() {
	var (
		name string
		pwd  string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入密码：")
	fmt.Scanln(&pwd)
	// account, ok := allAccounts[name]

	message, _, _ := services.LoginService(name, pwd)
	fmt.Println(message)

}

func ShowBalance() {
	msg, amounts, ok := services.ShowBalanceService()
	fmt.Println(msg)
	if ok {
		fmt.Println("你当前的余额是：", amounts)
	}
}

func ShowBalanceDetail() {
	msg, details, ok := services.ShowBalanceDetailService()
	fmt.Println(msg)
	if ok {
		fmt.Println("你当前的收支明细如下：")
		for index, d := range details {
			fmt.Printf("(%d)%s\t\t%d\t\t%s\n", index+1, d.Kind, d.Amounts, d.Message)
		}
	}
}

func UpBalance() {
	var (
		amounts int
		message string
	)
	fmt.Print("请输入收入金额： ")
	fmt.Scanln(&amounts)
	fmt.Print("请输入收入来源： ")
	fmt.Scanln(&message)
	msg, _ := services.UpBalanceService(amounts, message)
	fmt.Println(msg)

}

func DownBalance() {
	var (
		amounts int
		message string
	)
	fmt.Print("请输入支出金额： ")
	fmt.Scanln(&amounts)
	fmt.Print("请输入支出事由： ")
	fmt.Scanln(&message)

	msg, _ := services.DownBalanceService(amounts, message)
	fmt.Println(msg)

}

func Menu() {
	fmt.Println("欢迎使用酷酷熊记账本软件")
	var choice int
	menuInfo := `
		1  登录
		2  注册
		3  余额
		4  明细
		5  收入
		6  支出
		7  退出
	`
	for {
		fmt.Println(menuInfo)
		fmt.Print("请选择操作编号： ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Login()
		case 2:
			Register()
		case 3:
			ShowBalance()
		case 4:
			ShowBalanceDetail()
		case 5:
			UpBalance()
		case 6:
			DownBalance()
		case 7:
			os.Exit(0) //0表示正常退出，不是异常报错导致的退出
		default:
			fmt.Println("请选择合理的操作编号！")
		}
	}

}

func main() {
	Menu()
}
