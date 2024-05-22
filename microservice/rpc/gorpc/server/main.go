package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello~" + request
	return nil
}

func main() {
	//go内置rpc

	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		return
	}
	listener, _ := net.Listen("tcp", ":4321")
	conn, err := listener.Accept() //当一个新的连接进来后，
	rpc.ServeConn(conn)
	/*
		1、实例化server
			- net.Listen("tcp", ":1234")
		2、注册处理逻辑
			- rpc.RegisterName("HelloService", HelloService{})
		3、启动服务
			- rpc.ServeConn(conn)
	*/
}
