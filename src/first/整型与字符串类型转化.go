package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整形转换成字符串类型
	//v1 := 19
	//// Itoa 中只能接收 int 类型
	//result := strconv.Itoa(v1)
	//fmt.Println(result, reflect.TypeOf(result))

	// 字符串转换成整型
	v2 := "33"
	result, err := strconv.Atoi(v2)
	//fmt.Println(result, err)
	if err != nil {
		fmt.Println("转换失败")
	} else {
		fmt.Println("转换成功", result)
	}

}
