package main

import (
	"fmt"
)

func main(){
	//定义单个常量
	const a1 int = 1
	const b1 = 'A'
	const (
		text1 = "qwert"
		len1 = len(text1)
	)
	fmt.Println(a1,b1,text1,len1)
	//同时定义多个常量
	const a2, b2 = 2, 'B'
	const (
		text2, len2 = "asdfg", len(text1)
		//不可len(text2)，因为text2未定义
	)
	fmt.Println(a2,b2,text2,len2)
	//常量计数器
	const (
		a=iota//a的值为0
		b //d的值为1
		c //d的值为2
	)
	const (
		d = iota
		e //f的值为1
	)
	fmt.Println(a,b,c,d,e)
	//计算机存储单位的枚举
	const (
		_ = iota
		KB float64 = 1 << (iota * 10)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(KB,MB,GB,TB)
}