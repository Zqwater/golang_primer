package gobasic

import (
	langfeature "golang_primer/chapters/gobasic/1_langfeature"
	languagebasic "golang_primer/chapters/gobasic/2_languagebasic"
	innertypefun "golang_primer/chapters/gobasic/3_innertypefun"
	initwithmain "golang_primer/chapters/gobasic/4_initwithmain"
	calculatesignal "golang_primer/chapters/gobasic/5_calculatesignal"
	underline "golang_primer/chapters/gobasic/6_underline"
	command "golang_primer/chapters/gobasic/7_command"
	varwithconst "golang_primer/chapters/gobasic/8_varwithconst"
)

/* goalng基础入口函数 */
func GoBasicRun() {

	//golang语言的主要特横
	langfeature.GoFeatures()

	//golang语言的基础语法
	languagebasic.LanguageBasic()

	//golang内置类型和内置函数
	innertypefun.InnerType()
	innertypefun.InnerFunc()

	//golang main和init函数
	initwithmain.MainFun()
	initwithmain.InitFun()

	//golang运算符
	calculatesignal.CalculateSignal()

	//golang下划线
	underline.UnderLine()

	//golang命令
	command.Command()

	//golang变量和常量
	varwithconst.Var()
	varwithconst.Const()

}
