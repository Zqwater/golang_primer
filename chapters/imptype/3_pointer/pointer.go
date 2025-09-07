package pointer

import (
	"basicGo/basicLesson/golog"
	"unsafe"
)

var (
	logger = golog.Logger
)

/* golang指针类型 */
func Pointer() {

	/* 定义 */
	/*
	 1. 存储变量的内存地址的变量类型，一种间接访问和操作内存的方式
	 2. 取地址&和指针取值*为以对指针操作
	 3. 指针地址（存储地址的变量的值也就是指针变量的值），指针类型（指针变量的类型），指针变量（存储的变量地址的变量）
	 4. golang是安全指针，不支持指针运算，可用unsafe包实现
	*/

	/* 指针变量的声明 */
	var i1 int = 1001 //基本类型指针
	var p1 *int = &i1 //从变量i1中取出i1的内存地址赋值给p1

	var p2 *bool = new(bool) //使用new获取,该指针对应的值为该类型的零值 *p2 = false

	type User struct {
		Name string
	}

	p3 := &User{ //结构体指针获取
		"static",
	}

	arr1 := [3]complex128{complex(1.1, 3.14), complex(1, 2), complex(101, 22.33)}
	p4 := &arr1 //数组指针

	s1 := arr1[:]
	p5 := s1 //p4=p5

	f1 := func(a int) int {
		return a * a
	}
	p6 := &f1 //函数指针

	logger.Printf("p1=%p,p2=%p,p3=%p,p4=%p,p5=%p,p6=%p", p1, p2, p3, p4, p5, p6)
	logger.Printf("指针取值v:*p1=%d,*p2=%t,*p3=%v,*p4=%v,p5=%p\n", *p1, *p2, *p3, *p4, p5) //p6函数不能取地址和取得值

	/* 指针的使用 */
	//unsafe包操作指针
	//unsafe包 突破类型限制，是类型系统和内存布局之间的桥梁
	/*
	 1. unsafe.Pointer可以指向任意类型的指针,作为指针转换之间的桥梁
	 2. unsafe.Sizeof可以获取指针指向的变量的大小
	 3. unsafe.Alignof可以获取指针指向的变量的对齐方式
	 4. unsafe.Offsetof可以获取指针指向的变量的偏移量
	*/

	//unsafe.Pointer 是可以指向任何类型的特殊指针
	/*
		1.unsafe.Pointer <-> 任何类型指针
		2.unsafe.Pointer <-> uintptr(用于指针运算)
		3.uintptr运算结果 -> unsafee.Pointer
		4.可以转换指针值如下面的int和float32
	*/
	var i int = 20
	p := &i
	fp := (*float32)(unsafe.Pointer(p))
	*fp = 3.14
	logger.Printf("safe.Pointer转换*int为*float32:%d\n", i) //i的值已经更改为3.14

	//unsafe.Sizeof 获取类型内存大小
	/*
	  空结构体sizeof为0,而且相同进程的空结构体地址相同
	*/
	var ii int = 10                          //8
	var bb byte = 'A'                        //1
	var tt bool = true                       //1
	var cc complex128 = complex(3.14, 2.718) //16
	var ss string = "hello"                  //16
	var aa [2]int = [2]int{1, 2}             //16=8*2
	// eptStu := struct{}{} //size= 0
	stu := struct { //24 = 8 + 16
		age  int
		name string
	}{
		0,
		"world",
	}

	logger.Printf("值类型内存大小:int=%d,byte=%d,bool=%d,complex=%d,string=%d,array=%d,struct=%d\n",
		unsafe.Sizeof(ii), unsafe.Sizeof(bb), unsafe.Sizeof(tt), unsafe.Sizeof(cc),
		unsafe.Sizeof(ss), unsafe.Sizeof(aa), unsafe.Sizeof(stu),
	)

	pi := &ii //都是8
	pb := &bb
	pt := &tt
	pc := &cc
	ps := &ss
	pa := &aa
	pstu := &stu

	logger.Printf("值类型对应指针内存大小:*int=%d,*byte=%d,*bool=%d,*complex=%d,*string=%d,*array=%d,*struct=%d\n",
		unsafe.Sizeof(pi), unsafe.Sizeof(pb), unsafe.Sizeof(pt), unsafe.Sizeof(pc),
		unsafe.Sizeof(ps), unsafe.Sizeof(pa), unsafe.Sizeof(pstu),
	)

	sl := make([]int, 3)          //24
	mm := make(map[int]string, 3) //8
	ch := make(chan int, 3)       //8

	logger.Printf("切片类型:%d,map类型:%d,chan类型:%d\n", unsafe.Sizeof(sl), unsafe.Sizeof(mm), unsafe.Sizeof(ch))

}
