package main

import "fmt"

func main() {
	finger := 13
	//if finger == 1{
	//	fmt.Println(111)
	//} else if finger == 3{
	//	fmt.Println(2222)
	//}

	switch finger {
	case 1:
		fmt.Println(11111)
	case 2:
		fmt.Println(3333333)
	default:
		fmt.Println(finger)
	}

	// case 一次判断多个值
	num := 5
	switch num {
	case 1,3,5,7,9:
		fmt.Println("奇数")
	case 2,4,6,8:
		fmt.Println("偶数")
	}


	// case 中使用表达式
	age := 30
	switch  {
	case age>=18:
		fmt.Println("澳门首家线上赌场开业啦")
	case age<18:
		fmt.Println("warning")
	default:
		fmt.Println("嘿嘿嘿")
	}
}


