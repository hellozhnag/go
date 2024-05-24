package main

import (
	"io"
	"microservice/rpc/gorpc/hellorpc/hanlder"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello~" + request
	return nil
}

func main() {
	_ = rpc.RegisterName(hanlder.HelloServiceName, &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":4321", nil)
	if err != nil {
		return
	}
}
