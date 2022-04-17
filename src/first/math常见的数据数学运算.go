package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(-19))  // 取绝对值
	fmt.Println(math.Floor(3.14)) // 向下取整
	fmt.Println(math.Ceil(3.14)) //
	fmt.Println(math.Round(3.3478))
	fmt.Println(math.Round((3.3478*100)/100))
	fmt.Println(math.Mod(11, 3))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Pow10(2))
	fmt.Println(math.Max(1,2))
	fmt.Println(math.Min(1,2))

}
