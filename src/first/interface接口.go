package main

import "fmt"

// 定义 dog 结构体
type dog struct {

}

// 实现 狗叫的方法
func(d dog)say(){
	fmt.Println("汪汪")
}

// 定义 cat 结构体
type cat struct {

}

type person struct {
	name string
}

// 给人实现 say 函数
func (p person) say(){
	fmt.Println("aaa")
}

// 实现 猫叫的方法
func (c cat)say(){
	fmt.Println("喵喵")
}

// 定义一个类型，定义一个抽象类型，只要实现say 这个方法的类型，都可以称为sayer类型
type sayer interface {
	say()
}
// 接口不管你是什么类型，他只管你要实现什么方法

// 打的函数
func da(arg sayer){
	arg.say() // 不管传进来的函数是什么，打他就要叫，执行speak 方法
}


// 为什么需要接口
func main() {
	// 声明一个 c1 的变量
	//c1 := cat{}
	//da(c1)
	//
	//d1:=dog{}
	//da(d1)
	//
	//p1:=person{
	//	name:"vincent",
	//}
	//da(p1)

	// 实现接口的用处
	// 只要实现了 接口中的方法 就可以将值赋值给接口变量
	var s sayer
	c2 := cat{}
	s = c2
	p2 := person{name: "vincent"}
	s = p2
	fmt.Println(s)

}
