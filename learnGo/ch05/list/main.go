package main

import (
	"container/list"
	"fmt"
)

func main() {
	// list

	/*
		1、list初始化
		2、首尾插入元素
	*/

	//两种初始化
	//var mylist list.List
	var mylist = list.New()

	mylist.PushBack("go")
	mylist.PushFront("go1")
	mylist.PushBack("go2")
	fmt.Println(mylist) //{{0xc000108120 0xc000108150 <nil> <nil>} 3}

	/*
		3、遍历
	*/

	//正向遍历
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
		//go1
		//go
		//go2
	}
	//反向遍历
	for j := mylist.Back(); j != nil; j = j.Prev() {
		fmt.Println(j.Value)
		//go2
		//go
		//go1
	}
	/*
		4、中间插入元素
	*/
	//在go前边插入go0.5
	i := mylist.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "go" {
			break
		}
	}
	mylist.InsertBefore("go0.5", i)
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ") //go1 go0.5 go go2
	}
	fmt.Println()
	fmt.Println(i.Value) //go

	/*
		5、删除元素
	*/
	mylist.Remove(i)
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ") //go1 go0.5 go2
	}
}
