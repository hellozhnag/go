package main

import (
	"fmt"
	"microservice/rpc/gorpc/hellorpc/hanlder"
	"net/rpc"
)

func main() {
	//这里直接拨号就行了，因为我们使用json协议，就不走gob协议了，使用net
	client, _ := rpc.Dial("tcp", "localhost:4321")

	var reply string

	err := client.Call(hanlder.HelloServiceName+".Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply) //hello~bobby
}
