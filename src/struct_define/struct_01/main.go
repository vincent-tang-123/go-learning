package main

import "fmt"

// 嵌套结构体

type Address struct {

	Province string
	City string
}

type Person struct {
	Name string
	Gender string
	Age int
//	Province string
//	City string
//	Address Address // 嵌套结构体定义
	Address // 嵌套另外一个结构体
}

func main() {
	p1 := Person{
		Name: "vincent",
		Gender: "nan",
		Age: 18,
		Address: Address{ // 嵌套结构体赋值
			Province: "shandong",
			City: "qingdao",
		},

	}

	fmt.Printf("%#v\n", p1) // main.Person{Name:"vincent", Gender:"nan", Age:18, Address:main.Address{Province:"shandong", City:"qingdao"}}

	fmt.Println(p1.Name, p1.Age, p1.Gender, p1.Address) // vincent 18 nan {shandong qingdao}
	fmt.Println(p1.Name, p1.Age, p1.Gender, p1.Address.Province, p1.Address.City) // vincent 18 nan shandong qingdao


	// 通过嵌套的匿名结构体字段访问其内部的字段
	fmt.Println(p1.Address.Province) // shandong
	// 直接访问匿名结构体中的字段
	fmt.Println(p1.Province) // shandong




}
