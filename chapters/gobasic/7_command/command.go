package command

import (
	"flag"
	"golang_primer/golog"
)

/* golang命令 */
func Command() {
	/* go build命令 */
	//go build -o app main.go 指定输出文件名
	//go build -ldflags "-s -w" -trimpath -tags jsoniter main.go
	/*
		-o：指定输出文件名
		-ldflags：链接器标志（-s 移除符号表，-w 移除调试信息）
		-trimpath：移除文件系统路径信息（增强可复现性）
		-tags：条件编译标签
	*/

	logger := golog.Logger
	url := `https://blog.629622.xyz/2025/07/03/golang%E5%91%BD%E4%BB%A4%E8%AF%A6%E8%A7%A3/`
	logger.Printf("参考文章:%s\n", url)

	/* go run 命令 */
	//flag包来获取参数
	var name string //通过flag.TypeVar来绑定参数
	var married bool
	age := flag.Int("age", 0, "你的年龄") //flag.Type 返回指针
	port := flag.Int("port", 8080, "端口号")

	flag.StringVar(&name, "name", "static", "你的名字") //static是默认值 中文是说明
	flag.BoolVar(&married, "married", false, "是否结婚")

	flag.Parse()
	logger.Printf("从控制台获取flag信息,名字:%s,年龄:%d,是否已婚:%T,程序运行端口:%d\n", name, *age, married, *port)

	//go run -exec "sudo" main.go 使用sudo 权限运行
	//go run -race main.go启动竞态检测（调试并发问题）
	//go run -gomodcache=/tmp/cache main.go(临时gomodcache,避免污染全局缓存)
	/*
	 -a: 强制重新编译所有依赖
	 -n: 打印但不执行编译命令
	 -x: 显示完整执行过程
	 -work: 保留临时工作目录
	*/

	/* go test命令 */
	//go test -v -run TestXX test/XX_test.go
	/*
		-v：详细模式输出
		-cover：测试覆盖率
		-coverprofile：指定覆盖率文件输出
		-shuffle：随机化测试顺序
		-run：只运行匹配的测试函数，如 go test -run TestFunctionName
		-timeout：设置测试超时时间
	*/

}
