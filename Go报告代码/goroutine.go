package main

import (
	"fmt"
	"runtime"
)

func main(){
	names := []string{"Frank","Steven","Joyce","Phyllis"}
	for _, name :=range names {
		go func(who string) {
			fmt.Printf("Hello,%s.\n", who)
		}(name)
		runtime.Gosched()
	}
}