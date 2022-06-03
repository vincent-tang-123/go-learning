package main

import (
	"fmt"
	"sync"
)

// goroutine demo

// 监听 所有的 goroutine
var wg sync.WaitGroup


func hello(i int){
	fmt.Println("hello vincent", i)
	wg.Done()  // 通知 wg 把计数器 -1
}

func main() {  // 开启一个主 goroutine 去执行 main 函数

	wg.Add(10000)  // 启动一个 goroutine 并登记 计数牌 +1

	// 循环 执行 goroutine
	for i:=0;i<10000;i++{
		//wg.Add(1)
		go hello(i)  // 开启了一个独立的 goroutine 去执行 hello 函数

	}

	fmt.Println("hello main")
	//time.Sleep(time.Second)

	wg.Wait()  // 等所有小弟都干完活之后收兵   直到计数器为 0 时关闭
}
