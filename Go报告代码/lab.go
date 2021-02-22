package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
    "sort"
)

type Stu struct {
    stuNum  string
    stuName string
    goScore int
}

func (s *Stu) SetstuNum(stuN string) {
    s.stuNum = stuN
}
func (s *Stu) GetstuNum() string {
    return s.stuNum
}
func (s *Stu) SetstuName(stuN string) {
    s.stuName = stuN
}
func (s *Stu) GetstuName() string {
    return s.stuName
}
func (s *Stu) SetstuScore(stuS int) {
    s.goScore = stuS
}
func (s *Stu) GetstuScore() int {
    return s.goScore
}
func menu() {
    fmt.Println("1------录入成绩")
    fmt.Println("2------查找成绩")
    fmt.Println("3------修改成绩")
    fmt.Println("4------删除成绩")
    fmt.Println("5------成绩排序")
    fmt.Println("6------成绩统计")
    fmt.Println("7------退出")
}
func Clsscr() {
    cmd := exec.Command("cmd", "/c", "cls") // /c执行字符串指定的命令然后终止
    cmd.Stdout = os.Stdout
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}
func Goon() {
    fmt.Println("按回车键继续")
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
}
func Getdata(stuslice map[string]Stu) {
    stu1 := Stu{}
    n := 0
    var s1, s2 string
    var s3 int
    fmt.Println("输入学生个数")
    fmt.Scanln(&n)
    for i := 0; i < n; i++ {
        fmt.Println("input name of student", i+1)
        fmt.Scanln(&s1)
        stu1.SetstuName(s1)
        fmt.Println("input number of student", i+1)
        fmt.Scanln(&s2)
        stu1.SetstuNum(s2)
        fmt.Println("input goscore of student", i+1)
        fmt.Scanln(&s3)
        stu1.SetstuScore(s3)
        stuslice[stu1.stuName] = stu1
    }
}
func Find(stuslice map[string]Stu) {
    stuName1 := ""
    fmt.Println("请输入想要查找的学生姓名：")
    fmt.Scanln(&stuName1)
    studata, key := stuslice[stuName1]
    if key {
        fmt.Println("查询成功！学生信息为：", studata)
    } else {
        fmt.Println("未找到")
    }
}
func Change(stuslice map[string]Stu) {
    stuName1 := ""
    stuName2 := ""
    stuNum1 := ""
    stuscore1 := 0
    fmt.Println("请输入想要修改的学生的姓名：")
    fmt.Scanln(&stuName1)
    studata, key := stuslice[stuName1]
    if key {
        fmt.Println("input the change name:")
        fmt.Scanln(&stuName2)
        studata.stuName = stuName2
        fmt.Println("input the change stunum:")
        fmt.Scanln(&stuNum1)
        studata.stuNum = stuNum1
        fmt.Println("input the change goscore:")
        fmt.Scanln(&stuscore1)
        studata.goScore = stuscore1
        stuslice[stuName1] = studata

    } else {
        fmt.Println("输入错误")
    }
}
func Delete(stuslice map[string]Stu) {
    stuName1 := ""
    fmt.Println("请输入想要删除的学生的姓名：")
    fmt.Scanln(&stuName1)
    delete(stuslice, stuName1)
}
func Sort(stuslice map[string]Stu) {
    var stusli = make([]Stu, 0, 5)
    for _, v := range stuslice {
        stusli = append(stusli, v)
    }
    sort.Slice(stusli, func(i, j int) bool { return stusli[i].goScore > stusli[j].goScore })
    fmt.Println(stusli)
}
func Statistics(stuslice map[string]Stu) {
    sumA := 0
    sumB := 0
    sumC := 0
    sumD := 0
    sumE := 0
    for i, _ := range stuslice {
        if stuslice[i].goScore >= 90 && stuslice[i].goScore < 100 {
            sumA++
        }
        if stuslice[i].goScore >= 80 && stuslice[i].goScore < 90 {
            sumB++
        }
        if stuslice[i].goScore >= 70 && stuslice[i].goScore < 80 {
            sumC++
        }
        if stuslice[i].goScore >= 60 && stuslice[i].goScore < 70 {
            sumD++
        }
        if stuslice[i].goScore >= 0 && stuslice[i].goScore < 60 {
            sumE++
        }
    }
    fmt.Println("A:", sumA, "B:", sumB, "C:", sumC, "D", sumD, "E:", sumE)

}
func main() {
    studentmap := make(map[string]Stu)
    for {
        var i int
        menu()
        fmt.Scanln(&i)
        switch i {
        case 1:
            Getdata(studentmap)
            Goon()
            Clsscr()
        case 2:
            Find(studentmap)
            Goon()
            Clsscr()
        case 3:
            Change(studentmap)
            Goon()
            Clsscr()
        case 4:
            Delete(studentmap)
            Goon()
            Clsscr()
        case 5:
            Sort(studentmap)
            Goon()
            Clsscr()
        case 6:
            Statistics(studentmap)
            Goon()
            Clsscr()
        case 7:
            os.Exit(0)

        }
    }
}
