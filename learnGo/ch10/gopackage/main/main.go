package main

// 使用其他包，必须import
import (
	"fmt"
	"learnGo/ch10/gopackage/user"
	//也可以定义别名，这样的话就只能用别名访问，而不是package
	//u "learnGo/ch10/gopackage/user"
	//这种写法可以省略hi，意味着把这个目录下的所有导入main
	//. "learnGo/ch10/gopackage/user"
	// 引入了但不使用，主要是初始化,自动调用init
	//_ "learnGo/ch10/gopackage/user"
)

func main() {
	//通过package访问
	c := hi.Courses{
		Name: "go",
	}
	fmt.Println(c)             //{go}
	fmt.Println(hi.GetUser(c)) //go

}
