package utils

import (
	"crypto/md5"
	"fmt"
	"os"
	"path"
)

func IsExist(name string) (string, bool) {
	filepath := path.Join("accounts", fmt.Sprint(name, ".json"))
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return filepath, false
	}
	return filepath, true
}

func Md5Salt(str string) string {
	salt := "猜猜我是谁"
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte(salt)) //加盐
	return fmt.Sprintf("%x", h.Sum(nil))
}
