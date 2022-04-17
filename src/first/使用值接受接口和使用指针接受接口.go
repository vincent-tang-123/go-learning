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

// 使用值接受接口: 类型的值和类型的指针都能狗保存到接口变量中
//func (p person) move(){
//	fmt.Printf("%s谁在跑...\n", p.name)
//}

// 使用指针接受者实现接口： 只有类型指针能够保存到接口变量中
func (p *person) move(){
	fmt.Printf("%s谁在跑...\n", p.name)
}

func main() {
	var m mover
	p1 := person{ // p1 是person 类型的值
		name: "小王子",
		age: 18,
	}
	p2 := &person{// p2 是 person 类型的指针
		name: "娜扎",
		age: 18,

	}
	m = p1 // 无法赋值，因为p1是person类型的值没有实现mover接口
	m.move()
	m= p2
	m.move()
	fmt.Println(m)
}
