package innertypefun

import (
	"basicGo/basicLesson/golog"
)

/* 内置类型 */
func InnerType() {

	logger := golog.Logger

	/* bool类型 */
	//1.声明
	var b1 bool
	var b2 bool = false
	var b3 = true
	b4 := false
	//2.使用
	logger.Printf("bool类型:b1=%t,b2=%t,b3=%t,b4=%t\n", b1, b2, b3, b4)

	/* 数值类型 整型(无符号，有符号) 浮点型 */
	//1.声明
	var i1 int8
	var i2 int16 = -32768
	var i3 int32 = -2147483648
	var i4 rune = 2147483647
	var i5 int64 = -9223372036854775808

	var u1 uint8
	var u2 byte = 'A' //ASCII码对应65-90(A-Z) 97-122(a-z)
	var u3 uint16 = 65535
	var u4 uint32 = 4294967295
	var u5 uint64 = 18446744073709551615
	var u6 uintptr = uintptr(u2)

	var f1 float32 = 3.1415926         //6位精度
	var f2 float64 = 3.141592653589793 //15位精度

	//2.使用
	logger.Printf("整型带符号:i1=%d,i2=%d,i3=%d,i4=%d,i5=%d\n", i1, i2, i3, i4, i5)
	logger.Printf("整型无符号:u1=%d,u2=%d,u3=%d,u4=%d,u5=%d,u6=%d\n", u1, u2, u3, u4, u5, u6)
	logger.Printf("浮点型:f1=%f,f2=%f\n", f1, f2)

	/* 字符串类型 */
	//1.声明
	var s1 string
	var s2 string = "golang"
	var s3 = "gin"
	s4 := "go-zero"
	var s5 string = `hello 世界` //原生字符串，不进行转义

	//2.使用
	logger.Printf("字符串类型:s1=%s,s2=%s,s3=%s,s4=%s,s5=%s\n", s1, s2, s3, s4, s5)
	logger.Printf("len of s5 = %d\n", len(s5))

	/* complex复数类型 */
	//1.声明
	var c1 complex64 = 1.1 + 2i
	var c2 complex128 = 2.718 + 3i
	c3 := 3.14 + 2.2i
	c4 := complex(5.2, 6.6)
	//2.使用
	logger.Printf("复数类型:c1=%v,c2=%v,c3=%v\n", c1, c2, c3)
	logger.Printf("real of c4 = %f,imag of c4 = %f\n", real(c4), imag(c4))

	/* 数组  结构体 具体参考其他章节，此处仅作简单的展示*/
	//数组
	var a1 [3]int = [3]int{1, 3, 5}
	var a2 = [5]float32{1.2, 2.4, 3.6, 4.8, 5.10}
	a3 := [2]string{"hello", "go语言"}

	logger.Printf("数组类型:a1=%T,a2=%T,a3=%T\n", a1, a2, a3)

	//结构体
	type User struct {
		Name string
		Age  int
	}

	var user1 User
	var user2 = User{
		Name: "张三",
		Age:  18,
	}

	user3 := User{ //匿名结构体
		Name: "李四",
		Age:  20,
	}

	user3.Age = 21 //匿名结构体可以直接使用字段名访问字段

	var user4 = &User{ //结构体匿名字段
		"中中",
		22,
	}
	user4.Name = "小中中" //结构体匿名字段可以直接使用和修改字段值

	logger.Printf("user1=%v,user2=%v,user3=%v\n", user1, user2, user3)
	logger.Printf("address of user4 = %p,user4.Name = %s\n", user4, user4.Name)

	/* 切片 映射 通道 具体参考其他章节，此处仅作简单的展示*/
	//切片 声明 内置函数使用 make | len cap copy append
	var sl1 []int = []int{1, 2, 3} //字面量声明
	var sl2 = []float32{2.2, 3.3, 4.5}
	var sl3 []interface{}
	sl4 := []string{"hello", "dream", "skt"}

	var si1 []int
	var si2 = a2[:len(a2)-1]
	si3 := a2[:] //数组切片
	si4 := a1[1:2:3]

	var ss1 []int = make([]int, 1, 3) //使用make
	var ss2 = make([]int, 2, 3)
	ss3 := make([]int, 3)

	logger.Printf("sl1=%v,sl2=%v,sl3=%v,sl4=%v\n", sl1, sl2, sl3, sl4)
	logger.Printf("si1=%v,si2=%v,si3=%v,si4=%v\n", si1, si2, si3, si4)
	logger.Printf("ss1=%v,ss2=%v,ss3=%v\n", ss1, ss2, ss3)

	//映射 声明 函数使用 make | len cap delete
	var m1 map[int]int //字面量声明
	var m2 map[int]float32 = map[int]float32{
		1: 1.1,
		2: 2.2,
	}
	var m3 = map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	m4 := map[string]string{
		"skt": "faker",
		"g2":  "caps",
		"rng": "uzi",
	}

	var mm1 map[int]int = make(map[int]int, 2) //使用make
	var mm2 = make(map[int]string)
	mm3 := make(map[string]string, 3)

	logger.Printf("m1=%v,m2=%v,m3=%v,m4=%v\n", m1, m2, m3, m4)
	logger.Printf("mm1=%v,mm2=%v,mm3=%v\n", mm1, mm2, mm3)

	//chan通道声明 函数 make | close len
	var ch1 chan int = make(chan int, 2) //使用make
	ch2 := make(chan string)             //阻塞通道

	var ch3 chan<- bool = make(chan<- bool, 10) //单向通道 只能发送
	ch4 := make(<-chan bool)                    //单向通道 只能接收

	//chan通道使用 内置函数 close len cap
	defer close(ch1)

	go func() {
		select {
		case v := <-ch2:
			logger.Printf("receive data from ch2:%v\n", v)
		default:
			logger.Println("no data received")
		}
	}()
	ch2 <- "hello"

	logger.Printf("单向通道类型ch3:%T,ch4:%T\n", ch3, ch4)

}
