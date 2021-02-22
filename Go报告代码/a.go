package main

import (
	"fmt"
)

func main(){
	//变量
	c:=3
	var r int
	a:=c+r
	byte:='a'
	const d='A'
	const(
		l=iota
		j
	)//常量计数器//枚举效果
	defer fmt.Println("a")//延迟调用
	//if//else不能单独成行,大括号不能省略
	if c>3 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
	//switch//不需加break
	switch byte {
	case'a','b':
		r = 1
	default:
		r = 0
	}
	//for循环+range
	//range右边的表达式必须是array，slice，string，map或是指向array的指针，也可以是channel
	//数组
	x:=[]int{100,101,103,108}
	for i,v:=range x{
		fmt.Println(i,":",v)
	}
	//三种输入方式
	var b int
	fmt.Scan(&a)//数据间隔可以使用空格和回车
	//Scan接收字符型变量时只接收整型
	//并通过ASCII码值转换为字符型
	fmt.Scanf("%d", &a)//数据间隔只能使用回车
	//%d只能接收整型，不能接收字符型
	//所以在输入字符型变量时应该使用fmt.Scanf("%c",&a)
	fmt.Scanln(&a)
	fmt.Scan(&b)
	//函数
	//1.匿名函数
	sum:=func(a,b int) int{ 
		return a+b
 	}(2,3)//声明的同时直接调用匿名函数 
	fmt.Println(sum)
	//2.swap
	fmt.Println(swap(a, b))
	//3.闭包
	aa5:=aa(5)
	aa6:=aa(6)
	fmt.Println(aa5(aa6(1)))
}
func swap(a, b int) (int, int) {
	return b, a
}
func divide(a, b int) (quotient , remainder int) {
	quotient = a / b
	remainder = a % b
	return
}
func aa(x int)(func(int)int){
	return func(y int) int {return x+y}
}