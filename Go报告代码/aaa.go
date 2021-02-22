package main

import (
	"fmt"
)

func change(y []int, a int, n int) []int {
	x := make([]int, 0, n)
	for _, i := range y {
		if i != a {
			x = append(x, i)
		}
	}
	return x
}

func main() {
	var a [3][3]int
	x0 := make([]int, 0, 9)
	for i := 1; i < 10; i++ {
		x0 = append(x0, i)
	}
	for _, a[0][0] = range x0 {
		x1 :=change(x0, a[0][0], 8)
		for _, a[0][1] = range x1 {
			x2 :=change(x1, a[0][1], 7)
			for _, a[0][2] = range x2 {
				x3 :=change(x2, a[0][2], 6)
				for _, a[1][0] = range x3 {
					x4 :=change(x3, a[1][0], 5)
					for _, a[1][1] = range x4 {
						x5 :=change(x4, a[1][1], 4)
						for _, a[1][2] = range x5 {
							x6 :=change(x5, a[1][2], 3)
							for _, a[2][0] = range x6 {
								x7 :=change(x6, a[2][0], 2)
								for _, a[2][1] = range x7 {
									x8 :=change(x7, a[2][1], 1)
									a[2][2] = x8[0]
									b1 := a[0][0] + a[1][0] + a[2][0]
									b2 := a[0][1] + a[1][1] + a[2][1]
									b3 := a[0][2] + a[1][2] + a[2][2]
									b4 := a[0][0] + a[0][1] + a[0][2]
									b5 := a[1][0] + a[1][1] + a[1][2]
									b6 := a[2][0] + a[2][1] + a[2][2]
									if b1 == 14 && b2 == 15 && b3 == 16 && b4 == 8 && b5 == 18 && b6 == 19 {
										fmt.Println(a)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}