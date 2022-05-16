package main

import (
	"fmt"
	"strings"
)

// map 映射
func main() {
	// 只声明map 类型 但是没有初始化， a就是初始值 nil
	//var a map[string]int
	//fmt.Println(a == nil)
	//
	//// map 的初始化
	//a = make(map[string]int, 8) // 8 代表容量可不写
	//fmt.Println(a == nil)
	//
	//// map 中添加键值对
	//a["beijing"] = 100
	//a["shenzhen"] = 200
	//fmt.Println(a)
	//fmt.Printf("a:%#v\n", a) // a:map[string]int{"beijing":100, "shenzhen":200}
	//fmt.Printf("%T\n", a)
	//
	//// 声明 map 的同时完成初始化
	//b := map[int]bool{
	//	1:true,
	//	2:false,
	//}
	//fmt.Printf("b:%#v\n", b)
	//fmt.Printf("%T\n", b)
	//
	//var c map[int]int
	////c[100] = 200  // panic: assignment to entry in nil map
	//// c 这个map没有初始化不能直接进行操作
	//
	//// 对 c 进行初始化
	//c = make(map[int]int, 8)
	//// 赋值
	//c[100] = 200
	//fmt.Println(c)

	// 判断某个键存不存在
	var scoremap = make(map[string]int, 8)
	scoremap["沙河"] = 100
	scoremap["沙河小王子"] = 200
	// 判断 key 张二狗 在不在 map 中
	//value, ok := scoremap["张二狗"]
	//fmt.Println(value, ok)
	//if ok {
	//	fmt.Println("在里面", value)
	//}else{
	//	fmt.Println("未找到")
	//}

	// 遍历 map for range
	// map 是无序的，键值对和添加的顺序无关
	//for k, v := range scoremap{
	//	fmt.Println(k, v)
	//}
	//
	//// 只遍历 map 中的 key
	//for k := range scoremap{
	//	fmt.Println(k)
	//}
	//
	//// 只遍历 map 中的 value
	//for _, v := range scoremap{
	//	fmt.Println(v)
	//}

	// 删除键值对
	//delete(scoremap, "沙河")
	//fmt.Println(scoremap)

	// 按照固定顺序遍历 map
	//var score = make(map[string]int, 100)
	//// 添加 50 个键值对
	//for i:=0;i<50;i++{
	//	key := fmt.Sprintf("stu%02d", i)
	//	value := rand.Intn(100)
	//	score[key] = value
	//}
	//for k, v := range score{
	//	fmt.Println(k, v)
	//}
	//
	//// 按照 key 从小到大的顺序 遍历 map
	//// 1 先取出所有 key 存放到切片中
	//keys := make([]string, 50, 100)
	//for k := range score{
	//	keys = append(keys, k)
	//}
	//// 2 对 key 做排序
	//sort.Strings(keys) // 目前是一个有序的切片
	//// 3 排序后的 key 对 score 排序
	//for _, key := range keys{
	//	fmt.Println(key, score[key])
	//}




	// 元素类型为 map 的切片

	//var mapSlice = make([]map[string]int, 8) // 只是完成了切片的初始化
	//// [nil, nil, nil, nil, nil, nil, nil, nil]
	//fmt.Println(mapSlice[0] == nil)
	//// 还需要完成内部 map 元素的初始化
	//mapSlice[0] = make(map[string]int, 8) // 完成 map 的初始化
	//mapSlice[0]["北京"] = 100
	//fmt.Println(mapSlice) // [map[北京:100] map[] map[] map[] map[] map[] map[] map[]]



	// 值为切片的 map
	//var sliceMap = make(map[string][]int, 8) // 只完成了 map 的初始化
	//v, ok := sliceMap["中国"]
	//if ok{
	//	fmt.Println(sliceMap["中国"], v)
	//}else{
	//	// sliceMap 中没有中国这个键
	//	sliceMap["中国"] = make([]int, 3, 8) // 完成了对切片的初始化
	//	sliceMap["中国"][0] = 100
	//	sliceMap["中国"][1] = 100
	//	sliceMap["中国"][2] = 100
	//}
	//
	//// 遍历 sliceMap
	//for k, v := range sliceMap{
	//	fmt.Println(k, v)
	//}


	// 统计一个字符串中每个单词出现的次数
	// "how do you do" 中每个单词出现的次数
	// 0 定义一个map[string]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	// 1 字符串中都有那些单词
	words := strings.Split(s, " ")
	// 2 遍历单词做统计
	for _, word := range words{
		v, ok := wordCount[word]
		if ok{
			// map 中 有这个单词的统计记录
			wordCount[word] = v +1
		}else {
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount{
		fmt.Println(k, v)
	}

}
