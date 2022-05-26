package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统

// 需求：
// 1 添加学员信息
// 2 编辑学员信息
// 3 展示所有学员信息

func showMenu(){
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1 添加学员")
	fmt.Println("2 编辑学员信息")
	fmt.Println("3 展示所有学员信息")
	fmt.Println("4 退出系统")

}

// 获取用户输入的学员信息
func getInput() *Student{
	// 声明 学生结构体字段
	var (
		id int
		name string
		class string
	)

	fmt.Println("请按要求输入学员的信息")
	fmt.Println("请输入学员的学号:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学员的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学员的班级")
	fmt.Scanf("%s\n", &class)

	// 就能拿到用户输入的学员所有信息
	stu := NewStudent(id, name, class)  // 调用 student 的构造函数造一个学生
	return stu  // 将构造好的学生返回出去
}


// 有多个文件 编译的时候都要写上比如 go build main.go student.go 或者直接在项目目录下main.go同级执行go build
// go build main.go student.go  启动项目
func main() {

	sm := NewStudentMar()

	// 循环用户的输入
	for {
		// 1 打印系统菜单
		showMenu()
		// 2 等待用户选择要执行的选项
		var input int
		fmt.Println("请输入你要操作的序号:")
		fmt.Scanf("%d\n", &input)     // Scanf 扫描用户的输入 %d 按数字类型扫入，并且将值  赋给 input 指针变量
		fmt.Println("用户输入的是：", input) // 拿到用户的输入
		// 3 执行用户选择的选项
		switch input {
		case 1:
			// 添加学员
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// 编辑学员
			stu := getInput()
			sm.modify(stu)
		case 3:
			// 展示所有学员
			sm.showStudent()
		case 4:
			// 推出系统
			os.Exit(0)

		}
	}

}
