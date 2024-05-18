package main

import "fmt"

func main() {
	//map,无序集合，主要是查询方便；数组切片需要遍历，map不需要

	/*
	 1、定义
	*/
	var courseMap = map[string]string{
		"go":   "go攻城狮",
		"grpc": "grpc入门",
		"gin":  "深入理解",
	}
	/*
		2、取值
	*/
	fmt.Println(courseMap["gin"]) //深入理解

	/*
		3、放值
		4、两种初始化方法
		5、map必须初始化才能使用，slice可以不初始化
	*/
	courseMap["mysql"] = "数据库"
	fmt.Println(courseMap) //map[gin:深入理解 go:go攻城狮 grpc:grpc入门 mysql:数据库]

	//map类型要想放值必要要初始化
	//var courseMap1 map[string]string //没初始化，为nil
	//courseMap1["mysql"] = "数据库"      //报错，不能在nil map上添加
	//var courseMap1 map[string]string{}//这样初始化了，就不会报错

	// 使用make初始化，make可以初始化 slice、map、channel
	var courseMap2 = make(map[string]string, 3)
	courseMap2["mysql"] = "数据库"
	fmt.Println(courseMap2) //map[mysql:数据库]

	var m []string
	if m == nil {
		fmt.Printf("yes~\n") //yes~
	}
	m = append(m, "golang")
	fmt.Println(m) //[golang]

	/*
		6、赋值
	*/
	courseMap2["mysql"] = "mysql1"
	courseMap2["db"] = "db1"
	fmt.Println(courseMap2) //map[db:db1 mysql:mysql1]

	/*
		7、遍历
			- 无序，不保证每次打印都是相同顺序
			- map的key类型固定，只能是string、int...
			- value类型随意
	*/
	for k, v := range courseMap2 {
		fmt.Println(k, v)
		//mysql mysql1
		//db db1
	}
	//不需要key值
	for _, v := range courseMap2 {
		fmt.Println(v)
		//db1
		//mysql1
	}
	// 只有一个值，打印的是key！！！
	for key := range courseMap2 {
		fmt.Println(key)
		//mysql
		//db
	}

	/*
		8、获取map中的元素
	*/
	fmt.Println(courseMap2["go"]) //不能通过这样判断go是否在map中,这样返回空，无法判断
	fmt.Println(courseMap2)       //map[db:db1 mysql:mysql1]

	v, ok := courseMap2["go"]
	if ok {
		fmt.Println("in...", v)
	} else {
		fmt.Println("not in...") //not in...
	}
	//也可以这样写
	if v, ok := courseMap2["mysql"]; ok {
		fmt.Println("in...", v) //in... mysql1
	} else {
		fmt.Println("not in...")
	}
	if _, ok := courseMap2["db"]; ok {
		fmt.Println("in...") //in...
	} else {
		fmt.Println("not in...")
	}

	/*
		9、删除元素
	*/
	delete(courseMap2, "db")
	fmt.Println(courseMap2) //map[mysql:mysql1]

	//删除一个不存在的
	delete(courseMap2, "db") //不报错

	/*
		10、提示
			- map线程不安全
			-并发编程，要使用sync.map
	*/
}
