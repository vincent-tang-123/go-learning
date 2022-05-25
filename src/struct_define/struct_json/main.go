package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见行与JSON序列化

// 如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的
// 如果一个结构体中的字段名是大写的 ，那么该字段就是对外可见的

type student struct {
	ID int
	Name string
}

// student 的构造函数
func newStudent(id int, name string) student{
	return student{
		ID: id,
		Name: name,
	}
}


type class struct {               // 加上 tag 后变小写。方便前后端传值 这里的 json 是使用 json 包解析是 对应解析后的字段
	Title string `json:"title"`  // {"title":"火箭101","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},{"ID":4,"Name":"stu04"},{"ID":5,"Name":"stu05,{"ID":6,"Name":"stu06"},{"ID":7,"Name":"stu07"},{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"}]}
	// 意思是，在 json 包中字段解析成 student_list 在数据库 db 包中字段解析成 student，包支持多个中间用一个空格隔开，值必须用双引号
	Students []student  `json:"student_list" db:"student" xml:"ss"`   // 切片类型字段

}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title: "火箭101",
		Students: make([]student,0,20), // 切片使用时候  必须用 make 初始化
	}
	// 往班级c1 中添加学生
	for i := 0;i<10;i++{
		// 构造 10个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)

	// JSON 的序列化：Go语言中的数据 -> JSON 字符串
	// 注意 结构体中的字段首字母是小写的化，只能在当前包中显示 。小写时 使用json 包访问不到 该字段的值
	data, err := json.Marshal(c1)
	if err != nil{
		fmt.Println("json marshal failed, err", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data) // 结构体 title 小写时对外访问不到 // main.class{title:"", Students:[]main.student{main.student{ID:0, Name:"stu00"}, main.student{ID:1, Name:"stu01"}}}

	// JSON 反序列化: JSON 格式的字符串 -> Go语言中的数据
	jsonStr := `{"Title":"火箭101","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil{
		fmt.Println("json unmarshal fail, err", err)
		return
	}
	fmt.Printf("%#v\n", c2) // main.class{Title:"火箭101", Students:[]main.student{main.student{ID:0, Name:"stu00"}, main.student{ID:1, Name:"stu01"}}}

}
