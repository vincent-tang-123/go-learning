package main

import "fmt"

func main() {
	// 表达式
	//switch 1 + 1 {
	//case 1:
	//	fmt.Println("==1")
	//case 2:
	//	fmt.Println("==2")
	//case 3:
	//	fmt.Println("==3")
	//default:
	//	fmt.Println("都不满足")
	//}

	// 变量
	var age int
	fmt.Scanln(&age)
	switch age {
	case 2:
		fmt.Println("==1")
	case 5:
		fmt.Println("==5")
	case 7:
		fmt.Println("==3")
	default:
		fmt.Println("都不满足")
	}
}
