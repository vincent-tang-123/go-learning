package main

import (
	"fmt"
)

// 函数

// 定一个不需要参数也没有返回值的函数
func sayHello(){
	fmt.Println("hello word")
}

// 定义一个接受 string 类型的 name 参数
func sayHello1(name string){
	fmt.Println("hello", name)
}

// 定一个接受多个参数的函数
// func 函数名(参数名 参数类型) 返回值类型
// 1 也可以按 2 这种写法
func intSum(a int, b int) int{
	ret := a+b
	return ret
}
// 2 简写                      ret 为定义的返回值 函数体内不需要定义 ，同时 return 后也无需再写
func intSum1(a int, b int) (ret int){
	ret = a + b
	return
}


// 函数可以接受可变参数
// ... 表示可变参数， 在函数体内是切片类型  可以接受 0个或者多个参数
func intSum3(a ...int) int{
	//fmt.Println(a)
	//fmt.Printf("%T\n", a) // 切片
	ret := 0
	for _, arg := range a{
		ret += arg
	}
	return ret
}

// 固定参数和可变参数同时出现时， 可变参数要放在最后
// go 语言中没有默认参数
func intSum4(a int, b...int) int{
	ret := a
	for _, arg := range b{
		ret += arg
	}
	return ret
}


// go 语言中函数参数类型简写
func intSum5(a,b int)int{
	ret:= 0
	ret = a + b
	return ret
}

// 定义具有多个返回值的函数, 多返回值也支持类型简写
func calc(a, b int) (sum, sub int){
	sum = a + b
	sub = a - b
	return
}


func main() {
	// 函数调用
	sayHello()

	//
	name := "vincent"
	sayHello1(name)

	r := intSum(10, 20)
	r1 := intSum1(3,5)
	fmt.Println(r)
	fmt.Println(r1)

	r2 := intSum3(10,20,30)
	r3 := intSum3()
	fmt.Println(r2)
	fmt.Println(r3)

	r4 := intSum4(1) // a=1 ,b=[]
	r5 := intSum4(1, 20,30) // a=1 ,b=[20,30]
	fmt.Println(r4)
	fmt.Println(r5)
	x, y := calc(100, 200)
	fmt.Println(x , y)
}
