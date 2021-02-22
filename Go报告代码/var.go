package main

import (
	"fmt"
	"math"
)

func main(){
	var r float64
	r=3
	c:=2*math.Pi*r
	var s float64 = math.Pi*r*r
	var a int = 4
	b:=3
	var name = "circle"
	name1:="square"
	fmt.Println(r,c,s,name,a,b,name1)
}