package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robfig/cron/v3"
)

var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func main() {

	cronTab := cron.New(cron.WithSeconds())
	// 定义定时器调用的任务函数
	task := func() {
		fmt.Println("hello world", time.Now())
	}
	// 定时任务,cron表达式,每五秒一次
	spec := "*/5 * * * * ?"
	// 添加定时任务
	cronTab.AddFunc(spec, task)
	// 启动定时器
	cronTab.Start()
	// 阻塞主线程停止
	select {}

	// kkk := rand.Int31()
	// fmt.Printf("kkk: %v\n", kkk)
	// // details.Showinfo()
	// kk := detail.GenRound()
	// fmt.Printf("kk: %v\n", kk)
	// var builder strings.Builder
	// builder.WriteString("3244444444442")
	// builder.WriteString("\r\n")
	// builder.WriteString("65756756756")
	// builder.WriteString("\r\n")
	// builder.WriteString("78978978978978")
	// fmt.Println(builder.String())
	// tt := letters[rand.Intn(len(letters))]

	// fmt.Println(string(tt))
	// fmt.Println("tt")

	// b := make([]byte, 6)
	// for i := range b {
	// 	b[i] = byte(letters[rand.Int31()])
	// }
	// str := RandSeq(8)
	// fmt.Printf("str: %v\n", str)

	// var builder strings.Builder
	// builder.WriteString("111111")
	// builder.WriteString("2222")
	// fmt.Println(builder.String())

	// builder.Reset()               // 清空 Builder
	// builder.WriteRune('G')        // 写入一个 rune
	// builder.WriteRune('o')        // 再写入一个 rune
	// fmt.Println(builder.String()) // 输出: Go

	// builder.Reset()
	// builder.WriteByte('t')
	// builder.WriteByte('t')
	// fmt.Println(builder.String())

	// fmt.Println("now()时间:", time.Now(), "显示的是cst时间")
	// //解析并显示自定义时间
	// pt, _ := time.Parse("2006-01-02 15:04:05", "2020-02-06 21:58:59")
	// fmt.Println("parse()时间:", pt, "显示的是utc时间")
	// // //设定时区
	// loc, _ := time.LoadLocation("Asia/Shanghai")
	//根据时区解析并显示自定义时间
	// pts, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-02-06 21:58:59", loc)
	// fmt.Println("parse()时间:", pts.Unix(), "显示的是cst时间")

	// pp := time.Now()
	// kk := pp.UnixNano()
	// fmt.Printf("pp: %v\n", kk)
	// hh := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Printf("hh: %v\n", hh)

	// mm := time.Unix(int64(pp.Unix()), 0)
	// fmt.Printf("mm: %v\n", mm)

	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// pts, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-02-06 21:58:59", loc)
	// fmt.Println("parse()时间:", pts.Year(), "显示的是cst时间")
	// bmp := fmt.Sprintf("%v年%v月%v日", pts.Year(), int(pts.Month()), pts.Day())
	// fmt.Println(bmp)
	// //计算差值
	// fmt.Println("两个时间的差值：", pt.Sub(pts))
	// kk := "haha,nihao,wohao,zhangsan"
	// mm := strings.ToLowerSpecial("h", kk)
	// fmt.Printf("mm: %v\n", mm)

	// strings.ToLowerSpecial("h", "haha,nihao,wohao,zhangsan")
	// var s = make([]int, 3, 5)
	// s[0] = 2

	// s1 := append(s, []int{6, 7}...)
	// fmt.Printf("Slice address: %p\n", s)
	// fmt.Printf("Slice address: %p\n", s1)

	// fmt.Printf("Slice address: %v\n", s)
	// fmt.Printf("Slice address: %v\n", s1)

	// fmt.Printf("s: %T\ns: %v\n", s, s)
	// fmt.Printf("s: %T\ns: %v\n", s1, s1)
	// var p []int
	// fmt.Println(len(p), cap(p))
	// // p = append(p, 3)
	// // fmt.Printf("p: %v\n", p)
	// // fmt.Println(len(p), cap(p))
	// pp := make(map[int]string)
	// pp[1] = "aaa"
	// pp[2] = "bbb"
	// fmt.Println(pp)
	// i := 18
	// j := 9
	// res := show(i, j, sub)
	// fmt.Printf("res: %v\n", res)
}

// type mm func(x, y int) int

// func show(m, n int, sca mm) int {
// 	return sca(m, n)
// }

// func sub(x, y int) int {
// 	return x - y
// }

// func add(x, y int) int {
// 	return x + y
// }
