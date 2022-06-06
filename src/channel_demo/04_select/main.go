package main

import "fmt"

func main() {
	ch := make(chan int, 1) // 初始化一个 chan 容量为1
	for i :=0;i<10;i++{
		select {
		case x:= <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("啥也不干")
		}
	}
	/*
	0
	2
	4
	6
	8
	*/
}
