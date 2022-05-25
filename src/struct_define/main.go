package main

import "fmt"

// 定义结构体

// 1
//type person struct_01 {
//	name string
//	city string
//	age int
//}

// 2 结构体简化  内存对齐
type person struct {
	name, city string
	age int
}



// 结构体指针
type person1 struct {
	name, city string
	age int
}

func main() {
	var p person
	// 结构体赋值
	p.name = "vincent"
	p.city = "shanghai"
	p.age = 18
	fmt.Printf("p: %#v\n", p) // p: main.person{name:"vincent", city:"shanghai", age:18}


	fmt.Println(p.name) // vincent

	fmt.Println(p.city) // shanghai

	fmt.Println(p.age) // 18


	// 匿名结构体
	var user struct{
		name string
		married bool
	}
	// 结构体赋值
	user.name = "小王子"
	user.married = false

	fmt.Printf("user: %#v\n", user) // user: struct_01 { name string; married bool }{name:"小王子", married:false}


	// 结构体指针
	var p2 = new(person1)
	fmt.Printf("%T\n", p2) // *main.person1

	// 1
	// 指针内省结构体赋值
	//(*p2).name = "vincent"
	//(*p2).city = "beijing"
	//(*p2).age = 18
	//fmt.Printf("%#v\n", p2) // &main.person1{name:"vincent", city:"beijing", age:18}

	// 2
	// 指针内类型结构体赋值简化
	p2.name = "vincent"
	p2.city = "beijing"
	p2.age = 18
	fmt.Printf("%#v\n ", p2) // &main.person1{name:"vincent", city:"beijing", age:18}


	// 取结构体的地址进行实例话
	p3 := person{}
	fmt.Printf("%T\n", p3) // main.person
	// 初始化未传值时，里面元素的值 就是对应类型的 0 值
	fmt.Printf("%#v\n", p3) // main.person{name:"", city:"", age:0}

	p4 := &person{}
	fmt.Printf("%T\n", p4) // *main.person
	// 初始化未传值时，里面元素的值 就是对应类型的 0 值
	fmt.Printf("%#v\n", p4) // &main.person{name:"", city:"", age:0}
	p4.name = "娜扎"
	p4.age = 18
	p4.city = "北京"
	fmt.Println(p4) // &{娜扎 北京 18}  结构体地址变量实例化
	fmt.Println(*p4) // {娜扎 北京 18} * 取值


}
