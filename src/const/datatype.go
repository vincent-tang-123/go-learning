package main

import (
	"fmt"
	"math"
)

func main() {

	// 十进制打印成二进制
	n := 10
	fmt.Printf("%b\n", n)
	// 十进制
	fmt.Printf("%d\n", n)
	// 八进制
	m := 075
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)

	// 十六进制
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f)


	// uint8
	var age uint8
	age = 255 // 0 ~ 255
	fmt.Println(age)

	// 浮点数
	fmt.Println(math.MaxFloat32) // 32位的浮点数最大值
	fmt.Println(math.MaxFloat64) // 64位的浮点数最大值

	// 布尔值
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	// 字符串 双引号 ""
	s1 := "hello beijing"
	fmt.Println(s1)
	s2 := "你好北京"
	fmt.Println(s2)
	// 打印 windows 平台下的路径 c:\code\go.exe
	fmt.Println("c:\\code\\go.exe")
	fmt.Println("\t制表符\n换行符")
	s3 := `
	多行字符串
	两个反引号之间的内容
	会原样输出
	\n
	\t
`
	fmt.Println(s3)

}
