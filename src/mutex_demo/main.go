package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex // 互斥锁 保证同一时间只有一个 goroutine 访问共享资源
)

func add(){
	for i:=0;i<5000;i++{
		lock.Lock() // 加锁
		x+=1
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
