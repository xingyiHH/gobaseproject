package services

import (
	"bookkeep/models"
	"bookkeep/utils"
	"log"
	"os"
	"path"
)

// 自定义日志级别
var (
	Info    *log.Logger // 常规信息
	Warning *log.Logger // 警告信息
	Error   *log.Logger // 一般错误
	Fatal   *log.Logger // 严重错误
)

func init() {
	//设置日志前缀
	log.SetPrefix("[CoolBear]")
	//设置日志保存的内容
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	logFilePath := path.Join("logs", "coolbear.log")
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	//设置日志输出位置
	log.SetOutput(file)
	//自定义不同级别日志
	Info = log.New(file, "[INFO]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
	Warning = log.New(file, "[WARNING]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
	Error = log.New(file, "[Error]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
	Fatal = log.New(file, "[Fatal]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)

}

var currentAccount *models.Account

func LoginService(name, pwd string) (string, bool, *models.Account) {
	//判断name是否存在
	account := models.GetAccount(name)
	if account == nil {
		return "用户名不存在", false, nil
	}
	//账号存在，判断密码
	if utils.Md5Salt(pwd) != account.Pwd {
		return "密码错误", false, nil
	}
	currentAccount = account
	Info.Printf("%s登录成功", currentAccount.Name)
	return "登录成功", true, currentAccount
}

func RegisterService(name, pwd string) (string, bool, *models.Account) {
	account := models.GetAccount(name)
	if account != nil {
		return "当前账号已被注册，请换个账号重新注册", false, nil
	}
	//可以注册，对密码进行加密处理
	newAccount := models.NewAccount(name, utils.Md5Salt(pwd))
	models.SaveAccount(newAccount)
	return "注册成功", true, newAccount

}

func ShowBalanceService() (string, int, bool) {
	if currentAccount == nil {
		return "未登录", 0, false
	}
	return "余额查询成功", currentAccount.ShowBalance(), true

}

func ShowBalanceDetailService() (string, []models.Detail, bool) {
	if currentAccount == nil {
		return "未登录", nil, false
	}
	details := currentAccount.ShowBalanceDetail()
	if len(details) == 0 {
		return "您当前账号没有明细", nil, false
	}
	return "明细查询成功", details, true

}

func UpBalanceService(amounts int, message string) (string, bool) {
	if currentAccount == nil {
		return "未登录", false
	}
	currentAccount.UpBalance(amounts, message)
	models.SaveAccount(currentAccount)
	return "收入记录成功", true
}

func DownBalanceService(amounts int, message string) (string, bool) {
	if currentAccount == nil {
		return "未登录", false
	}
	if amounts > currentAccount.ShowBalance() {
		return "账户余额不足", false
	}
	currentAccount.DownBalance(amounts, message)
	models.SaveAccount(currentAccount)
	return "支出记录成功", true
}
