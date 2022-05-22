package main

import "fmt"

// 结构体构造函数： 构造一个结构体实例的函数
// 结构体实值类型
type person struct {
	name, city string
	age int8
}

// 构造函数 实例化一个结构体 并返回值
func newPerson(name, city string, age int8) person{
	return person{
		name: name,
		city: city,
		age: age,
	}
}

// 构造函数，实例化一个结构体 并返回指针
func newPerson1(name, city string, age int8) *person{
	return &person{
		name: name,
		city: city,
		age: age,
	}
}


func main() {
	p1 := newPerson("vincent", "beijing", int8(18))
	fmt.Printf("type: %T value: %#v\n", p1, p1) // type: main.person value: main.person{name:"vincent", city:"beijing", age:18}

	p2 := newPerson1("vincent", "beijing", int8(18))
	fmt.Printf("type: %T value: %#v\n", p2, p2) //type: *main.person value: &main.person{name:"vincent", city:"beijing", age:18}
	fmt.Printf("type: %T value: %#v\n", p2, *p2) //type: *main.person value: main.person{name:"vincent", city:"beijing", age:18}

}
