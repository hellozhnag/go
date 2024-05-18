package main

import "fmt"

func main() {

	/*
		1、切片的定义和添加元素
	*/
	var courses []string
	fmt.Printf("%T\n", courses)     //[]string
	courses = append(courses, "go") //接受类型和放置类型要一样！！！，都是courses
	fmt.Printf("%T\n", courses)     //[]string
	fmt.Println(courses)            //[go]
	fmt.Println(courses[0])         //go

	/*
		2、切片初始化
			- 从数组直接创建
			- 使用string{}
			- make
	*/
	allCourses := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	fmt.Println(allCourses) //[go grpc gin mysql elasticsearch]
	// 第一种
	courseSlice := allCourses[0:2] //[go grpc]，左闭右开[)
	fmt.Println(courseSlice)
	fmt.Printf("%T\n", courseSlice)
	// 第二种
	coursesSlice1 := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	fmt.Println(coursesSlice1)        //[go grpc gin mysql elasticsearch]
	fmt.Printf("%T\n", coursesSlice1) //[]string
	//第三种
	courseSlice2 := make([]string, 3) //先预分配空间3
	courseSlice2[0] = "go"
	courseSlice2[1] = "grpc"
	courseSlice2[2] = "gin"
	fmt.Println(courseSlice2) //[go grpc gin]

	var courseSlice3 []string //没有预先分配空间，只能使用append添加
	//courseSlice3[0] = "go"    //运行报错，index out of range [0] with length 0
	courseSlice3 = append(courseSlice3, "go")

	/*
		3、访问切片的元素
			- 访问单个，跟数组一样
			- 访问多个
	*/
	courseSlice4 := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	fmt.Println(courseSlice4[1])          //grpc
	fmt.Println(courseSlice4[:])          //[go grpc gin mysql elasticsearch]
	fmt.Println(courseSlice4[2:])         //[gin mysql elasticsearch]
	fmt.Println(courseSlice4[:3])         //[go grpc gin]
	fmt.Printf("%T\n", courseSlice4[1:2]) //[]string

	/*
		4、slice for循环
			- 跟数组一样
	*/

	/*
		5、slice 添加
	*/
	courseSlice5 := []string{"go", "grpc"}
	courseSlice5 = append(courseSlice5, "gin", "mysql", "elasticsearch")
	fmt.Println(courseSlice5) //[go grpc gin mysql elasticsearch]
	courseSlice6 := []string{"go1", "grpc1"}
	courseSlice6 = append(courseSlice6, courseSlice5...)
	fmt.Println(courseSlice6) //[go1 grpc1 go grpc gin mysql elasticsearch]
	courseSlice6 = append(courseSlice6, courseSlice5[1:3]...)
	fmt.Println(courseSlice6) //[go1 grpc1 go grpc gin mysql elasticsearch grpc gin]

	/*
		6、slice 删除元素
			- 左右slice拼接，重新生成slice
	*/
	courseSlice7 := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	//删除 gin
	myslice := append(courseSlice7[0:2], courseSlice7[3:]...)
	fmt.Println(myslice) //[go grpc mysql elasticsearch]
	// 删除mysql、elasticsearch
	myslice1 := courseSlice7[:3]
	fmt.Println(myslice1) //[go grpc mysql]

	/*
		7、复制 slice
	*/
	coursesSliceCopy := courseSlice7 //本质没有拷贝，只是指向同一
	coursesSliceCopy1 := courseSlice7[:]
	fmt.Println(coursesSliceCopy, " ", coursesSliceCopy1)

	var coursesSliceCopy2 = make([]string, len(courseSlice7))
	copy(coursesSliceCopy2, courseSlice7)
	fmt.Println(coursesSliceCopy2) //[go grpc mysql elasticsearch elasticsearch]
	courseSlice7[0] = "hello~"
	fmt.Println(coursesSliceCopy, coursesSliceCopy1, coursesSliceCopy2) //输出为下面3行
	//[hello~ grpc mysql elasticsearch elasticsearch]
	//[hello~ grpc mysql elasticsearch elasticsearch]
	//[go grpc mysql elasticsearch elasticsearch]

	coursesSliceCopy4 := courseSlice7[:3]
	fmt.Println(coursesSliceCopy4) //[hello~ grpc mysql]
	courseSlice7[1] = "test"
	fmt.Println(coursesSliceCopy4) //[hello~ test mysql]

}
