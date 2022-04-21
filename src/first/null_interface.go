package main

import "fmt"

// 空接口
// 接口中没有定义任何方法时，该接口就是一个空接口
// 任意类型都实现了空接口 --》 空接口可以存储任意类型的值

// 空接口一般不需要提前定义
type xxx interface {

}


// 空接口的应该
// 1。空接口类型作为函数的参数
// 2 空接口类型可以作为map 的 value类型



func main() {

	var x interface{} // 定义一个空接口类型

	x = "hello"
	x = 100
	x = false
	//fmt.Println(x) // 为什么 Println 作为函数 不管传什么类型 都能打印出来了？ 空接口

	//var m = make(map[string]interface{}, 16) // 定义了一个对象，key类型 string, 值类型为任意
	//m["name"] = "vincent"
	//m["age"] = 18
	//m["hobby"] = []string{"篮球", "足球", "双色球"}
	//fmt.Println(m)
	ret, ok := x.(string) // 类型断言,猜的类型不对时 ok=false, ret=string 类型的零值，会返回一个布尔值，
	if !ok{
		fmt.Println("不是字符串类型")
	}else{
	fmt.Println("是字符串类型", ret)}
	

	// 使用switch 进行类型断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型， value:%v\n", v)
	case bool:
		fmt.Printf("是布尔类型, value:%v\n", v)
	case int:
		fmt.Printf("是int类型， value：%v\n", v)
	default:
		fmt.Printf("猜不到了，value:%v\n", v)
	}

}
