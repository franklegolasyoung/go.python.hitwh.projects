package main

import (
	"fmt"
)

func main(){
	for x:=0;x<3;x++ {
		fmt.Println(x)
	}//普通for循环
	x:=0
	for x<3 {
		fmt.Println(x)
		x++
	}//相当于C中的while(x<3) {……}
	for {
		fmt.Println(x)
		x--
		if x<0 {
			break
		}
	}//相当于C中的while(true)或do...while(x>=0);
	K:=[]int{100,101,102}
	for i,v:=range K {
		fmt.Println(i,":",v)
	}//for...range可同时返回集合中数据的索引,以及数据的值
}