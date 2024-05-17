package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	name1 := "\tgo\"课程\""
	fmt.Println(name1) //	go"课程"

	name2 := `go
			 "课程"`
	fmt.Println(name2) //ln默认加了\r\n

	//格式化输出
	uname := "tom"
	out := "hello " + uname
	fmt.Println(out)
	age := 18
	address := "china"
	mobile := "15936178333"
	fmt.Println("用户名:"+uname, "年龄:"+strconv.Itoa(age), "地址:"+address, "电话:"+mobile) //这个性能好
	fmt.Printf("用户名:%s,年龄:%d,地址:%s,电话:%s \r\n", uname, age, address, mobile)        //这个好维护
	userMsg := fmt.Sprintf("用户名:%s,年龄:%d,地址:%s,电话:%s \r\n", uname, age, address, mobile)
	fmt.Println(userMsg)

	//通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名")
	builder.WriteString(uname)
	builder.WriteString("用户名")
	builder.WriteString(strconv.Itoa(age))
	re := builder.String()
	fmt.Println(re)

}
