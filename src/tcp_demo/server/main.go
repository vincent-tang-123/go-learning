package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server top


func process(conn net.Conn){
	defer conn.Close() // 处理完之后要关闭这个连接
	// 针对当前的连接做数据的发送和接收操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil{
			fmt.Printf("read from conn failed, err:%v\n", err)
		}
		recv := string(buf[:n])
		fmt.Printf("接收到的数据：%v\n", recv)
		conn.Write([]byte("ok")) // 给客户端返回 ok
	}
}

func main() {
	// 1 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil{
		fmt.Printf("listen failed, err:%v\n", err)
		return // 监听失败 return 退出函数
	}
	for {
		// 2 等待客户端来建立连接
		conn, err := listen.Accept()
		if err != nil{
			fmt.Printf("accept failed, err:%v\n", err)
			continue // 连接失败 continue 继续等待下一次连接
		}

		// 3 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}


}
