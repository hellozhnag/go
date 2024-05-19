package hi

// 一般package和目录名一致
// package 用来组织源码，多个go源码的集合，代码复用的基础,比如 fmt、os...
// 每个源码文件开始都必须声明package，同一个文件夹下的源码用同一个包名，包名可以自定义
// 在同文件夹下的其他源文件中可直接访问
// 要想导出在别的包使用，首字母大写！！！
type Courses struct {
	Name string
}
