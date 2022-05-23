package main

import (
	"fmt"
	"unsafe"
)

// go 内存对齐

type Args struct {
	num1 int
	num2 int
}

type Flag struct {
	num1 int16
	num2 int32
}


func main() {

	// link: https://geektutu.com/post/hpg-struct-alignment.html

	// 打印内存占用
	fmt.Println(unsafe.Sizeof(Args{})) // 16
	fmt.Println(unsafe.Sizeof(Flag{})) // 8

	/*
	Args 由 2 个 int 类型的字段构成，在 64位机器上，一个 int 占 8 字节，因此存储一个 Args 实例需要 16 字节。
	Flag 由一个 int32 和 一个 int16 的字段构成，成员变量占据的字节数为 4+2 = 6，但是 unsafe.Sizeof 返回的结果为 8 字节，多出来的 2 字节是内存对齐的结果。
	因此，一个结构体实例所占据的空间等于各字段占据空间之和，再加上内存对齐的空间大小。
	*/


}
