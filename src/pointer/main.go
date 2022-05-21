package main

import "fmt"

// 指针  即存放内存地址的类型

// 指针取值 是根据内存地址取值
//
//func main() {
//	a := 10 // int 类型
//	b := &a // *int 类型 （int 指针）
//	fmt.Printf("%v %p\n", a, &a) // 10 0xc0000b2008
//	fmt.Printf("%v %p\n", b, b) // 0xc0000b2008 0xc0000b2008
//	// 变量 b 是一个 int 类型的指针（*int）， 它存储的是变量 a 的内存地址
//
//	c := *b // 根据内存地址去取值
//	fmt.Println(c)
//}

func modify1(x int){
	x = 100
}

func modify2(y *int){
	*y = 100
}

func main(){
	a := 1
	modify1(a)
	fmt.Println(a) // 1

	modify2(&a)
	fmt.Println(a) // 100
}
