package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
type Point struct{
	x,y float64
}
func (s Point)Length() float64{
	return math.Sqrt(s.x*s.x+s.y*s.y)
}
func main() {
	Rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var (
		a Point
		m int = 0
		n int = 1000000000
		r float64 = 0.0
	)
	for i := 0; i < n; i++ {
		a.x = Rand.Float64()
		a.y = Rand.Float64()
		if a.Length()<=1 {
			m++
		}
	}
	r = 4*float64(m)/float64(n)
	fmt.Println("π =",r)
}