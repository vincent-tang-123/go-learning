package main

import "fmt"

// 学生结构体

type Student struct {
	id int // 学号唯一
	name string
	class string
}

// newStudent 是 student 类型的构造函数
// 返回指针

func NewStudent(id int, name, class string) *Student{
	return &Student {
		id: id,
		name: name,
		class: class,
	}
}


// 学员管理类型

type StudentMgr struct {
	allStudent []*Student // 使用指针类型切片
}

// newStudentMar 是 studentMgr 的构造函数

func NewStudentMar()*StudentMgr{
	return &StudentMgr{
		allStudent: make([]*Student, 0, 100), // 切片使用前需要初始化
	}
}

// 1,2,3 均是结构体 studentMgr 的方法

// 1 添加学生
func (s *StudentMgr)addStudent(newStu *Student){
	s.allStudent = append(s.allStudent, newStu)
}

// 2 编辑学生
func (s *StudentMgr)modify(newStu *Student){
	for i, v := range s.allStudent{
		if newStu.id == v.id{      // 当学号相同时 就表示找到来需要修改的学生
			s.allStudent[i]=newStu // 根据切片的索引值直接把新学生赋值进来
			return // 操作完成 return 结束
		}
	}
	// 如果走到这里说明输入的学生没有找到
	fmt.Printf("输入的学生信息有误，系统中没有学号是：%d的学生\n", newStu.id)
}


// 3 展示学生
func (s *StudentMgr)showStudent(){
	for _, v := range s.allStudent{
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}