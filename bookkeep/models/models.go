package models

import (
	"bookkeep/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Account struct {
	Name    string   `json:"name"`
	Pwd     string   `json:"pwd"`
	Balance int      `json:"balance"`
	Details []Detail `json:"details"`
	isLogin bool
}
type Detail struct {
	Kind    string `json:"kind"`
	Amounts int    `json:"amounts"`
	Message string `json:"message"`
}

func NewAccount(name, pwd string) *Account {
	return &Account{
		Name:    name,
		Pwd:     pwd,
		Balance: 0,
	}
}

func SaveAccount(account *Account) {
	data, _ := json.Marshal(account)
	os.WriteFile(account.Filepath(), data, 0666)
}

func GetAccount(name string) *Account {
	filepath, ok := utils.IsExist(name)
	if !ok {
		return nil
	}
	accountJson, _ := ioutil.ReadFile(filepath)
	var account Account
	json.Unmarshal(accountJson, &account)
	return &account
}

func (ac Account) Filepath() string {
	return path.Join("accounts", fmt.Sprint(ac.Name, ".json"))

}

func (ac *Account) ShowBalance() int {
	return ac.Balance
}

func (ac *Account) ShowBalanceDetail() []Detail {
	return ac.Details
}

func (ac *Account) UpBalance(amounts int, message string) {

	ac.Balance += amounts
	detail := Detail{
		Kind:    "收入",
		Amounts: amounts,
		Message: message,
	}
	ac.Details = append(ac.Details, detail)

}

func (ac *Account) DownBalance(amounts int, message string) {
	ac.Balance -= amounts
	// if amounts > ac.Balance {
	// 	fmt.Println("余额不足！")
	// 	return
	// }
	detail := Detail{
		Kind:    "支出",
		Amounts: amounts,
		Message: "message",
	}
	ac.Details = append(ac.Details, detail)

}
