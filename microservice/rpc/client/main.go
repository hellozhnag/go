package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/httpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

// 对远程调用封装
func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get("http://127.0.0.1:8001/add?a=1&b=2")
	fmt.Println(res)
	body, _ := res.Body()
	fmt.Println(string(body))
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	fmt.Println(Add(1, 2))
	//打印：
	//&{1 http://127.0.0.1:8001/add?a=1&b=2 0xc000180000 []}
	//{"data":3}
	//3
}
