package main

import (
	"fmt"
)

func main(){
	a:=10
	if a>0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	//精简化
	if a=-10;a>0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	//switch
	switch a {
	case 1:
		fmt.Println("1")
	case 2,10:
		fmt.Println("10")
	default:
		fmt.Println("0")
	}
}