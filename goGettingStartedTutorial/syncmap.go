package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var m sync.Map // 不需要通过make初始化

func aaa(n int) {
	// map写的操作
	key := strconv.Itoa(n)
	m.Store(key, n)
	//map读操作
	fmt.Println(m.Load(key))
	wg.Done()
}

func main() {
	// concurrent map writes
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go aaa(i)
	}
	wg.Wait()
}
