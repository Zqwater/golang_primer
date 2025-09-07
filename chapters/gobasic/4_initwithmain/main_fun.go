package initwithmain

/* main函数 */
func MainFun() {
	/*
		解释：
		Go程序的入口点是main包中的main函数
		每个Go程序必须包含一个main包
		main函数不需要任何参数，也不返回任何值

		执行流程：
		1.先执行所有包的init函数
		2.执行main函数
		3.程序正常执行结束，退出程序

		特点：
		1.每个程序只能有一个main函数
		2.main函数只能在main包中
		3.不能被其他函数调用
		4.是程序的唯一入口点

	*/
}
