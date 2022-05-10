package main

import "fmt"

// go 语言中的运算符号
func main() {
	// 1 算数运算
	a := 10
	b := 20
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(5 % 2)
	a++ // ++ 在原来的基础上加一  // 同理 --
	fmt.Println(a)

	// 2 关系运算符

	fmt.Println(10>2)
	fmt.Println(10 !=2)

	// 3 逻辑运算符
	fmt.Println(10>5 && 1 == 1)  // and
	fmt.Println(10>5 || 2>1) // or
	fmt.Println(!(2>3)) // not

	// 4 位运算符

	c := 1 // 001
	d := 5 // 101

	fmt.Println(c & d) // 001 => 1
	fmt.Println(c | d) // 101 => 5
	fmt.Println(c ^ d) // 100 => 4

	fmt.Println(1<<2) // 100 => 4
	fmt.Println(4>>2) // 001 => 1
	fmt.Println(1<<10) //  => 1024

	// 5 赋值运算符
	s := 5
	s += 5 // s = s+5
	fmt.Println(s)


}
