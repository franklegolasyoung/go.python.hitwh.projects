package main 
import "fmt"
func main(){/*
	//结构体数组
	type user struct{ 
		name string
		age byte
	}//定义user结构或类 
	d := []user{{"Frank", 20},{"Jinmi", 18}}//定义结构体数组，有2个元素 
	fmt.Printf("%v\n%+v\n%#v\n",d,d,d)
	//多维数组
	a := [2][2]int{{1,2}, {3,4}} 
	b := [][2]int{
		{10,20}, 
		3:{30,40}, 
	} 
	c := [][2][2]int{
 		{
			{1,2},
			{3,4},
		},
		{
			{10, 20},
			{30, 40},
		},
	}
	fmt.Printf("%v\n%v\n%v\n",a,b,c)
	//数组和指针
	//数组名不代表数组的首地址
	x, y := 10, 20 
 	s := []*int{&x,&y}		//指针数组，各个元素为指针 
	p := &s                //指向数组地址的指针 
	fmt.Printf("%T,%v\n", s, s) 
	fmt.Printf("%T,%p,%v,%v,%p\n", p, p, p, *p, &p)
	//
	
	k := make([]int, 0, 3) //创建长度为0,容量为3的切片 
	for i := 0; i < 8; i++ { 
		k = append(k, i) //追加数据,超出容量限制时,容量自动翻倍 
	} 
	fmt.Println(k, len(k), cap(k))
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}//定义数组x有10个元素 
	s := x[2:5]//创建切片,指向数组的2,3,4号元素，容量为len(x)-2 
	for i, v := range s { 
		fmt.Printf("%v:%v\n", i, v) 
	} 
	fmt.Println("长度即可操控元素个数:",len(s),"\n容量:",cap(s)) 
	s[0] = 5 
	fmt.Println(x)*/
	//map
	pp:=make(map[string]int) 
	//map赋值 
	pp["a"] = 100
	pp["b"] = 90
	pp["c"] = 78
	pp["d"] = 86
	pp["e"] = 97
	//遍历map 
	for people, score := range pp {  
		fmt.Println(people, score) 
	} 
	delete(pp, "d") //map删除 
	//map查询 
	if score, ok := pp["d"]; ok { 
		fmt.Println("d's score is:", score) 
	} else { 
		fmt.Println("d is not in") 
	}
	fmt.Println(pp) //一次性输出map
}