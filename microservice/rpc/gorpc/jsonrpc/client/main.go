package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//这里直接拨号就行了，因为我们使用json协议，就不走gob协议了，使用net
	conn, _ := net.Dial("tcp", "localhost:4321")

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err := client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply) //hello~bobby
}
