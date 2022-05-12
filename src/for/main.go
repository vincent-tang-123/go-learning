package main

import "fmt"

// for 循环
func main() {
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	// 2 省略初识语句，但是必须保留初识语句后面的分号
	var i = 0
	for ; i<10 ; i++{
		fmt.Println(i)
	}

	// 3 省略初识语句和结束语句
	var a = 10
	for a > 0{
		fmt.Println(a)
		a--
	}

	// 4 无限循环
	//for {
	//	fmt.Println("hello beijing")
	//}

	// 5 break 跳出 for 循环

	//for i:=0; i<5; i++{
	//	fmt.Println(i)
	//	if i == 3{
	//		break
	//	}
	//}

	// continue 继续下一次循环
	for i:=0; i<5; i++{
		if i == 3{
			continue // 跳过本次循环
		}
		fmt.Println(i)

	}

}
