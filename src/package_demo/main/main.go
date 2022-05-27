package main


// Go 语言中不允许导入包不使用
// GO 语言中不允许循环导入

// 多行导入
import (
	"fmt"
	// 给导入的包起别名 test_1 写在包前面中间有空格隔开 即可
	test_1 "vincent441/src/package_demo/calc" // 导入包 需要将包放入 go.mod 文件中
	// 如项目没有 go.mod 文件
	// 需要 go mod init <module name>  eg: go mod init vincent441
	// 自己本地新创建的包 需要被 其他包导入 ，需要执行  `go mod tidy` 命令  将包添加到 go.mod 文件中 管理起来
)

func main() {
	a := test_1.Add(10, 20) // 导包成功 注意 如果想另外一个包中函数可见，该函数名需大写
	fmt.Println(a) // 30
	fmt.Println(test_1.Name) // vincent
}

// init 多用作一些初始化的工作
func init()  {
	fmt.Println("main init")
}