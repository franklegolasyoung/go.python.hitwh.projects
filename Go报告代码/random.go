package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	a:=[10]int{}
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		a[i]=rand.Intn(100)
		k:=1
		for j:=i-1;j>=0;j--{
			if a[i]==a[j]{
				k=0
				i--
				break
			}
		}
		if k==1{
			fmt.Println(a[i])
		}
	}
}