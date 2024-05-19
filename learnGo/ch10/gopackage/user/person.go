package hi

// 要想导出在别的包使用，首字母大写
func GetUser(u Courses) string {
	return u.Name
}
