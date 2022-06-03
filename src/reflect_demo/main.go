package main

import (
	"fmt"
	"reflect"
)

// reflect demo

func reflectType(x interface{}){
	// 我是不知道别人调用我这个函数的时候会传进来什么类型的变量
	// 1 方式1： 通过类型断言
	// 2 方式2： 借助反射

	obj := reflect.TypeOf(x)
	fmt.Println(obj)
	fmt.Println(obj, obj.Name(), obj.Kind())  // main.Dog Dog struct  []string  slice
	fmt.Printf("%T\n", obj)  // *reflect.rtype
}

func reflectValue(x interface{}){
	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Printf("%v %T\n", v, v)

	k := v.Kind() // 拿到值对应的类型种类
	fmt.Println(k)  // int32

	// 如何拿到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)  // 100 int32


	}
}

func reflectSetValue(x interface{}){
	v := reflect.ValueOf(x)

	// Elem() 用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:  // 如果是 int32 类型就将值设置成100
		v.Elem().SetInt(100)
	case reflect.Float32:  // 如果是 float32 类型就将值设置成 3。21
		v.Elem().SetFloat(3.21)

	}
}


type Cat struct {

}

type Dog struct {

}


func main() {
	//var a float32 = 1.234
	//reflectType(a)  // float32
	//
	//var b int8 = 10
	//reflectType(b)  // int8
	//
	//var c Cat
	//var d Dog
	//
	//reflectType(c)  // main.Cat
	//reflectType(d)  // main.Dog
	//
	//// slice
	//var e []int
	//var f []string
	//reflectType(e)
	//reflectType(f)


	// valueOf
	var aa int32 = 100
	reflectValue(aa) // 100 reflect.Value

	var bb float32 = 1.234
	reflectValue(bb)  // 1.234 float32

	// set value
	var aaa int32 = 18
	reflectSetValue(&aaa)  // 修改函数中的值 一定传的是地址 而不是传值 aaa
	fmt.Println(aaa)  // 100


}
