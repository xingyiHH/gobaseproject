package main

import "fmt"

type Person1 struct {
	name   string
	age    int
	gender string
}

// 构造函数
func newPerson(name, gender string, age int) Person1 {
	return Person1{name: name, gender: gender, age: age}
}
func newPerson2(name, gender string, age int) *Person1 {
	return &Person1{name: name, gender: gender, age: age}
}
func main() {
	p1 := newPerson2("小明", "male", 19)
	fmt.Println(p1)
}
