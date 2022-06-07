package main

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的map

var m = make(map[int]int)

var m2 = sync.Map{}

var (
	wg sync.WaitGroup
)

func get(key int)int{
	return m[key]
}

func set(key int, value int){
	m[key] = value
}

//func main() {
//	for i:=0;i<500;i++{
//		wg.Add(1)
//		// 并发 匿名函数
//		go func(i int){
//			set(i, i+100) // 设置键值对
//			fmt.Printf("key:%v value:%v", i, get(i)) // 打印键值对
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}


func main() {
	for i:=0;i<500;i++{
		wg.Add(1)
		// 并发 匿名函数
		go func(i int){
			m2.Store(i, i+100) // map 中设置键值对
			//set(i, i+100) // 设置键值对
			value, _ := m2.Load(i)
			fmt.Printf("key:%v value:%v\n", i, value) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
