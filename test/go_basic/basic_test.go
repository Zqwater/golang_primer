package gobasic_test

import (
	"testing"
)

/* golang示例 */
func TestReal(t *testing.T) {

	/* langguage_basic示例 */

	//变量命名规则
	var userName string = "static"
	//var 1_job string = "dev" //错误，不能以数字开头
	var _name = "tianshen" //可以以下划线开头
	//type := "newtype" //不能以关键字为变量名字
	//const type string = "Hello Go" //不能以关键字为常量名字

	//常用关键字和保留字的使用例子
	defer func() { //recover捕获异常
		if err := recover(); err != nil {
			t.Logf("发生了错误:%v", err)
		}
	}()

	defer func(strs ...string) { //defer延迟函数
		for k, v := range strs {
			t.Logf("第%d个参数的值为:%s", k, v)
		}
		//panic主动触发
		panic("这是一个主动的panic")
	}(userName, _name)

	a := new(*int)                      //new函数
	t.Logf("a的类型为:%T,a的值是%#v\n", a, *a) //**int类型

	/* inner_type和innner_fun示例 */
	

}
