package main

import (
	"errors"
	"fmt"
)

/*
	1、函数的定义
	2、变长参数
*/

//	func 函数名(参数1 类型1，参数2 类型2 ...) (返回值类型1，返回值类型2 ...){
//			return 0	//返回值可以没有、1个、多个
//	}

// 定义一个求和函数
// func add(a int, b int) int {
// 也可以这样写
func add(a, b int) (int, error) {
	return a + b, nil
}

// 求和函数、也可以这样写
func add0(a, b int, c float64) (sum int, err error) {
	fmt.Println(a, b, c)
	sum = a + b
	//return sum, nil
	err = nil
	// 如果上边我们已经定义好sum、err了，也可以只写return
	return
}

// 变长参数
func add1(a ...int) (sum int) {
	fmt.Printf("%T\n", a) // []int、slice类型
	for _, v := range a {
		sum += v
	}
	return
}

// 函数返回值为函数
func cal(op string, item ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Printf("这是加法 \n")
		}
	case "-":
		return func() {
			fmt.Printf("这是减法 \n")
		}
	default:
		return func() {
			fmt.Println("其他运算。。。")
		}
	}
}

// 函数作为参数
func cal1(myfunc func(item ...int) int) int {
	return myfunc()
}

// 定义一个全局变量
var local int

// 每次调用函数，全局变量+1
func autoIncrement() int {
	local += 1
	return local
}

// 闭包
func autoIncrement1() func() int {
	local := 0
	//一个函数不能访问另一个函数的局部变量
	//但是在一个函数中定义匿名函数，匿名函数是可以访问该函数的局部变量的---闭包

	return func() int {
		local += 1
		return local
	}
}

// defer使用
func deferReturn() (ret int) {
	defer func() {
		ret++ //defer有能力修改函数返回值
	}()
	return 10
}

// 返回error
func A() (int, error) {
	return 1, errors.New("this is an error")
}

// panic
func B() (int, error) {
	panic("this is a panic") //会导致程序退出
}

func main() {
	/*
		- go函数支持普通函数、匿名函数、闭包
		- go中函数是“一等公民”
			- 函数本身可以当做变量
			- 匿名函数、闭包
			- 函数满足接口
		- go语言中，函数参数全是值传递

	*/
	fmt.Println(add(1, 2))     //3 <nil>
	fmt.Println(add1(1, 2, 3)) //6

	/*
		3、函数一等公民
	*/

	//函数可以像变量一样赋值
	funcVar := add             //这里不能有（），不然就成调用函数了
	fmt.Println(funcVar(3, 3)) //6 <nil>

	//函数返回函数
	cal("+")() //这是加法

	//函数作为参数传递
	fmt.Println(cal1(func(item ...int) int {
		return 1
	})) //1

	// 匿名函数
	myFunc := func() {
		fmt.Println("hello1") //hello1
	}
	myFunc()
	func() {
		fmt.Println("hello2") //hello2
	}()

	/*
		4、go函数闭包性
			- 需求：
				-希望有个函数每调用一次返回结果都是增加一次之后的值
	*/
	//这种解决方案需要额外定义一个全局变量
	for i := 0; i < 5; i++ {
		fmt.Print(autoIncrement(), " ") //1 2 3 4 5
	}
	fmt.Println()
	//不需要定义全局变量也能实现每次调用+1
	next := autoIncrement1() //返回函数
	for i := 0; i < 5; i++ {
		fmt.Print(next(), " ") //1 2 3 4 5
	}
	fmt.Println()
	//如何归0---再定义一个next1
	next1 := autoIncrement1() //返回函数
	for i := 0; i < 5; i++ {
		fmt.Print(next1(), " ") //1 2 3 4 5
	}
	fmt.Println()

	/*
		5、函数中的defer
			-连接数据库、打开文件、开始锁，无论如何最后都要记得去关闭数据库、关闭文件、解锁
	*/
	//var mu sync.Mutex
	//mu.Lock()
	//defer mu.Unlock() //defer后面的代码会放在函数return之前执行

	//如果有多个defer,执行顺序（类似栈，前面的defer后执行）：
	//defer fmt.Print("1", " ")
	//fmt.Print("3", " ")
	//defer fmt.Print("2", " ")
	//输出顺序：3 2 1

	ret := deferReturn()
	fmt.Println(ret) //11~

	/*
		6、error设计理念
			go语言错误处理
				- error 值
				- panic 函数
				- recover 函数
			go语言错误处理理念
				- 其他语言 try...catch...
				- 如果函数出错，函数返回error值去告诉调用者是否成功，要求必须处理error，不能往上抛
					- 因此代码大量出现 if err!=nil
					- 也叫防御性编程
	*/
	_, err := A()
	if err != nil {
		fmt.Println(err) //this is an error
	}
	//也可以这样简洁写：
	if _, err := A(); err != nil {
		fmt.Println(err) //this is an error
	}

	/*
		7、panic & recover
			- 会导致程序退出
			- 使用场景：
				-比如我的服务想要启动，必须有些依赖服务准备好，日志文件存在、mysql能链接通、比如配置文件没有问题，这个时候服务才能肩动
				 如果我们的服务启动检查中发现了这些任何一个不满足那就调用panic
			- 即使我们没写panic，也会触发，比如map没初始化就添加值
				- 引出 recover
					- 可以捕获panic，避免程序因某个异常而崩掉
					- defer需要放在panic之前定义，另外recover只有在defer调用的函数中才会生效
					- recover处理异常后，逻辑并不会恢复到panic的那个点去,而是直接return
					- 多个defer会形成栈，后定义的defer.会先执行
	*/
	//B() //报错，panic: this is a panic

	//map没有初始化就添加值没有初始化
	//var names map[string]string //没有初始化
	//names["go"] = "123"         //会出错，panic: assignment to entry in nil map

	// recover 使用
	defer func() { //注意要放在出现会panic错误代码的前边！！！
		if r := recover(); r != nil {
			fmt.Println("没事了~", r) //没事了~ assignment to entry in nil map
		}
	}()
	var names map[string]string
	names["go"] = "123" //不报错了~，因为panic被recover捕获了
	fmt.Println(names)

}
