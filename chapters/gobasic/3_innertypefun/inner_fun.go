package innertypefun

import (
	"basicGo/basicLesson/golog"
	"strings"
)

/* 内置函数 */
func InnerFunc() {
	logger := golog.Logger

	/* 基础查询函数 */

	/* len 查看string，数组 slice map chan的元素个数 */

	//string类型 底层是byte数组，不可变序列
	var str1 string = "golang length" //13
	str2 := "go语言"                    //8
	var str3 strings.Builder
	str3.WriteString(str1)
	str3.WriteString(str2) //13+8=21
	logger.Printf("len of str1=%d,str2=%d,str3=%d\n", len(str1), len(str2), len(str3.String()))

	//array
	var arr1 [3]int = [3]int{1, 7, 5}
	arr2 := [5]string{"ok", "fine", "good"}
	logger.Printf("len of arr1=%d,arr2=%d\n", len(arr1), len(arr2))

	//slice 直接读取slice header中的len字段
	var s1 []int = []int{1, 2, 3}
	s2 := []string{"ok", "fine", "good"}
	s3 := make([]float32, 2, 3)
	logger.Printf("len of s1=%d,s2=%d,s3=%d\n", len(s1), len(s2), len(s3))

	//map 返回当前存储的键值对数量
	var m1 map[int]int = map[int]int{
		1: 1,
		2: 4,
		3: 9,
	}
	m1[5] = 25

	m2 := make(map[int]string, 3)
	logger.Printf("len of m1=%d,m2=%d\n", len(m1), len(m2)) //4,0

	//chan 返回缓冲区中未读取的元素数
	ch1 := make(chan int, 3)
	ch2 := make(chan<- string, 2)
	ch2 <- "dream"
	ch2 <- "boy"
	logger.Printf("len of ch1=%d,ch2=%d\n", len(ch1), len(ch2)) //0,2

	/* cap  slice和chan的容量 */
	//slice 读取slice header中的cap字段
	s4 := make([]int, 2, 3)
	logger.Printf("before cap of s4=%d\n", cap(s4)) //3
	s4 = append(s4, 2, 3, 4, 5)
	logger.Printf("after cap of s4=%d\n", cap(s4)) //6 参考扩容机制

	//chan 返回初始化时指定的缓冲区大小
	ch3 := make(chan int, 3)
	logger.Printf("before cap of ch3=%d\n", cap(ch3)) //3
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	// ch3 <- 4 //向已经满了的chan中发送数据fatal error
	logger.Printf("after cap of ch3=%d\n", cap(ch3)) //3

	/* 内存管理函数  */
	/* make 初始化slice、map或channel
	slice：分配底层数组并返回slice header
	map：初始化哈希表数据结构
	channel：创建通道及其缓冲区
	*/

	/* new  分配内存并返回指针，指向类型的零值 */
	var p *int = new(int)
	logger.Printf("p=%d\n", *p) //0

	type User struct {
		FamilyMember []string
		Age          int
	}
	var u *User = new(User)
	logger.Printf("u=%v\n", u)

	/* 集合操作函数 */
	/* append用于切片追加元素 */
	s5 := append(s4, []int{1, 2, 3}...)
	logger.Printf("append s4 get s5=%v\n", s5)

	/* copy函数用于复制切片内容 */
	s6 := []int{100, 200, 300}
	s7 := make([]int, 2, 5)
	s7[0] = 1
	copy(s7, s6)
	logger.Printf("copy s6 to s7 get s7=%v\n", s7)
	logger.Printf("len of s7=%d,cap of s7=%d\n", len(s7), cap(s7)) //2,5
	s8 := s7[1:]
	logger.Printf("len of s8=%d,cap of s8=%d\n", len(s8), cap(s8)) //1,4

	/* delete函数用于删除map中的元素 */
	m3 := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	if v, ok := m3[4]; ok { //判断某一个键值对是否存在
		logger.Printf("四月是%s\n", v)
	}
	delete(m3, 4) //删除键值对
	logger.Printf("删除键值对4后,m3=%v\n", m3)

	/* 并发控制函数 */
	/* close函数 用于关闭channel */
	ch4 := make(chan bool)
	defer close(ch4) //关闭channel
	go func() {
		ch4 <- true
	}()
	for {
		if b, ok := <-ch4; ok {
			logger.Printf("ch4 receive %v\n", b)
			break
		}
	}

	/* 错误处理函数 */
	/* recover函数 捕获panic值并恢复正常执行 recover只在defer函数中有效*/
	defer func() {
		if err := recover(); err != nil {
			if msg, ok := err.(string); ok {
				logger.Printf("recover from panic:%s\n", msg)
			}
		}
	}()

	/* panic函数 用于触发运行时panic */
	defer panic("panic error")

	/* 复数函数 */
	/* complex函数 用于创建复数 */
	var cm1 complex128 = complex(2, 3)
	logger.Printf("c1=%v\n", cm1)
	logger.Printf("real of cm1=%v,imag of cm1=%v\n", real(cm1), imag(cm1))

	/* 新增函数 */
	/* clear函数 用于清空map、slice中的所有元素 */
	var s10 []int = []int{1, 2, 3, 4, 5}
	var m10 = map[string]string{
		"1": "one",
		"2": "two",
	}
	clear(s10)
	clear(m10)
	logger.Printf("clear set elements:s10=%v,m10=%v\n", s10, m10) //s10=[0 0 0 0 0],m10=map[]

	/* max min函数 得到极值 */
	maxValue := max(10, 20, 39)
	minValue := min("apple", "dream", "lady")
	logger.Printf("maxValue=%d,minValue=%s\n", maxValue, minValue)

}
