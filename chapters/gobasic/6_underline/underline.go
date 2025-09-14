package underline

import (
	"database/sql"
	"errors"
	"golang_primer/golog"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDB  *sql.DB
	mysqlDns = "static:295295@(localhost:3306)/test?charset=utf8mb4&parseTime=True"
)

type Writer interface {
	Write([]byte) (int, error)
}

type SelfWriter struct {
	file *os.File
}

func (w *SelfWriter) Write(p []byte) (int, error) {

	//实现写入逻辑
	n, err := w.file.Write(p)
	return n, err
}

/* 下划线用法 */
func UnderLine() {
	var err error
	logger := golog.Logger
	/* 1.忽略不需要使用的值 */
	//自定义函数 获取slice中的中间值
	var s []int
	selfFun := func(s []int) (int, error) {
		if s == nil {
			return 0, errors.New("slice is nil")
		}
		for i := 0; i < len(s); i++ { //冒泡排序
			for j := 0; j < len(s)-i-1; j++ {
				if s[j] > s[j+1] {
					s[j], s[j+1] = s[j+1], s[j]
				}
			}
		}
		return s[len(s)/2], nil
	}
	_, err = selfFun(s)
	if err != nil {
		logger.Printf("下划线忽略获取值:%v\n", err)
	}

	/* 调用其他包的init函数 */
	//使用mysql的包 注册mysql驱动
	initMysqlConn := func(dns string) (*sql.DB, error) {
		db, err := sql.Open("mysql", dns)
		if err != nil {
			return nil, errors.New("parse dns failed")
		}
		err = db.Ping()
		if err != nil {
			return nil, errors.New("connect mysql failed")
		}
		logger.Printf("connect to mysqldb success!")

		return db, err
	}

	MysqlDB, err = initMysqlConn(mysqlDns)
	if err != nil {
		logger.Printf("connect to mysql error:%v\n", err)
	}

	/* 编译期接口检查 */
	var _ Writer = (*SelfWriter)(nil) //编译期检查

	/* 结构体中的占位字段 */
	//用途：内存对齐、预留扩展位、防止结构体字面量初始化时遗漏字段
	type ServerConfig struct {
		Host    string
		Port    int
		_       [16]byte //预留扩展空间
		TimeOut time.Duration
	}

	var sc ServerConfig
	logger.Printf("_作为结构体中的占位字段:%v\n", sc)

	/* 字面量分割符 不影响原本的数值 */
	const Gigabyte = 1_000_000
	const BinaryGB = 1_073_741_824
	const PhoneNum = 173_8056_3622
	logger.Printf("_作为字面量分割符Gigbyte:%v 二进制GB:%v 手机号:%v\n", Gigabyte, BinaryGB, PhoneNum)

}
