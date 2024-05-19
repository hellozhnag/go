package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	b := "hello"
	c := "bello"
	fmt.Println(a == b) //true
	fmt.Println(a > c)  //true，因为h的assic码比b的assic码大

	//strings常用方法
	name := "haha攻城狮"
	fmt.Println(strings.Contains(name, "狮"))              //true
	fmt.Println(len([]rune(name)))                        //7
	fmt.Println(strings.Count(name, "h"))                 //2
	fmt.Println(strings.Split(name, "攻城"))                //[haha 狮]
	fmt.Println(strings.HasPrefix(name, "ha"))            //true，作为前缀
	fmt.Println(strings.HasSuffix(name, "攻城狮"))           //true,作为后缀
	fmt.Println(strings.Index(name, "城"))                 //7，字符串存在则返回第一个字节位置
	fmt.Println(strings.IndexRune(name, []rune(name)[5])) //7,rune存在则返回第一个字节位置
	fmt.Println(strings.Replace(name, "ha", "hello", -1)) //hellohello攻城狮
	//-1表示全替换，1表示替换前一个，2表示替换前两个
	fmt.Println(strings.Replace(name, "ha", "hello", 1)) //hellohello攻城狮
	fmt.Println(strings.ToLower("GO"))                   //go
	fmt.Println(strings.Trim("##he#llo#", "#"))          //he#llo

}
