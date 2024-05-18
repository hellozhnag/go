package main

import "fmt"

func main() {
	//数组
	/*
		1、数组定义
	*/
	var courses1 [3]string
	//[3]string为只有三个元素的数组类型,和[4]string不是一种数据类型！！！
	//[]string和[3]string不是一种类型！！！，[]string为切片
	var courses2 [4]string

	//%T表示输出类型
	fmt.Printf("%T\r\n", courses1)
	fmt.Printf("%T\n", courses2)
	//courses2=courses1 //报错，类型不匹配

	/*
		2、数组赋值
	*/
	courses1[0] = "go"
	courses1[1] = "python"
	courses1[2] = "java"
	fmt.Println(courses1)

	/*
		3、数组遍历
	*/
	for _, course := range courses1 {
		fmt.Println(course)
	}
	for i := 0; i < len(courses1); i++ {
		fmt.Println(courses1[i])
	}

	/*
		4、数组初始化
	*/
	var courses3 [3]string = [3]string{"go", "python", "java"} //这里的定义类型可以省略
	fmt.Println(courses3)                                      //[go python java]
	var courses4 [4]string = [4]string{2: "gin"}
	fmt.Println(courses4) //[  gin ]
	courses5 := [...]string{"go", "java"}
	fmt.Println(courses5)          //[go java]
	fmt.Printf("%T\r\n", courses5) //[2]string

	/*
		5、数组间的比较
	*/
	//if courses4==courses5 {}//报错，无法比较，因为类型不匹配
	if courses1 == courses3 { //可以比较，类型匹配
		//会比对数组每一个元素
		fmt.Println("equal...")
	}

	/*
		6、多维数组定义赋值
	*/
	var courseinfo [3][4]string //3行4列
	courseinfo[0] = [4]string{"go", "1h", "tom", "go体系"}
	//也可以这样赋值
	//courseinfo[0][0]="go"
	//courseinfo[0][1]="python"
	courseinfo[1] = [4]string{"grpc", "2h", "tom", "grpc框架"}
	courseinfo[2] = [4]string{"gin", "1.5h", "tom", "gin高级开发"}

	/*
		7、多维数组遍历
	*/
	fmt.Println(len(courseinfo)) //3
	for i := 0; i < len(courseinfo); i++ {
		for j := 0; j < len(courseinfo[i]); j++ {
			fmt.Print(courseinfo[i][j], " ")
		}
		fmt.Println()
	}
	for _, course := range courseinfo {
		for _, value := range course {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}

}
