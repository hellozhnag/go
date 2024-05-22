package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}
type Employee struct {
	Name    string
	company Company
}
type PrintResult struct {
	Info string
	Err  error
}

// 远程的
func RpcPrintln(employee Employee) {
	/*
		客户端：
			1、建立连接，  tcp/http/http2.0
			2、将employee序列化为JSON字符串 --- 序列化
			3、发送JSON字符串---调用成功后实际接受的是一个二进制数据
			4、等待服务器发送结果
			5、将服务器返回的数据解析成PrintResult对象 --- 反序列化
		服务端
			1、监听网络端口
			2、读取数据 --- 二进制的JSON数据
			3、对数据进行反序列化Employee对象
			4、开始处理业务逻辑
			5、将处理后的结果PrintResult对象序列化
			6、将序列化数据网络返回
	*/
}

func main() {
	/*
		将Add函数变成一个远程调用函数
			1、数据编码协议：JSON、XML、protobuf、msgpack
				- json并不是一个高性能编码协议
				- 本地对象先序列化为json，发给另一台服务器；另一台服务器接收后再反序列化
			2、Call id
			3、序列化、反序列化
			4、网络传输协议
	*/
	fmt.Println(Add(1, 2))
	// 如果调用的是远程的打印
	fmt.Println(Employee{
		Name: "zhang3",
		company: Company{
			Name:    "tencent",
			Address: "hongkong",
		},
	})
}
