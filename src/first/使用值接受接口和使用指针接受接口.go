package main

import "fmt"

// 使用值接收者接首接口

type mover interface {
	move()
}

type person struct {
	name string
	age int
}

//使用值接受接口: 类型的值和类型的指针都能狗保存到接口变量中
//func (p person) move(){
//	fmt.Printf("%s谁在跑...\n", p.name)
//}

// 使用指针接受者实现接口： 只有类型指针能够保存到接口变量中
func (p *person) move(){
	fmt.Printf("%s谁在跑...\n", p.name)
}

func (p *person) say(){
	fmt.Printf("%s谁在叫...\n", p.name)
}




type sayer interface {
	say()
}

// 接口的嵌套
type animal interface {
	mover
	sayer
}


// 空接口的定义
// 空接口是指没有定义任何方法的接口。因此任何礼物ii选哪个都实现了空接口
// 空接口类型变量可以存储任意类型的变量



func main() {
	//var m mover
	////p1 := person{ // p1 是person 类型的值
	////	name: "小王子",
	////	age: 18,
	////}
	p2 := &person{// p2 是 person 类型的指针
		name: "娜扎",
		age: 18,

	}
	////m = p1 // 无法赋值，因为p1是person类型的值没有实现mover结偶
	//m.move()
	//m= p2
	//m.move()
	//fmt.Println(m)


	var a animal // 定一个 animal 类型的变量

	a = p2
	a.move()
	a.say()
	fmt.Println(a)

}
