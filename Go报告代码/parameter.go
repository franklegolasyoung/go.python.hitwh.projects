package main

import (
	"fmt"
)

func main(){
	f1(1,8,0,4,0,0,7,1,5)
}

func f1(args ...interface{}) {
	f2(args...)
	f2(args[5:]...)
}

func f2(args ...interface{}) {
	fmt.Println(args)
}