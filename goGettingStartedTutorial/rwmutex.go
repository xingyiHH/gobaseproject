package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func write() {
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(time.Second) // 假设读操作耗时10毫秒
	rwlock.Unlock()         // 解写锁
	wg.Done()
}

func read() {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
