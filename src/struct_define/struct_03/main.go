package main

import "fmt"

// 结构体的继承  （用结构体的嵌套来模拟继承，注意嵌套的必须是匿名字段）
// 要注意一点，Dog结构体中，Animal必须是匿名字段，dog才会继承animal的方法。否则不会

type Animal struct {
	name string
}

func (a *Animal) move(){
	fmt.Printf("%s 会动\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal // 匿名嵌套 而且嵌套的是一个结构体指针  必须是匿名结构体
}

func (d *Dog)wang(){
	fmt.Printf("%s 会汪汪\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}

	d1.wang() // 乐乐 会汪汪

	d1.move() // 乐乐 会动

}
