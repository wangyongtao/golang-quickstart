package main 

import (
	"fmt"
	"time"
	// "strconv"
)

func main() {

	fmt.Println(time.Now()) // 2019-04-30 14:41:59.661602 +0800 CST m=+0.000225294
	fmt.Println(time.Now().String()) // 2019-04-30 14:41:59.661826 +0800 CST m=+0.000448434

	fmt.Println(time.Now().Unix()) // 1556615702
	fmt.Println(time.Now().UnixNano() / 1e6) // 1556615702009
	fmt.Println(time.Now().UnixNano()) // 1556615702009257000

	fmt.Println(time.Now().Format("2006-01-02")) // 2019-04-30
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2019-04-30 14:43:26
	fmt.Println(time.Now().Format("2006-01-02 00:00")) // 2019-04-30 00:00
	fmt.Println(time.Now().Format("2006/01/02 15:04")) // 2019-04-30 14:43
	fmt.Println(time.Now().Format("2006/Jan/02 15:04")) // 2019/Apr/30 17:28

	// 使用 Date 构造时间
	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// 当前时间: time.Now().Date() 返回三个参数: 年月日
	year1, month1, day1 := time.Now().Date()
	fmt.Printf("year: %v, type: %T \n", year1, year1) // year: 2019, type: int 
	fmt.Printf("month: %v, type: %T \n", month1, month1) // month: April, type: time.Month 
	fmt.Printf("day: %v, type: %T \n", day1, day1) // day: 30, type: int 

	// 直接指定时间: 年，月，日，时，分，秒，时区
	t2 := time.Date(2019, time.November, 28, 11, 35, 46, 0, time.UTC)
	fmt.Printf("=> 输出UTC时间: %s\n", t2.String()) // => 输出UTC时间: 2019-11-28 11:35:46 +0000 UTC
	fmt.Printf("=> 输出本地时间: %s\n", t2.Local())   // => 输出本地时间: 2019-11-28 19:35:46 +0800 CST
	fmt.Printf("=> 输出时间戳: %d\n", t2.Unix())   // => 输出本地时间: 1574940946
	fmt.Printf("=> 输出日期格式: %s\n", t2.Format("2006-01-02 15:04:05"))   // => 输出日期格式: 2019-11-28 11:35:46
	fmt.Printf("=> 输出日期格式: %s\n", t2.Format("2006/01/02 15:04:05"))   // => 输出日期格式: 2019/11/28 11:35:46
	fmt.Printf("=> 输出日期格式: %s\n", t2.Format("06/01/02 15:04:05"))   // => 输出日期格式: 19/11/28 11:35:46
 

	t := time.Now()	// 2019-04-30 16:38:49.786745 +0800 CST m=+0.000237155
	y := t.Year()   // 年 (类型:int)
	m := t.Month()  // 月 (类型:string) (t.Month()返回类型:Month)
	d := t.Day()    // 日 (类型:int)
	h := t.Hour()   // 小时 (类型:int)
	i := t.Minute() // 分钟 (类型:int)
	s := t.Second() // 秒 (类型:int)
	fmt.Println(y, m, d, h, i, s) // 2019 April 30 16 38 49
	// datetime = 
	fmt.Printf("=> %v %T\n", t, t)
	fmt.Printf("=> %v %T\n", y, y)
	fmt.Printf("=> %v %T\n", m, m)
	fmt.Printf("=> %v %T\n", d, d)
	fmt.Printf("=> %v %T\n", h, h)
	fmt.Printf("=> %v %T\n", i, i)
	fmt.Printf("=> %v %T\n", s, s)

	// fmt.Println(strconv.Atoi(t.Month()))



	fmt.Println("--------Weekday----------")
	fmt.Println(time.Now().Weekday())
	now := time.Now()
	// fmt.Println(today)
	// fmt.Println(time.Sunday)
	// fmt.Println(time.Saturday)
	// 获取当前时间的月份: 
	month2 := now.Month().String()
	weekday2 := now.Weekday().String()
	hour2 := now.Hour()
	minute2 := now.Minute()
	fmt.Printf("=> 当前时间是几月: %s %T\n", month2, month2) // => 当前时间是几月: April string
	fmt.Printf("=> 当前时间是周几: %s %T\n", weekday2, weekday2) // => 当前时间是周几: Tuesday string
	fmt.Printf("=> 当前时间是几点: %d %T\n", hour2, hour2) // => 当前时间是几点: 15 int
	fmt.Printf("=> 当前时间是几分: %d %T\n", minute2, minute2) // => 当前时间是几分: 47 int

	// 字符串转化成时间戳
	x := "2018-12-27 18:44:55"
    p, _ := time.Parse("2006-01-02 15:04:05", x)
    fmt.Println( p.Unix() ) // 输出: 1545936295
}