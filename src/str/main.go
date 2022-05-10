package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串常见操作

	// len 统计的是字节
	s := "hello"
	fmt.Println(len(s)) // 5

	s1 := "hello你好"
	fmt.Println(len(s1)) // 11

	// 拼接字符串
	fmt.Println(s+s1)
	s3 := fmt.Sprintf("%s - %s", s, s1)
	fmt.Println(s3)

	// 字符串的分割
	s4 := "how do you do"
	fmt.Println(strings.Split(s4, " ")) // 按空格切割
	fmt.Printf("%T\n", strings.Split(s4, " ")) // %T 得到切片数据类型 []string

	// 判断是否包含
	fmt.Println(strings.Contains(s4, "do"))
	// 判断前缀
	fmt.Println(strings.HasPrefix(s4, "how"))
	// 判断后缀
	fmt.Println(strings.HasSuffix(s4, "how"))
	// 判断字串的位置
	fmt.Println(strings.Index(s4, "do"))
	// 判断子串出现的位置
	fmt.Println(strings.LastIndex(s4, "do"))
	// join
	s5 := []string{"how", "do", "u", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5 ,"+"))
}
