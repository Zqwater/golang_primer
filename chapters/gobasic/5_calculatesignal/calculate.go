package calculatesignal

import (
	"basicGo/basicLesson/golog"
)

/* 运算符 */
func CalculateSignal() {
	logger := golog.Logger

	/* 按照操作数个数分类*/

	/* 一元运算符
	+ - ！^(按位取反) & * <-
	++ -- 在go中算单独语句
	*/
	var i int = -0b10101
	reverse := ^i
	logger.Printf("^按位取反:%b\n", reverse)
	/* 二元运算符 */
	/* 算术运算符 + - * / % */
	/* 关系运算符 == !=  < > <= >= */
	/* 逻辑运算符 && || !*/
	var a, b bool
	if !a || b { //逻辑非
		logger.Printf("逻辑或")
	}
	if !a && !b {
		logger.Printf("逻辑与")
	}
	/* 位运算符 & | ^(异或) << >>（乘除2^n） */
	left := i << 2
	right := i >> 2
	logger.Printf("左移2位:%b\n", left)
	logger.Printf("右移2位:%b\n", right)

	/* 赋值运算符 = += -= *= /= %= <<= >>= &= |= ^=*/

	/* 三元运算符 没有三元运算符 a?b:c 可以用if else实现*/

	/* 运算符优先级
	1.括号()
	2.++ -- 前缀 后缀
	3.* / %
	4.+ -
	5.位运算符 << >> & ^ |
	6.关系运算符 == != < > <= >=
	7.逻辑运算符 && || !
	8.赋值运算符 = += -= *= /= %= <<= >>= &= |= ^=
	9.短路求值 左边求值决定是否还需要计算右边
	*/

}
