package main

import (
	"golang_primer/chapters"
	"golang_primer/golog"
)

func init() {
	//创建日志文件以输出信息
	golog.InitLog()
}

func main() {
	//程序入口
	chapters.Run()
}
