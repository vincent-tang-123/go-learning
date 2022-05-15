package main

import (
	"fmt"
	"sort"
)

// 开发中 定义切片 应适当想好切片的容量大小 避免动态扩容造成内存分配开销

func main() {
	// 切片
	//var a []string
	//var b []int
	//var c = []bool{true, false}
	//
	//fmt.Println(a, b, c)

	// 1 基于数组得到切片
	//a := [5] int {55,56,57, 58, 59}
	//b := a[1:4]
	//fmt.Println(b)
	//fmt.Printf("%T\n", b) // 切片类型
	//
	//// 2 切片再次切片
	//c := b[:] // b[0:len(b)]
	//fmt.Println(c)
	//fmt.Printf("%T\n", c)
	//
	//// 3 make 函数构造切片
	//d := make([]int,5, 10) // 5是切片的长度，10 是切片容量，也可以不写，此时切片的长度和容量都是5
	//fmt.Println(d)
	//fmt.Printf("%T\n", d)
	//// 通过 len() 函数获取切片的长度
	//fmt.Println(len(d))
	//// 通过 cap() 函数获取切片长度
	//fmt.Println(cap(d))

	// nil
	//var a [] int // 声明 int 类型切片
	//var b = [] int {} // 声明并且初始化
	//c := make([]int, 0)
	//if a == nil {
	//	fmt.Println("a 是一个 nil")
	//}
	//fmt.Println(a, len(a), cap(a))
	//if b == nil {
	//	fmt.Println("b 是一个 nil")
	//}
	//fmt.Println(b, len(b), cap(b))
	//if c == nil {
	//	fmt.Println("c 是一个nil")
	//}
	//fmt.Println(c, len(c), cap(c))

	// 判断空切片 用 len 判断


	// 切片的赋值拷贝
	//a := make([]int,5, 10) // [0,0,0,0,0]
	//b := a // a 和 b 公用一个底层数组
	//b[0] = 100
	//fmt.Println(a) // [100 0 0 0 0]
	//fmt.Println(b) // [100 0 0 0 0]



	// 切片的遍历
	//a := []int{1,2,3,4,5}
	//
	//// 根据索引来遍历
	//for i:=0;i<len(a);i++{
	//	fmt.Println(i, a[i])
	//}
	//
	//// for range 遍历
	//for index, value := range a{
	//	fmt.Println(index, value)
	//}


	// 切片的扩容

	//var a []int // 此时它并没有申请内存
	// 切片的初始化之后才能使用
	//a = append(a, 10) // 必须要赋值，可能返回原来的切片又可能返回扩容后的切片
	//fmt.Println(a)

	//for i:=0;i<10;i++{
	//	a = append(a, i)
	//	fmt.Printf("%v len: %d cap:%d ptr:%p \n", a, len(a), cap(a), a)
	//	/*
	//	[0] len: 1 cap:1 ptr:0xc0000b2008
	//	[0 1] len: 2 cap:2 ptr:0xc0000b2030
	//	[0 1 2] len: 3 cap:4 ptr:0xc0000b6020
	//	[0 1 2 3] len: 4 cap:4 ptr:0xc0000b6020
	//	[0 1 2 3 4] len: 5 cap:8 ptr:0xc0000b8040
	//	[0 1 2 3 4 5] len: 6 cap:8 ptr:0xc0000b8040
	//	[0 1 2 3 4 5 6] len: 7 cap:8 ptr:0xc0000b8040
	//	[0 1 2 3 4 5 6 7] len: 8 cap:8 ptr:0xc0000b8040
	//	[0 1 2 3 4 5 6 7 8] len: 9 cap:16 ptr:0xc0000ba000
	//	[0 1 2 3 4 5 6 7 8 9] len: 10 cap:16 ptr:0xc0000ba000
	//	*/
	//}

	// 一次往切片中追加多个元素
	//a = append(a, 1,2,3,4,5,6)
	//fmt.Println(a)

	// 追加切片
	//b := []int{1,2,3,4,5}
	//a = append(a, b...) // ... 意思是将切片打散 成单个元素追加
	//fmt.Println(a)


	// 切片的拷贝
	//a := []int{1,2,3,4,5}
	//b := make([]int, 5)
	//copy(b, a)
	//b[0] = 100
	//// copy 后 b 和 a 没有指向同一块地址
	//fmt.Println(a) // [1 2 3 4 5]
	//fmt.Println(b) // [100 2 3 4 5]

	// 切片的删除
	// 从 a 中删除 索引为 index 的元素， append(a[:index], a[index+1]...)
	//a := []string{"北京", "上海", "广州", "深圳"}
	//a = append(a[0:2],a[3:]...)  // 删除元素广州
	//fmt.Println(a)


	// 使用内置的sort包对数组 var a = [...]int{3,7,8,9,1} 进行排序
	var a = [...]int{3,7,8,9,1}
	// a[:] 得到一个从头切到尾的切片，指向了底层的数组 a
	sort.Ints(a[:])
	fmt.Println(a)

}
