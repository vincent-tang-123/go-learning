package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁 读的场景远远大于写的场景

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwLock sync.RWMutex
)

func read(){
	//lock.Lock()
	rwLock.RLock()  // 读写锁
	time.Sleep(time.Microsecond)
	rwLock.RUnlock()
	wg.Done()
}

func write(){
	//lock.Lock()
	rwLock.Lock()
	x += 1
	time.Sleep(time.Microsecond*10)
	//lock.Unlock()
	rwLock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i:=0;i<1000;i++{
		wg.Add(1)
		go read()
	}

	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
