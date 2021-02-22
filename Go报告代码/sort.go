package main

import (
	"fmt"
	"sort"
)
type IPhone struct {
	Name string
	Year int
}
func main(){
	//简单数值型切片
	score := []int{60,70,50,40,80,90}
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	fmt.Println(score)
	//升序
	sort.Slice(score, func(i,j int) bool {return score[i]<score[j]})
	fmt.Println(score)
	//结构体排序
	iPhones := []IPhone{{"iPhone 4",2010},{"iPhone X",2017},{"iPhone 5",2012},{"iPhone 6 Plus",2014}}
	sort.Slice(iPhones,func(i,j int)bool{return iPhones[i].Year < iPhones[j].Year})
	fmt.Println(iPhones)
}