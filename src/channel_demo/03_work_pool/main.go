package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs<-chan int, results chan<- int){
	for job:= range jobs{
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results<-job*2
		time.Sleep(time.Microsecond*500)
		fmt.Printf("worker:%d end job:%d\n", id, job)

	}

}


func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启 3 个 goroutine  开启goroutine 池 完成5个任务
	for j:=0;j<3;j++{
		go worker(j, jobs, results)
	}

	// 发送5个任务
	for i:=0;i<5;i++{
		jobs <- i
	}
	close(jobs)

	// 输出结果
	// for range 会自动检测 通道关闭情况，因为没有关闭results通道  且通道内无数据所以造成一直娶不到数据 导致死锁
	//for ret := range results{  // fatal error: all goroutines are asleep - deadlock!  死锁
	//
	//	fmt.Println(ret)
	//}
	for i:=0;i<5;i++{  // 显式取3次
		ret:=<-results
		fmt.Println(ret)
	}

	/*
	worker:2 start job:0
	worker:0 start job:1
	worker:1 start job:2
	0
	2
	4
	worker:1 start job:4
	8
	worker:2 start job:3
	6
	*/

}
