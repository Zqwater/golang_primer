package structtype

import (
	"golang_primer/golog"
	"time"
	"unsafe"
)

/* 空结构体*/
/*
 1. 是一种占零内存不包含任何字段的类型，底层返回的都是局变量zerobase的地址
 2. 同一个进程地址唯一
*/
func EmptyStruct() {

	logger := golog.Logger
	/* 定义和声明 */
	var empty1 struct{}
	empty2 := struct{}{} //匿名类型

	logger.Printf("sizeof empty1=%d\n", unsafe.Sizeof(empty1))
	logger.Printf("sizeof empty2=%d\n", unsafe.Sizeof(empty2))

	logger.Printf("address of empty1=%p\n,address of empty2=%p\n", &empty1, &empty2) //地址唯一

	/* 空结构体的应用场景 */

	// 实现方法接收器(无状态对象)
	EpyImp1()

	//实现高效集合Set
	EpyImp2()

	//实现通道信号传递
	EpyImp3()

}

/* 实现方法接收器(无状态对象) */
/*
	1.零内存开销
	2.保持方法集组织性
	3.便于未来扩展字段
*/
type Emp1 struct{}

func (e Emp1) Open() {
	logger.Printf("Emp1 open")
}

func (e Emp1) Close() {
	logger.Printf("Emp1 close")
}

func EpyImp1() {
	var emp1 Emp1
	emp1.Open()
	emp1.Close()
}

/* 实现高效集合Set */

type Set[K comparable] map[K]struct{}

func (s Set[K]) Add(v K) {
	s[v] = struct{}{}
}

func (s Set[K]) Delete(v K) {
	delete(s, v)
}

func (s Set[K]) Contain(v K) bool {
	_, ok := s[v]
	return ok
}

func EpyImp2() {

	set1 := make(Set[int])
	set1.Add(20)
	set1.Delete(1)
	r := set1.Contain(20)
	logger.Printf("set1 contains 20:%t\n", r)

}

/* 实现通道信号传递 */
/*
 1.明确表达”无数据仅信号”语义
 2.标准库 context.Context 的 Done() 采用相同设计
*/

func EpyImp3() {

	stop := make(chan struct{})
	defer close(stop)

	go func(st chan struct{}) {
		for {
			select {
			case <-st:
				logger.Println("work stop")
			default:
				logger.Println("go on working")
			}
		}
	}(stop)

	time.Sleep(50 * time.Microsecond)

}
