package array

import "basicGo/basicLesson/golog"

/* golang数组类型 值类型*/
func Array() {

	logger := golog.Logger

	//数组的定义
	/*
		1.由长度和元素类型共同定义
		2.内存布局：一段连续的内存空间，相同类型的元素
		3.编译时检查：支持==和!=操作符，要求长度和类型完全匹配
		4.不支持数组展开，但是支持切片展开
	*/

	//编译期间处理
	/*
		1.长度确定:必须在编译时确定（常量表达式）
		2.边界检查：已知索引的越界访问在编译期检查
		3.长度推导：支持[...]语法自动计算长度
	*/

	//数组的声明和初始化
	var arr1 [3]int
	var arr2 [3]string = [3]string{"a", "b", "c"}
	var arr3 = [...]byte{'A', 'a'}
	arr4 := [3]float32{3.14, 2.718, 1.41}
	arr5 := [...]struct { //匿名结构体数组
		Age  int
		Name string
	}{
		{18, "static"},
		{22, "water"},
	}

	arr6 := [...][]int{ //切片数组
		{1, 2, 3},
		{200, 301, 404},
		{1024, 1024 * 1024},
	}

	arr7 := [...]map[int]string{ //map数组
		{1: "a", 2: "b"},
		{100: "aa", 200: "bb"},
	}

	arr8 := [...]func(a int) int{ //函数数组
		func(a int) int {
			return -a
		},
		func(a int) int {
			return a * a
		},
		func(a int) int {
			return a + a
		},
	}

	logger.Printf("arr1=%v,arr2=%v,arr3=%v,arr4=%v,arr5=%v,arr6=%v,arr7=%v,arr8=%v",
		arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8,
	)

	//数组和切片的关系
	/*
		1.切片是对底层数组的一种引用，切片本身不存储任何数据，只是描述了底层数组的一段
		2.切片的长度可以动态变化，而数组的长度是固定的。切片支持展开，数组不支持
		3.切片是引用传递，对切片元素的修改会影响底层数据的元素数据
		4.切片是不可比较类型，不能通过== !=来比较，但是可以通过一些深度比较函数判断:reflect.DeepEqual go-cmp库的cmp.Equal
	*/

	var arr [4]int = [4]int{100, 101, 102, 103}
	var s []int = arr[1:] //数组转为切片,隐式
	arrs := (*[2]int)(s)  //切片转为数组
	logger.Printf("数组切片转换。数组转为切片s=%v,切片转为数组arrs=%v", s, arrs)

	//多维数组
	var marr1 [3][2]string = [...][2]string{
		{"a", "b"},
		{"aa", "bb"},
		{"golang", "java"},
	}

	marr2 := [...][3]bool{
		{true, false, true},
		{false, true, false},
		{true, true, true},
	}
	logger.Printf("遍历marr1:")
	for i := 0; i < len(marr1); i++ {
		for j := 0; j < len(marr1[i]); j++ {
			logger.Printf("%v\n", j)
		}
	}

	logger.Printf("遍历marr2地址:")
	for _, v := range marr2 {
		for _, v2 := range v {
			logger.Printf("%p\n", &v2)
		}
	}

	//数组的使用

	for i := range marr2 { //高效遍历多维数组
		for j := range marr2[i] {
			logger.Printf("marr2[%d][%d] = %v\n", i, j, marr2[i][j])
		}
	}

	//使用建议
	/*
		小数组直接在栈上分配，大数组使用切片控制分配到堆上
	*/

}
