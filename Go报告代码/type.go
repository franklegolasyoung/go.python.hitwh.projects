package main
import (
	"fmt"
	"math"
)
//对象
type Point struct{ 
	x,y float64
}
//方法
func(self Point)Length() float64{ 
	return math.Sqrt(self.x*self.x+self.y*self.y) 
}//距离
func(self *Point)Scale(factor float64){ 
	self.x=self.x*factor
	self.y=self.y*factor
}//坐标变换
//接口
func(p Point) Print() {
	fmt.Printf("%+v\n", p) 
}
//继承
type Circle struct{
	Point//匿名成员 
	r int
}
func main(){
	p := Point{3,4} 
	d:=p.Length()//=>5
	fmt.Println(d)
	p.Scale(2) 
	e:=p.Length()//=>10
	fmt.Println(e)
	p.Print()
	var C Circle
	C.r=3
	C.Point=p
	C.Point.x=C.Point.y
	C.Print()
}