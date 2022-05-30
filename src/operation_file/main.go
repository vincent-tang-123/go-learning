package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile(){
	fileObj, err:= os.Open("./xxx.txt")
	if err != nil{
		fmt.Printf("OPEN file failed, err %v\n", err)
		return
	}
	defer fileObj.Close() // 关闭文件

	// 循环读取文件，至少文件读完
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // END OF FILE  文件读完会报这个错误
			// 把当前读了多少个字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err: %v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file\n", n) // read 7 bytes from file

		// 读出的内容是 bytes 类型 需要转成字符串
		fmt.Println(string(tmp[:n])) // abc
		                            //  xyz
	}
}

// 文件操作


func readAll(){
	fileObj, err:= os.Open("./xxx.txt")
	if err != nil{
		fmt.Printf("OPEN file failed, err %v\n", err)
		return
	}
	defer fileObj.Close() // 关闭文件

	// 循环读取文件，至少文件读完
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // END OF FILE  文件读完会报这个错误
			// 把当前读了多少个字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err: %v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file\n", n) // read 7 bytes from file

		// 读出的内容是 bytes 类型 需要转成字符串
		fmt.Println(string(tmp[:n])) // abc
		//  xyz
	}
}


// read by bufio
func readByBufio(){
	fileObj, err:= os.Open("./xxx.txt")
	if err != nil{
		fmt.Printf("OPEN file failed, err %v\n", err)
		return
	}
	defer fileObj.Close() // 关闭文件

	reader := bufio.NewReader(fileObj)
	for {
		// 按行来读
		line, err := reader.ReadString('\n') // 注意此处用单引号, 返回的 line 是一个 string
		if err == io.EOF{
			fmt.Println(line)
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio failed, err %v\n", err)
			return
		}
		fmt.Println(line)
	}


}

// ioutil  一次性读完 尽量不要读太大的文件
func readByoutil(){
	content, err := ioutil.ReadFile("./xxx.txt") // 返回值 content 是一个 byte
	if err != nil{
		fmt.Printf("read file by ioutil failed, err:%v\n", err)
		return
	}
	fmt.Println(string(content)) // 专成 string
}


// write
func write(){
	// 得到文件对象                                   os.O_CREATE 文件不存在则创建文件  os.O_APPEND 追加 | os.O_TRUNC 清空
	fileObj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Printf("open file failed %v\n, err", err)
		return
	}
	defer fileObj.Close()

	str := "沙河小王子"
	fileObj.Write([]byte(str)) // []byte
	fileObj.WriteString("hello word") // string

}

func writeByBufio(){
	// 得到文件对象                                   os.O_CREATE 文件不存在则创建文件  os.O_APPEND 追加 | os.O_TRUNC 清空
	fileObj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Printf("open file failed %v\n, err", err)
		return
	}
	defer fileObj.Close()
	write := bufio.NewWriter(fileObj)
	write.WriteString("vincent") // 将内容写入缓冲区
	write.Flush() // 将缓冲区的内容写入磁盘
}

func writeByIoutil(){
	str := "我命由我不由天"
	err := ioutil.WriteFile("./xxx.txt", []byte(str), 0644)
	if err != nil{
		fmt.Printf("write file failed , err:%v\n", err)
		return
	}
}

func str_2_byte(){
	//str := "我命由我不由天"
	//fmt.Println([]byte(str)) // [230 136 145 229 145 189 231 148 177 230 136 145 228 184 141 231 148 177 229 164 169]
	//fmt.Printf("%T\n", []byte(str))  // type []uint8

	var a []int // 声明一个切片
	a = make([]int, 8, 8)  // [0 0 0 0 0 0 0 0] // 初始化一个切片
	fmt.Println(a)
	fmt.Printf("%T\n", a) // []int  // 打印切片的类型

}



func main() {
	//readAll()
	//readByBufio()
	//readByoutil()
	//write()
	//writeByBufio()
	//writeByIoutil()
	str_2_byte()
}
