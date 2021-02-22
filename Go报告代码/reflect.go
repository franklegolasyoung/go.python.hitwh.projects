package main

import (
	"fmt"
	"reflect"
)

type Phone struct {
	Name string
	Year int
}

func SetValue(o interface{}) {
	v:=reflect.ValueOf(o)
	if v.Kind() !=reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("Can't set")
		return 
	} else {
		v=v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Int:
			v.Field(i).SetInt(2018)
		case reflect.String:
			v.Field(i).SetString("iPhone X")
		}
	}
}

func main(){
	phone:=Phone{"iPhone 4", 2010}
	SetValue(&phone)
	fmt.Println(phone)
}