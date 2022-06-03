package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type student struct {
	Name string `json:"name" ini:"s_name"`
	Score int   `json:"score" int:"s_score"`
}

func main() {
	stu1 := student{
		Name: "vincent",
		Score: 90,
	}
	// 通过反射获取结构体中所有字段
	t := reflect.TypeOf(stu1)
	fmt.Printf("name: %v kind : %v\n", t.Name(), t.Kind())  // name: student kind : struct


	for i:=0; i<t.NumField();i++{

	}

}
