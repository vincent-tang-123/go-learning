package main

import "fmt"

var sd string = "alex"

func main() {
	fmt.Println("hello world")
	fmt.Println(8+9)
	fmt.Println( sd)

	// 声明 赋值
	var age int = 73
	fmt.Println(age)

	var flag bool = true
	fmt.Println(flag)

	// 先声明 后赋值
	var a string
	a = "qqq"
	fmt.Println(a)

	// 存储数据
	var v1 int = 99
	var v2 int = 33
	var v3 int = v1 + v2
	var v4 int = v1 * v2
	var v5 int = v1 + v2 + v3 + v4
	fmt.Println(v1, v2 , v3, v4, v5 )

	//var name string
	//fmt.Scanf("%s", &name)
	//if name == "alex" {
	//	fmt.Println("用户输入正确")
	//} else {
	//	fmt.Println("用户名输入错误")
	//}
	//

	name := "vincent"
	fmt.Println(name)

	// 因式分解 ，例如声明5个变量，分别有字符串 整形
	var (
		gender = 1
		ages = 18
		salary = 30000
	)
	fmt.Println(gender, ages, salary)
	// 拓展： go 编辑器会认为声明变量不使用就是耍流氓

	// 作用域
	/*
	如果我们定义了 打括号
	不能被上级使用
	可以被同级使用
	可以在子级别使用
	*/
}
