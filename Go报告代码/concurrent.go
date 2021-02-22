package main

import (
	"fmt"
	"time"
)

func main(){
	go taskone()//启动任务1
	go tasktwo()//启动任务2
	time.Sleep(time.Second * 6)
}
func taskone() { 
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Second)
	}
}
func tasktwo() {
	for i := 0; i < 5; i++ {
		fmt.Println("世界")
		time.Sleep(time.Second)
	}
}