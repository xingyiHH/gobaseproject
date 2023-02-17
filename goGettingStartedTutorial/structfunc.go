package main

import "fmt"

// 结构体做函数的参数
//
type Person struct {
	name   string
	age    int
	gender string
}

func prinInfo(person Person) {
	fmt.Println("姓名：", person.name)
	fmt.Println("年龄：", person.age)
	fmt.Println("性别：", person.gender)
}

func prinInfo2(person *Person) {
	// fmt.Println("姓名：", (*person).name)
	fmt.Println("姓名：", person.name) // 简写
	fmt.Println("年龄：", person.age)
	fmt.Println("性别：", person.gender)
}

func editName(person Person, newName string) {
	person.name = newName
	fmt.Println(person)
}

func editName2(person *Person, newName string) {
	person.name = newName
	fmt.Println(person)
}

func main() {
	p1 := Person{"小明", 18, "男"}
	// prinInfo(p1)
	prinInfo2(&p1)
	// editName(p1, "大明")
	editName2(&p1, "大明")
	fmt.Println(p1)
}
