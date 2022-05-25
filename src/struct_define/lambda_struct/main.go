package main

import "fmt"

// 结构体的匿名字段

// 使用匿名结构体 相同类型的字段必须唯一

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"vincent",
		18,
	}
	fmt.Println(p1) // {vincent 18}

	fmt.Println(p1.string, p1.int) // vincent 18


}
