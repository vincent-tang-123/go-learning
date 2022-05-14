package main

import "fmt"

func main() {
	// 数组类型 数组长度和类型
	//var a [3]int
	//var b [4]int
	//
	//fmt.Println(a)
	//fmt.Println(b)

	// 数组的初始化
	// 1 定义时 使用初始值列表的方式初始化
	//var cityArray = [4]string{"北京", "shanghai", "guangzhou", "shengzhen"}
	//fmt.Println(cityArray)
	//fmt.Println(cityArray[0])
	//fmt.Println(cityArray[3])
	//// 2 编译器推导数组的长度  ... 代表可变长
	//var boolArray = [...]bool{true, false, true, false}
	//fmt.Println(boolArray)
	//// 3 使用索引值方式初始化
	//var langArray = [...]string{1:"golang", 3:"python", 7:"java"}
	//fmt.Println(langArray)
	//fmt.Printf("%T\n",langArray)

	// 数组的遍历

	//var cityArray = [4]string{"北京", "shanghai", "guangzhou", "shengzhen"}
	// 1 for 循环遍历
	//for i:=0; i<len(cityArray); i++{
	//	fmt.Println(cityArray[i])
	//}

	// 2 for range 遍历
	// 带索引
	//for index, value := range cityArray{
	//	fmt.Println(index, value)
	//}
	//// 只要值 不要索引
	//for _, value := range cityArray{
	//	fmt.Println( value)
	//}

	// 二维数组
	// 多维数组只有最外层可以使用 ... 这样的简写， 内层必须写上数字
	//cityArray := [...][2]string{
	//	{"beijing", "xian"},
	//	{"shanghai", "chongqing"},
	//	{"hangzhou", "shenzhen"},
	//}
	//fmt.Println(cityArray) // [[beijing xian] [shanghai chongqing] [hangzhou shenzhen]]
	//
	//fmt.Println(cityArray[1][1]) // 二维数组取值
	//
	//// 二维数组的遍历  两层循环
	//for _, v1 := range cityArray{
	//	for _, v2 := range v1{
	//		fmt.Println(v2)
	//	}
	//}

	// 数组是值类型
	x := [3][2]int{
		{1,2}}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	y := x
	y[0][0] = 1000
	fmt.Println(y)


}

func f1(a [3][2]int){
	a[0][1]=100
}
