package main
import (
	"fmt"
	"math"
)
func run(n int) float64 {
	ch := make(chan float64, 100)
	for i := 0; i < n; i++ {
		go pi(ch, float64(i))
	}
	var f float64
	f = 0
	for i := 0; i < n; i++ {
		f += <-ch
	}
	return f
}
func pi(ch chan float64, f float64) {
	ch <- 4 * math.Pow(-1, f) / (2*f + 1)
}
func main() {
	var n int
	fmt.Println("输入n:")
	fmt.Scanln(&n)
	fmt.Println(run(n))
}