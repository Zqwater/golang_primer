package structtype

import (
	"basicGo/basicLesson/golog"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

var (
	logger = golog.Logger
)

/* golang结构体类型 */
func StructType() {

	/* 自定义类型和类型别名 */
	//自定义类型可以通过内置类型或struct来定义
	type MyString string
	type MyChan chan bool
	type User struct {
		Name string
		Age  int
	}
	logger.Printf("自定义类型:MyString=%T,MyChan=%T,User=%T\n", //通过类型转换初始化快速声明
		MyString(""), MyChan(make(chan bool)), User{})

	//类型别名 相同类型更换一个名字
	type NewString = MyString
	type NewUser = User
	logger.Printf("类型别名:NewString=%T,NewUser=%T\n", NewString(""), NewUser{})

	/* 结构体 */
	/*
	 1.与java通过类和继承实现oop不同，使用结构体和接口的组合来实现oop
	 2.结构体内存布局，内存对齐
	 3.匿名结构体，结构体的嵌套
	 3.结构体字段的访问，匿名字段（只有类型，没有字段），tags标签使用
	 4.结构体方法的定义和调用
	*/

	//结构体的定义
	type Man struct {
		Name    string `json:"name"` //json tags用于json序列化
		Age     int    `json:"age"`
		Married bool   `json:"married"`
	}

	type Address struct {
		Province string `json:"province"`
		City     string `json:"city"`
	}

	//结构体的实例化（实例化就是分配内存地址）
	var m1 Man //基本实例化
	m1.Age = 18
	m1.Name = "static" //m1.married取零值
	logger.Printf("Man基本实例化:%+v\n", m1)

	var m2 = new(Man) //使用new进行实例化 返回的是指向零值的指针
	m2.Name = "tianshen"
	m2.Age = 22 //m2.married取零值
	logger.Printf("Man使用new实例化:%+v\n", m1)

	var m3 = &Man{} //取结构体指针实例化
	m3.Name = "water"
	m3.Age = 26 //m3.married取零值
	logger.Printf("Man使用取结构体指针实例化:%+v\n", m3)

	//结构体初始化
	var a1 Address
	a1.Province = "sichuang" //基本初始化
	a1.City = "chengdu"

	a2 := new(Address)
	a2.Province = "shanghai"
	a2.City = "shanghai"

	a3 := &Address{}
	a3.Province = "beijing"
	a3.City = "beijing"

	a4 := Address{ //键值对初始化
		Province: "guangdong",
		City:     "shenzhen",
	}

	a5 := &Address{
		Province: "guangxi",
		City:     "nanning",
	}

	a6 := Address{
		"taiwan",
		"gaoxiong",
	}

	logger.Printf("结构体初始化:a1=%+v,a2=%+v,a3=%+v,a4=%+v,a5=%+v,a6=%+v\n",
		a1, a2, a3, a4, a5, a6)

	//匿名结构体
	as1 := struct { //匿名结构体使用
		Length int
		time   time.Duration
	}{
		Length: 17,
		time:   30 * time.Minute,
	}
	logger.Printf("匿名结构体:%+v\n", as1)

	as2 := struct {
		Address
	}{
		Address{
			"tokyo",
			"daban",
		},
	}

	logger.Printf("匿名结构体嵌套as2=:%+v\n", as2)

	//匿名字段
	type Teacher struct { //嵌套匿名字段的结构体实现类似java的继承
		City    string
		int     // 只有类型，没有字段的匿名字段
		Address //匿名字段嵌入
	}

	te1 := Teacher{
		City: "shanghai",
		int:  30,
		Address: Address{
			"china",
			"shanghai",
		},
	}
	logger.Printf("匿名字段外部city=%s,匿名字段内部city=%s\n", te1.City, te1.Address.City)

	//结构体tags
	type Information struct { //json序列化
		Age  int    `json:"age" db:"user_age" xml:"user_id"`
		Name string `json:"name"`
	}

	type Boy struct {
		Dream  string `json:"dream"`
		Vargin bool   `json:"vargin"`
		*Information
	}

	boy1 := &Boy{
		Dream:  "the world beauty",
		Vargin: true,
		Information: &Information{
			Age:  18,
			Name: "static",
		},
	}

	boy2 := &Boy{
		Dream:  "the world awful",
		Vargin: false,
		Information: &Information{
			Age:  22,
			Name: "water",
		},
	}

	boys := make([]Boy, 0)
	boys = append(boys, *boy1, *boy2) //添加元素

	go func() {
		if err := recover(); err != nil {
			logger.Printf("catch panic:%v\n", err)
		}
	}()

	bs, err := json.Marshal(boys)
	if err != nil {
		panic(err)
	}

	logger.Printf("json序列化boys:%s\n", string(bs))

	boy := reflect.TypeOf(Boy{}) //使用反射读取标签
	filed, _ := boy.FieldByName("Age")
	logger.Printf("反射读取标签json:%v\n", filed.Tag.Get("json"))
	logger.Printf("反射读取标签db:%v\n", filed.Tag.Get("db"))

	//内存布局
	/*
	 1.结构体在内存中是连续分配的字段序列。Go编译器会根据平台自动进行内存对齐以优化访问速度
	 2.字段顺序影响内存占用 大前小后原则
	 3. unsafe.Offsetof可以查看内存细节
	*/

	type Example struct {
		a bool  //1字节，扩充到4
		b int32 //4
		c byte  //1字节，扩充到8
		d int64 //8
	}
	var example Example

	logger.Printf("Example内存对齐细节:%v\t%v\t%v\t%v\n",
		unsafe.Offsetof(example.a), unsafe.Offsetof(example.b), unsafe.Offsetof(example.c), unsafe.Offsetof(example.d))

	//结构体方法
	/*
	 值类型和指针接受类型都拥有全部的方法,编译器提供的语法糖，自动解引用和取引用
	 1.值类型方法可以通过指针调用，编译器会自动解引用
	 2.指针类型方法也可以通过值类型调用，自动取地址
	 3.关键限制：值类型接受者调用指针方法时要，值必须是可以寻址的（变量或者结构体字段）
	*/
	StructMethod()

	//结构体和接口嵌套实现oop:组合实现继承，接口实现多态
	StructOOP()

}

/* structmethod */
type Class struct {
	ClassMem  int    //班级人数
	ClassName string //班级名称
}

func (c Class) GetMem() int {
	return c.ClassMem
}

func (c *Class) SetClassName(name string) bool {
	if name == "" {
		return false
	}
	c.ClassName = name
	return true
}

type Stu struct {
	Name  string
	Age   int
	Class //内嵌匿名字段 实现继承
}

func (s Stu) GetStuAge() int {
	return s.Age
}

func (s *Stu) SetStuAge(age int) bool {
	if age < 0 {
		return false
	}
	s.Age = age
	return true
}

func MakeStu(name string, age int, className string, classMem int) *Stu {
	return &Stu{
		Name: name,
		Age:  age,
		Class: Class{
			ClassMem:  classMem,
			ClassName: className,
		},
	}
}

func StructMethod() {

	stu1 := MakeStu("static", 22, "1003", 35)

	stu2 := Stu{
		Name: "water",
		Age:  28,
		Class: Class{
			ClassMem:  20,
			ClassName: "2004",
		},
	}

	logger.Printf("stu1=%v,stu2=%v\n", stu1, stu2)
}

/* structoop */
type Abler interface {
	Run(name string)
	TimesCount() int
}

type Father struct {
	Time int
	Name string
}

func (f Father) Run(name string) {
	str := fmt.Sprintf("%s run %s\n", f.Name, name)
	logger.Printf("father %s\n", str)
}

func (f *Father) TimesCount() int {
	f.Time++
	return f.Time
}

type Son struct {
	Common bool
	Father //嵌入匿名字段
}

func (s Son) Run(name string) {
	logger.Printf("son run %s\n", name)
}

func (s Son) TimesCount() int {
	logger.Printf("son has come with:%t\n", s.Common)
	return 0
}

// 实现继承和多态
func StructOOP() {

	var able Abler //声明接口变量

	f1 := Father{
		Time: 0,
		Name: "father1",
	}

	// able = f1 //wrong: 接口方法规则
	able = &f1

	able.Run("go")

	s1 := Son{
		Common: true,
		Father: Father{
			Time: 10,
			Name: "godfather",
		},
	}
	able = s1 //多态
	able.Run("python")
	s2 := &s1

	able = s2 //true 自动解引用
	able.Run("java")

}
