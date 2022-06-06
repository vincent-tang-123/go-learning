package main

import "fmt"

/*
两个 goroutine
	1 生成 0-100 的数字发送到 ch1 中
	2 从 ch1 中取出数据计算它的平方，把结果发送到 ch2 中
*/

// 生成 0-100的数字发送到ch1 中
func f1(ch chan<- int){
	for i:=0;i<100;i++{
		ch <- i // 将数据装入 channel 中
	}
	close(ch) // 操作完 channel 后 关闭 通道
}

// 从 ch1 中取出数据计算它的平方，把结果发送到 ch2 中
func f2(ch1 <-chan int, ch2 chan<- int){   // chan 关键字前加 <-  符号代表此管道只能取  反之之后 只能存

	// 从通道中取值的方式1

	for {
		tmp, ok := <- ch1  // 从 channel 中取出数据
		if !ok { // 取完数据了 结束 for 循环
			break
		}
		ch2 <- tmp*tmp // 将数据装入 ch2 通道中
	}
	close(ch2) // 关闭 ch2 通道
}

func main() {
	ch1 := make(chan int, 100) // 初始化 一个 channel 长度 100
	ch2 := make(chan int, 200)

	go f1(ch1)
	go f2(ch1, ch2)
	// 从通道中取值的方式2
	for ret := range ch2{
		fmt.Println(ret)
	}
}
