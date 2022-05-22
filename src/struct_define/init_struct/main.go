package main

import "fmt"

// 结构体初始化

type person struct {
	name, city string
	age int
}


// 结构体内存布局
// 结构体占用一块连续的内存

type test struct {
	a, b, c, d int
}


func main() {

	// 1 键值对初始化
	p4 := person{
		name: "小王子",
		city: "beijing",
		age: 18,
	}
	fmt.Printf("%#v\n", p4) // main.person{name:"小王子", city:"beijing", age:18}
	// 字段可以不全写
	// 最后一行 需要加上逗号
	p5 := &person{
		name: "vincent",
		city: "shanghai",
		age: 18, // 需要加上逗号
	}
	fmt.Printf("%#v\n", p5) // &main.person{name:"vincent", city:"shanghai", age:18}

	// 2 值的列表进行初始化
	// 字段需要对上定义的结构体类型
	p6 := person{
		"vincent",
		"beijing",
		18,
	}
	fmt.Printf("%#v\n", p6) // main.person{name:"vincent", city:"beijing", age:18}

	n := test{
		1,
		2,
		3,
		4,
	}

	fmt.Printf("n.a %p\n", &n.a) // n.a 0xc0000b8020
	fmt.Printf("n.b %p\n", &n.b) // n.b 0xc0000b8028
	fmt.Printf("n.c %p\n", &n.c) // n.c 0xc0000b8030
	fmt.Printf("n.d %p\n", &n.d) // n.d 0xc0000b8038

}
