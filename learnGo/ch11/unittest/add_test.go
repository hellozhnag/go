package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

//go test命令是一个按照一定约定和组织的测试代码驱动程序、在包目录中，所有以_test.go为后缀的源码文件都会被gotest运行到
// go build 不会把这些以_test.go结尾的测试文件打包到最后的可执行文件中
// test文件有4类，只关心前两类。Test开头的为功能测试，Benchmark开头的为性能测试，example，模糊测试

//单元测试文件运行可以没有main函数
//单元测试写在同目录下，避免需测试的函数首字母小写而访问不了

// 编写测试用例
// 点击运行或者命令行go test . 运行
func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("add(1,3) = %d; want 4", re)
	}
}

// 跳过耗时的单元测试
// go test -short
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	re := add(1, 5)
	if re != 6 {
		t.Errorf("add(1,5) = %d; want 6", re)
	}
}

// 基于表格驱动测试
func TestAdd3(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{1, 3, 4},
		{2, 4, 6},
	}

	for _, v := range dataset {
		re := add(v.a, v.b)
		if re != v.out {
			t.Errorf("add(%d, %d) = %d; want %d", v.a, v.b, re, v.out)
		}
	}
}

// benchmark 性能测试
// go test -bench=".*"
// 我们希望核心函数，性能高
func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 123
	b = 456
	c = 579
	for n := 0; n < bb.N; n++ {
		if actual := add(a, b); actual != c {
			bb.Errorf("add(%d,%d) = %d; want %d", a, b, actual, c)
		}
	}
}

// 测试字符串三种相加操作性能，sprintf、builder、+
// sprintf 和 + 差别不大
// builder比那俩快了接近100倍
const numbers = 10000

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str += strconv.Itoa(j)
		}
	}
	b.StopTimer()
}
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
