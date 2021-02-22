package main

import (
	"fmt"
)

type IPhone struct{
	Name string
	Year int
}
type Tester interface{}
func main() {
	iPhone := IPhone{"iPhone X", 2018}
	it :=make([]Tester, 4)
	it[0]=2008
	it[1]="Apple"
	it[2]=iPhone
	it[3]=true
	for i, e:=range it {
		switch val := e.(type) {
		case int:
			fmt.Printf("it[%d] type is int, val=%d\n", i, val)
		case string:
			fmt.Printf("it[%d] type is int, val=%s\n", i, val)
		case IPhone:
			fmt.Printf("it[%d] type is int, val=%v\n", i, val)
		case bool:
			fmt.Printf("it[%d] type is int, val=%v\n", i, val)
		default:
			fmt.Printf("Unknown type of it[", i,"]")
		}
	}
}