package main

import "fmt"

// 为任意类型添加方法

// 基于内置的基本类型造一个我们自己的类型

type MyInt int

func (m MyInt)sayHi(){
	fmt.Println("hi")
}

func main() {
	var a MyInt
	a = 100
	fmt.Println(a) // 100
	a.sayHi() // hi
}
