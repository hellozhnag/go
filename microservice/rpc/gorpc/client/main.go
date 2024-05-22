package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1、建立连接
	client, _ := rpc.Dial("tcp", "localhost:4321")
	//var reply *string = new(string)
	//err := client.Call("HelloService.Hello", "bobby", reply)
	//或者这样写
	var reply string
	err := client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	//fmt.Println(*reply) //hello~bobby
	fmt.Println(reply) //hello~bobby
}
