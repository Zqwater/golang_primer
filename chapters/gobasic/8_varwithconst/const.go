package varwithconst

import "golang_primer/golog"

/* go语言常量 */
func Const() {
	logger := golog.Logger
	/* 常量内存特性 */
	/*
		编译期确定：常量值在编译阶段确定，不会在运行时计算
		无内存地址：常量没有内存地址，无法使用 & 获取指针
		直接替换：编译器在代码中使用常量值时直接进行替换
		类型安全：类型化常量遵循严格的类型检查规则
	*/

	/* 常量的声明 */
	//标准声明
	//批量声明
	const (
		HTTP_OK  = 200
		Pi       = 3.1415927
		MAX_CONN = 1024
	)

	//类型化常量
	const prefix string = "ERR_"
	const http_code int = 301

	/* iota的用法 iota 是 Go 语言的常量计数器，在 const 关键字出现时重置为 0，每新增一行常量声明 iota 自增 1*/
	//基础枚举
	const (
		Monday    = iota + 1 // 1
		Tuesday              // 2，与上面的一致
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)
	//表达式中的枚举
	const ( //简单表达式
		_  = iota             // 跳过0
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB = 1 << (10 * iota) // 1 << 20 = 1,048,576
		GB = 1 << (10 * iota) // 1 << 30
		TB = 1 << (10 * iota) // 1 << 40
	)

	const ( //复杂表达式
		Read    = 1 << iota //001
		Write               //010
		Execute             //100
	)
	// 权限组合
	userPermisson := Read | Write             //011
	adminPermission := Read | Write | Execute //111
	logger.Printf("userPermisson: %b, adminPermission: %b", userPermisson, adminPermission)

	//iota重置和恢复
	const (
		A = iota // 0
		B
		C = 100      //中断iota
		D            //100(和上面保持一致)
		E = iota + 1 //5 恢复计数
		F            //6
	)
	logger.Printf("A: %d, B: %d, C: %d, D: %d, E: %d, F: %d", A, B, C, D, E, F)

	//自定义启始值
	const (
		_     = iota + 100
		one   //101
		two   //102
		three //103
	)
	logger.Printf("自定义启始值:one: %d, two: %d, three: %d", one, two, three)

	//跳值技巧
	const (
		_     = iota //0
		Red          //1
		_            //跳过2
		Blue         //3
		Green = iota //4
	)

}
