package main

import (
	"fmt"
	"runtime"
)
type vector []int64
var result int64
func (v vector) DoAll() {
	var NCPU int = runtime.NumCPU()
	c:=make(chan int64,NCPU-1)
	for i:=0; i<NCPU-1; i++{
		go v.Dosome(i*len(v)/(NCPU-1), (i+1)*len(v)/(NCPU-1), c)
	}
	for i:=0;i<NCPU-1;i++{
		value := <-c
		result += value
	}
}
func (v vector) Dosome(i, n int, c chan int64){
	for ; i<n-1; i++{
		v[i+1]+=v[i]
	}
	c <- v[i]
}
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	v:=make(vector, 100000000,100000000)
	for i:=0;i<100000000;i+=2{
		v[i]=int64(i+1)
	}
	v.DoAll()
	fmt.Println(result)
}