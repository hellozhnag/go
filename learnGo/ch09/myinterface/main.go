package main

import (
	"fmt"
	"strings"
)

type Duck interface {
	//方法，可以省略func
	Gaga()
	Walk()
	Swimming()
	//只要拥有这3个方法，那就是Duck
}

// 鸭子类型
type pskDuck struct {
	legs int
}

// 3个方法都要实现
func (p *pskDuck) Gaga() {
	fmt.Println("pskDuck GAGA")
}
func (p *pskDuck) Walk() {
	fmt.Println("pskDuck Walk")
}
func (p *pskDuck) Swimming() {
	fmt.Println("pskDuck Swimming")
}

// 多接口实现
type MyWriter interface {
	Writer(string) error
}
type MyCloser interface {
	Closer() error
}

type writerCloser0 struct{}

func (wc *writerCloser0) Writer(s string) error {
	fmt.Println("Writer")
	return nil
}
func (wc *writerCloser0) Closer() error {
	fmt.Println("Closer")
	return nil
}

// 结构体中嵌入interface，常用
type writerCloser struct {
	MyWriter //想放一个写文件/写数据库的实现，到时候实例化时传对应的实例
}

// 写文件的结构体
type fileWriter struct {
	filePath string
}

func (fw *fileWriter) Writer(s string) error {
	fmt.Println("写文件~", s)
	return nil
}

// 写数据库结构体
type dbWriter struct {
	host string
	post string
}

func (dw *dbWriter) Writer(s string) error {
	fmt.Println("写数据库~")
	return nil
}

// 断言
// 实现一个通用的add方法
func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case float64:
		return a.(float64) + b.(float64)
	case string:
		return a.(string) + b.(string)
	default:
		return 0
	}
}

// 接口嵌套
type MyWriter1 interface {
	Write1(string)
}
type MyReader1 interface {
	Read1() string
}

// 继承
type MyReadWriter1 interface {
	MyWriter1
	MyReader1

	ReadWrite1()
}

type SreadWriter struct{}

func (s SreadWriter) Write1(s2 string) {
	//TODO implement me
	fmt.Println("write")
}

func (s SreadWriter) Read1() string {
	//TODO implement me
	fmt.Println("read")
	return ""
}

func (s SreadWriter) ReadWrite1() {
	//TODO implement me
	fmt.Println("read write")
}

func mprint(datas ...interface{}) {
	for _, data := range datas {
		fmt.Println(data)
	}
}

type myinfo struct{}

func (mi myinfo) Error() string {
	return "我不是error"
}

func main() {

	/*
		1、定义接口
	*/
	// 我们定义的方法的接收器为指针类型，所以这里要用&取地址
	var d Duck = &pskDuck{}
	d.Walk() //pskDuck Walk

	//如果一个interface什么都没有，那就可以把任何的struct都赋值给他，即any

	/*
		2、多接口实现
			- 一个接口的多类型实现
			- 一个类型的多接口实现
	*/
	var mv0 MyWriter = &writerCloser0{}
	mv0.Writer("666") //Writer
	var mc0 MyCloser = &writerCloser0{}
	mc0.Closer() //Closer

	//结构体中嵌入interface
	var mv MyWriter = &writerCloser{
		MyWriter: &fileWriter{},
	}
	mv.Writer("123") //写文件~ 123

	var mv1 MyWriter = &writerCloser{
		MyWriter: &dbWriter{},
	}
	mv1.Writer("123") //写数据库~

	/*
		3、通过interface解决动态类型传参
			- 断言
	*/
	fmt.Println(add(1, 2))     //3
	fmt.Println(add(1.2, 2.8)) //4
	re := add("hi ", "tom~")
	fmt.Println(re) //hi tom~
	res, _ := re.(string)
	fmt.Println(strings.Split(res, " ")) //[hi tom~]

	/*
		4、接口嵌套
	*/
	//虽然我们给结构体添加的方法的接收器是普通类型，但是这里仍然可以用&
	//但是如果添加的方法的接收器是指针类型，只能用&
	//一般情况呀，我们采用指针类型的接收器
	//var mrw1 MyReadWriter1 = SreadWriter{}
	var mrw1 MyReadWriter1 = &SreadWriter{}
	mrw1.ReadWrite1()  //read write
	mrw1.Read1()       //read
	mrw1.Write1("555") //write

	/*
		5、空接口
	*/
	//空接口，可装万物
	var data = []interface{}{
		"bobby", 18, 1.80,
	}
	mprint(data) //[bobby 18 1.8]
	mprint(data...)
	//bobby
	//18
	//1.8

	var data1 = []string{
		"bobby1", "bobby2", "bobby3",
	}
	//fmt.Println(data1...)//无法将 'data1' (类型 []string) 用作类型 []any
	//mprint(data1...) //报错，不可以
	mprint(data1) //[bobby1 bobby2 bobby3]
	var datai []interface{}
	for _, d := range data1 {
		datai = append(datai, d)
	}
	mprint(datai...)
	//bobby1
	//bobby2
	//bobby3

	/*
		6、error本质
			- 接口
	*/
	var err error = &myinfo{}
	fmt.Println(err.Error()) //我不是error

}
