package main

import "fmt"

// 嵌套结构体的字段冲突

type Address struct {
	Province string
	City 	 string
	UpdateTime string
}

type Email struct {
	Addr string
	UpdateTime string
}


type Person struct {
	Name string
	Gender string
	Age int
	Address // 嵌套另外一个结构体
	Email // 嵌套另一个结构体
}

func main() {
	 p1 := Person{
		 Name: "vincent",
		 Gender: "nan",
		 Age: 18,
		 Address: Address{
			 Province: "shandong",
			 City: "qingdao",
			 UpdateTime: "2022-02-17",
		 },
		 Email: Email{
			 Addr: "xiaowangzi@163.com",
			 UpdateTime: "2022-02-18",
		 },
	 }
	 fmt.Printf("%#v\n", p1) // main.Person{Name:"vincent", Gender:"nan", Age:18, Address:main.Address{Province:"shandong", City:"qingdao", UpdateTime:"2022-02-17"}, Email:main.Email{Addr:"xiaowangzi@163.com", UpdateTime:"2022-02-18"}}%
	 //fmt.Println(p1.UpdateTime)  报错 嵌套结构体中包含多个同名的 UpdateTime 字段
	 fmt.Println(p1.Address.UpdateTime) // 2022-02-17
	 fmt.Println(p1.Email.UpdateTime) // 2022-02-18
}
