package main

import (
	"fmt"
)

func aa(x int)(func(int)int){
	return func(y int) int {return x+y}
}

func main(){
	aa5:=aa(5)
	aa6:=aa(6)
	aa7:=aa(7)
	aa4:=aa(4)
	fmt.Println(aa5(aa6(1)))
	fmt.Println(aa7(aa4(2)))
}
package main

import “fmt”

func main() {
	callback(1, Add)
}
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

var a [4]int //元素自动初始化为零
b := [4]int{2, 5} //未提供初始值的元素自动初始化为零 
c := [4]int{5, 3: 10} //可指定索引位置初始化 
d := [...]int{1, 2, 3} //编译器按初始化值数量确定数组长度 
e := [...]int{10, 3: 100}//支持索引初始化，但数组长度与此有关 
fmt.Println(a) //可直接输出素组各个元素，比C方便



defer fmt.Println("Program ended")
srcFile,err:=os.Open(srcName)
defer srcFile.Close()//关闭文件句柄，结构清晰，避免忘记关闭