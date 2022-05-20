package main

import "fmt"

// panic or recover

func a(){
	fmt.Println("a")
}

func b(){
	// defer 一定要在可能引发 panic 的语句定义之前
	defer func(){
		// 注意 recover 必须搭配 defer 使用
		err := recover() // 捕获函数内部的错误  使程序正常运行
		if err != nil{
			fmt.Println("func b error")
		}
	}() // 定义完匿名函数后要调用匿名函数

	panic("panic  in b")
	fmt.Println("b")
}

func c(){
	fmt.Println("c")
}

func main() {
	a()
	b()
	c()
}
