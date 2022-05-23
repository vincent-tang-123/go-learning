package main

import "fmt"

// 方法的定义示例

// Person 是一个结构体 首字母大写 对外暴露
type Person struct {
	name string
	age int
}

// 构造函数

func NewPerson(name string, age int) *Person{
	return &Person{
		name: name,
		age: age,
	}
}


// 定义方法

// Dream 是为 Person 类型定义的方法
func (p Person)Dream(){
	fmt.Printf("%s 的梦想是学好go语言\n", p.name)
}

// SetAge 是一个修改年龄的方法
// 指针接受者指的就是接受者的类型是指针
func (p *Person)SetAge(newAge int){
	p.age = newAge
}

// SetAge2 是一个使用值接受者来修改年龄的方法
func (p Person)SetAge2(newAge int){
	p.age = newAge // 此时的 p 已经是原结构体的一种拷贝 所以无法修改原结构体
}

func main() {
	p1 := NewPerson("vincent", 18)
	//(*p1).Dream() // vincent 的梦想是学好go语言
	p1.Dream() // vincent 的梦想是学好go语言


	// 调用修改年龄的方法
	fmt.Println(p1.age) // 修改前 18
	p1.SetAge(20)
	fmt.Println(p1.age) // 修改后 20


	// 用方法修改结构体字段值 一定要用指针接受者

	fmt.Println(p1.age) // 修改前 20
	p1.SetAge2(30)
	fmt.Println(p1.age) // 修改后 20


}
