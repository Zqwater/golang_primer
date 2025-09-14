package imptype

import (
	array "golang_primer/chapters/imptype/1_array"
	structtype "golang_primer/chapters/imptype/2_struct"
	pointer "golang_primer/chapters/imptype/3_pointer"
)

func ImpType() {
	/* 值类型 */

	//数组类型
	array.Array()

	//结构体类型
	structtype.StructType()
	structtype.EmptyStruct()

	/* 引用类型 */

	//指针类型
	pointer.Pointer()

	//切片slice

	//映射map

	//通道channel

}
