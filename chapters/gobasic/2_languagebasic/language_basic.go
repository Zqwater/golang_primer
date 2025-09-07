package languagebasic

/* golang语言基础知识 */

func LanguageBasic() {

	//1.命名规则
	/*
		1. 命名规则： 1. 以字母或下划线开头（不能以数字开头） 2. 后面可以跟字母、数字、下划线 3. 不能使用关键字 4. 不能使用空格 5. 不能使用中文
		2. 命名规范： 1. 驼峰命名法 2. 下划线命名法 3. 缩写命名法
		3. 示例: 首字母大写表示可以其他包可见和使用
			1. 驼峰命名法： firstName lastName
			2. 下划线命名法： first_name last_name
			3. 缩写命名法： firstN lastN
	*/

	//2.25个关键字 37个保留字
	/*
		1.关键字：
		定义: package import var const type interface chan map struct func go defer
		程序控制：if else switch case default for range break continue return goto fallthrough select
		2.保留字：
		常量: true false iota nil
		类型：int int8 int16 int32(rune) int64 uint uint8(byte) uint16 uint32 uint64 uintptr float32 float64 complex64 complex128
		函数: make len cap new append copy close delete complex real imag panic recover
	*/

	//可见性
	/*
		1.声明在函数内部，仅函数内部可见
		2.声明在函数外部，对当前包可见，类似于全局变量
		3.声明在函数外部，且首字母大写，对其他包可见，可以被其他包导入使用
	*/

}
