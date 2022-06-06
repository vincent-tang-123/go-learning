package main

import "fmt"

func main() {
	//var ch1 chan int // 引用类型 需要初始化之后才能使用
	//ch1 = make(chan int, 1)

	// 带缓冲区的通道相当于快递柜是一个队列， size 大小相当于快递柜大小  又称为 异步通道
	ch1 := make(chan int, 1) // 声明并初始化通道  size 是 1   带缓冲区通道


    // 无缓冲区的通道相当于阻塞了是同步的，需要马上消费    又称为 同步通道
	//ch1 := make(chan int) // 不带缓冲区通道  fatal error: all goroutines are asleep - deadlock!
	ch1 <- 20  // 将 20 装进通道中

	fmt.Println(len(ch1)) // len 取通道中元素的数量  1

	x := <-ch1 // 从通道中取出值 赋值给 x
	fmt.Println(x) // 20

	//len(ch1)
	fmt.Println(len(ch1)) // len 取通道中元素的数量  0

	// cap(ch1)
	fmt.Println(cap(ch1)) // cap 拿到通道的容量  1

	close(ch1) // 关闭通道

}
