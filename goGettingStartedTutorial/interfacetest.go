package main

import "fmt"

// 定义一个接口 animal
type animal interface {
	say()
}

type cat struct {
	name string
}

// cat这个结构体实现了say方法
func (c cat) say() {
	fmt.Printf("%v叫了,喵喵喵\n", c.name)
}

type dog struct {
	name string
}

// dog这个结构体实现了say方法
func (d dog) say() {
	fmt.Printf("%v叫了,旺旺旺\n", d.name)
}

func aaa(a animal) {
	a.say()
}

func main() {

	c := cat{"蓝猫"}
	d := dog{"哮天犬"}

	// c.say()
	// d.say()

	aaa(c)
	aaa(d)
}
