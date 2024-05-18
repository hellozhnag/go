package main

import (
	"fmt"
	"strconv"
)

// 基于基本类型自定义类型
type MyInt2 int

// 自定义一个类型，并扩展方法
func (mi MyInt2) string() string {
	return strconv.Itoa(int(mi))
}

// 想要放多个person的信息到一个集合中
var persons [][]string

// 定义结构体
type Person struct {
	name   string
	age    int
	addr   string
	height float32
}

// 结构体嵌套
//   - 第一种嵌套方式
type Student struct {
	p     Person
	score float32
}

// - 第二种嵌套-匿名嵌套
type Student1 struct {
	Person
	score float32
}

// - 匿名嵌套---覆盖效果
type Student2 struct {
	Person
	name string
}

// 在Person结构体上定义方法
// func 接收器 方法名() (返回值类型){}
// 接收器有两种形态
//   - 第一种，值传递
func (p Person) print() {
	p.name = "111"
	fmt.Printf("name:%s,age:%d,addr:%s,height:%f\n", p.name, p.age, p.addr, p.height)
}

// - 第二种，指针,数据较大时，考虑指针传递
func (p *Person) print1() {
	p.name = "111"
	//这里我们无需将结构体值从指针中取出来，而是直接使用指针获取name、age...
	// 因为go语言在结构体指针这做了优化
	fmt.Printf("name:%s,age:%d,addr:%s,height:%f\n", p.name, p.age, p.addr, p.height)
}

func main() {

	/*
		1、type关键字
			- 定义结构体
			- 定义接口
			- 定义类型别名
				- 别名是为了更好地理解代码
			- 类型定义
				- 基于已有类型自定义一个类型
				- 扩展方法
			- 类型判断
	*/
	// 类型别名
	type MyInt = int
	var i MyInt = 12      //在编译时，MyInt会转换成int
	fmt.Printf("%T\n", i) //int
	var j int = -2
	fmt.Println(i + j) //10

	//  自定义类型
	type MyInt1 int
	var i1 MyInt1 = 12
	fmt.Printf("%T\n", i1)   //main.MyInt1
	fmt.Println(int(i1) + j) //10

	var i2 MyInt2 = 12
	fmt.Println(i2.string()) //12

	//类型判断
	var a interface{} = "abc"
	//判断类型,type拿到真正类型
	switch a.(type) {
	case string:
		fmt.Println("a string") //a string
	}
	//类型断言
	m := a.(string)
	fmt.Println(m) //abc

	/*
		2、结构体定义、初始化
	*/
	//这样比较麻烦
	persons = append(persons, []string{"a", "b", "c", "d"})
	persons = append(persons, []string{"a", "b", "c"})
	fmt.Println(persons) //[[a b c d] [a b c]]
	// 初始化结构体
	//这种初始化必须全部给值，且必须顺序对应上
	p1 := Person{
		"tom",
		15,
		"北京",
		1.80,
	}
	// 这种初始化可以不全给值
	fmt.Println(p1) //{tom 15 北京 1.8}
	p2 := Person{
		name: "baby",
	}
	fmt.Println(p2) //{baby 0  0}

	var persons []Person
	persons = append(persons, p1)
	persons = append(persons, p2)
	persons = append(persons, Person{
		name: "zhang3",
	})
	fmt.Println(persons) //[{tom 15 北京 1.8} {baby 0  0} {zhang3 0  0}]

	persons2 := []Person{
		{
			name: "a",
		},
		{
			"b", 13, "h", 8,
		},
	}
	fmt.Println(persons2) //[{a 0  0} {b 13 h 8}]

	// 结构体也可以这样赋值
	var p Person
	p.name = "tom"
	fmt.Println(p.name) //tom

	/*
		3、匿名结构体
	*/
	//先声明、并实例化好
	adddres := struct {
		province string
		city     string
		address  string
	}{
		"北京市", "朝阳区", "酒店",
	}
	fmt.Println(adddres) //{北京市 朝阳区 酒店}

	/*
		4、结构体的嵌套
	*/
	s := Student{
		p1,
		99,
	}
	fmt.Println(s) //{{tom 15 北京 1.8} 99}

	//这种赋值比较麻烦，因为每次写都有中间p，这就引入匿名嵌套
	s1 := Student{}
	s1.p.name = "tom"
	fmt.Println(s1.p.name) //tom

	// 匿名嵌套---只是访问方便了，初始化是仍然要走p
	s2 := Student1{}
	s2.name = "tom"
	fmt.Println(s2.name) //tom
	// 匿名嵌套---初始化
	s3 := Student1{
		//name:"tom",//报错，要走p
		Person{
			name: "tom",
		},
		99,
	}
	fmt.Println(s3) //{{tom 0  0} 99}

	//匿名嵌套 --- 覆盖效果
	s4 := Student2{}
	s4.Person.name = "111"
	s4.name = "222"
	fmt.Println(s4) //{{111 0  0} 222}

	/*
		5、结构体绑定方法
	*/
	// 接收器---值传递
	p1.print()           //name:111,age:15,addr:北京,height:1.800000
	s1.p.print()         //name:111,age:0,addr:,height:0.000000
	s3.print()           //name:111,age:0,addr:,height:0.000000
	fmt.Println(p1.name) //tom

	//接收器 --- 引用传递
	p1.print1()          //name:111,age:15,addr:北京,height:1.800000
	fmt.Println(p1.name) //111

	// 指针可以调用值绑定的方法，值可以调用指针上绑定的方法
	pp := &Person{
		name: "pp",
	}
	pp.print()  //name:111,age:0,addr:,height:0.000000
	p1.print1() //name:111,age:15,addr:北京,height:1.800000
}
