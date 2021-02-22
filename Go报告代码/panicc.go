package main

import (
	"fmt"
)

func main(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("程序异常退出:",err)
		} else {
			fmt.Println("程序正常退出")
		}
	}()
	f(101)
}
func f(a int) {
	if a > 100 {
		panic("参数值超出范围")
	} else {
		fmt.Println("函数f成功调用")
	}
}