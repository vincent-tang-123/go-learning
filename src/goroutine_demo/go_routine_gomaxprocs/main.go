package main

import (
	"fmt"
	"runtime"
	"sync"
)


// 管理 goroutine
var wg sync.WaitGroup


func a(){
	for i:=1;i<10;i++{
		fmt.Println("a", i)
	}
	wg.Done()
}

func b(){
	for i:=1;i<10;i++{
		fmt.Println("b", i)
	}
	wg.Done()
}


func main() {
	runtime.GOMAXPROCS(6) // 开启4个核心
	wg.Add(2) // 添加管理 2个 goroutine
	go a()
	go b()
	wg.Wait()
}
