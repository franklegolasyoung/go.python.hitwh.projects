package main
import (
	"fmt"
)
func f(a int) {
	if a>100{
		panic("exceed")
	} else {
		fmt.Println("succeed")
	}
}
func main() {
	defer func() {
		if err:=recover(); err!=nil{
			fmt.Println("error",err)
		} else {
			fmt.Println("exit")
		}
	}()
	f(100)
}