package main

import (
	"fmt"
)

type Person struct {
	name string
}

func changeName(person *Person) {
	person.name = "tom"
}

func (p *Person) sayHi() {
	fmt.Println("Hi")
}

// 通过指针交换两个值
func swap(a, b *int) {
	a, b = b, a
}
func swap1(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	p := Person{"123"}
	changeName(&p)
	fmt.Println(p.name) //tom

	/*
		1、指针的定义
	*/
	var po *Person
	po = &p
	(*po).name = "po"
	fmt.Println(po) //&{po}
	//go在这里做了优化，可以直接通过指针访问元素
	po.name = "po1"
	fmt.Println(po)        //&{po1}
	fmt.Printf("%p\n", po) //0xc00010a030
	// go语言的指针不能参加运算如加法运算
	// 其实也是可以做运算的，在unsafe包里

	/*
		2、指针初始化
	*/
	var b *int
	fmt.Println(b) //<nil>

	//var p1 *Person
	//fmt.Println(p1.name)//报错，invalid memory address or nil pointer dereference

	//为了避免nil，需要对指针进行初始化
	//指针第一种初始化方式
	p2 := &Person{}
	fmt.Println(p2.name) //打印空
	//指针第二种初始化方式
	var p3 Person //会进行默认初始化
	p4 := &p3
	fmt.Println(p4.name) //打印空
	//指针第三种初始化方式
	var p5 *Person = new(Person) //new方法会初始化并把地址返回
	fmt.Println(p5.name)         //打印空

	//初始化的两个关键字：
	//	- map、channel、slice ---make，map必须初始化
	//	- 指针 ---new,指针最好初始化

	/*
		3、通过swap交换指针的值
	*/
	a1, b1 := 1, 2
	//& 指的是传递地址值
	swap(&a1, &b1)
	fmt.Println(a1, " ", b1) //1   2
	swap1(&a1, &b1)
	fmt.Println(a1, " ", b1) //2   1

	/*
		4、nil
			- 不同的数据类型零值不同
				- bool --- false
				- numbers --- 0
				- string --- ""
				- pointer --- nil
				- slice --- nil
				- map --- nil
				- channel --- nil
				- myinterface --- nil
				- function --- nil
				- struct --- 默认不是nil，默认值是具体字段的默认值
					- struct可以使用 == 判断是否等于（判断每一个属性是否都相等）
	*/
	//slice默认值
	//当我们判断slice是否为nil，实际上是判断slice底层的struct的element指针指向的是否为nil

	var ps []Person
	if ps == nil {
		fmt.Println("ps is nil") //ps is nil
	}
	var ps2 = make([]Person, 0)
	if ps2 == nil {
		fmt.Println("ps2 is nil")
	}

	var m map[string]string             //nil map
	var m2 = make(map[string]string, 0) //empty map
	if m == nil {
		fmt.Println("m is nil") //m is nil
	}
	if m2 == nil {
		fmt.Println("m2 is nil")
	}
	//用for遍历，都没报错，因为go做了优化
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	//赋值
	//m["ha"] = "ha"//报错，panic: assignment to entry in nil map
	m2["ha"] = "ha" //正常

}
