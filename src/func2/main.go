package main

import "fmt"

// 函数进阶之变量作用域
// 函数一等公民，既可以做对象，也可以做参数

// 定义全局变量 num
var num int = 10

// 定义函数
func testGlobal(){
	// 可以在函数中使用变量
	// 1 先在自己的函数中查找，找到了就用自己函数中的
	// 2 函数中找不到变量就在外层找全局变量
	num := 100
	fmt.Println("全局变量", num)
}

func main() {
	//testGlobal()
	//
	//// 变量i此时只在for 循环语句块中生效
	//for i:=0;i<5;i++{
	//	fmt.Println(i)
	//}
	////fmt.Println(i)  外层访问不到内部for语句块中的变量

	// 函数可以作为变量
	//abc := testGlobal
	//fmt.Printf("%T\n", abc)
	//abc()

	r1 := calc(100, 200, add)
	r2 := calc(100, 200, sub)
	fmt.Println(r1)
	fmt.Println(r2)
}

func sub(x,y int)int{
	return x -y
}

func add (x, y int) int{
	return x + y
}

// op 为函数类型 func(参数) 返回值
func calc(x, y int, op func(int, int) int)int{
	return op(x, y)
}