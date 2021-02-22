package main

import (
	"fmt"
	"sort"
)
var strscore = map[string]int{"Frank": 100, "Joyce":98, "Steven":0, "Phyllis":60}
func main(){
	keys := make([]string, len(strscore))
	i:=0
	for k,_:=range strscore {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i,j int)bool{return keys[i] < keys[j]})
	for _,k:=range keys {
		fmt.Printf("(%v,%v)", k, strscore[k])
	}
}