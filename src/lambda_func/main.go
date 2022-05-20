package main

import (
	"fmt"
	"strings"
)

// 匿名函数就是没有函数名的函数 匿名函数多用于回调函数或闭包

//匿名函数的定义格式如下
/*
func(参数)(返回值){
	函数体
}
*/


// 定一个函数返回值是一个函数
// 把函数作为返回值
func a(name string) func(){
	//name := "beijing"
	return func(){
		fmt.Println("hello", name)
	}
}

// 使用闭包做文件后缀名检测
func makeSuffixFunc(suffix string) func(string)string{
	return func(name string)string{
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int)int, func(int)int){
	add := func(i int)int {
		base += i
		return base
	}

	sub := func(i int)int{
		base -= i
		return base
	}
	return add, sub
}


func main() {
	//func(x,y int){
	//	fmt.Println(x+y)
	//}(10, 20) // 匿名函数定义完加（）直接执行
	//add_num := func(x,y int){
	//	fmt.Println(x+y)
	//}// 匿名函数定义完加（）直接执行
	//add_num(10, 20)

	// 闭包 = 函数 + 外层变量的引用
	//names := "vincent"
	//r:=a(names) // r 此时就是一个闭包
	//r() // 相当于执行了a函数内部的匿名函数

	r1 := makeSuffixFunc(".text")
	ret := r1("北京")
	fmt.Println(ret) // 北京.text

	x, y := calc(100)
	ret1 := x(200)
	fmt.Println(ret1)
	ret2 := y(200)
	fmt.Println(ret2)

}
