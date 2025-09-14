package varwithconst

import "golang_primer/golog"

/* go语言变量 */
func Var() {
	logger := golog.Logger

	/* 变量的声明（基本类型为例） */
	//标准声明和短变量声明
	//集体声明
	var (
		b1   bool
		i1   int
		by1  byte = 97 //uint8
		r1   rune      //int32
		f1   float32
		c1   complex64
		str1 string
		up1  uintptr //以存储指针的uint32或uint64整数
	)
	logger.Printf("变量的集体声明b1=%t,i1=%d,by1=%c,r1=%d,f1=%0.2f,c1=%v,str1=%s,up1=%v\n",
		b1, i1, by1, r1, f1, c1, str1, up1,
	)
	//匿名变量 _

	/* 变量内存分配机制 */
	/*
	 1.函数内变量分配到栈上，随着函数的周期创建和销毁
	 2.变量逃逸出函数作用域时分配在堆上，通过指针共享变量，由gc管理声明周期
	 3.零值初始化：未显示初始化的变量会被赋予其类型的默认零值
	*/

	/* 变量作用域规则 */
	/*
	 1.函数内变量：在函数内部定义的变量，只能在函数内部访问，函数调用结束后变量会被销毁
	 2.全局变量：在函数外部定义的变量，在整个包中都可以访问
	 3.局部变量：在函数内部定义的变量，只能在函数内部访问，函数调用结束后变量会被销毁
	*/
	var blockVar bool

	{
		var blockVar bool = true
		logger.Printf("块内变量blockVar=%t\n", blockVar)
	}

	logger.Printf("块内变量在块外是否可以访问:%t\n", blockVar)

}
