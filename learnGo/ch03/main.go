package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "ha慕课"
	//len方法获取字节长度
	fmt.Println(len(s)) //8
	//在utf8中，英文占一个字节，中文占三个字节
	fmt.Println("%X\n", []byte(s)) //[104 97 230 133 149 232 175 190]

	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
		//（字节索引，Unicode码）
		//(0 68)
		//(1 61)
		//(2 6155) //中文字符占3字节即第2 3 4
		//(5 8BFE) //索引不连续
	}
	fmt.Println()
	// 查看字符串中字符数量
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //4

	//转换为字节数组
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	/*
		转换为rune数组，就能通过索引获取每一个字符
	*/
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
