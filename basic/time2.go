package main

import "fmt"
import "time"
// import "github.com/cnwyt/golang-web-api/demo"
// import "github.com/cnwyt/golang-web-api/demo/package1"
// import "github.com/cnwyt/golang-web-api/demo/package2"
import "github.com/cnwyt/golang-web-api/tools"

func main() {
    // 我的hello world程序
    fmt.Println("Hello, World!");

    // demo.SayHello();
    // demo1.SayHello();
    // demo1.SayHelloTest1();
    // demo2.SayHello();
    // demo2.SayHelloTest2();
    
    fmt.Println(tools.GetDatetime());

    // // test2()
    // fmt.Println("-------getDate-------->");

    // fmt.Println(getDate())
    // fmt.Println(getDate(time.Now().Unix()))

    // fmt.Println("-------getDatetime-------->");

    // fmt.Println(getDatetime())
    // fmt.Println(getDatetime((time.Now().Unix())))
    // fmt.Println("-------getDatetime-end------->");
}

func getDate(params ...int64) string {
	    fmt.Println("-------params-end------->");
	    fmt.Println(params);

	res := ""
	format := "2006-01-02"
	if len(params) > 0 {
		timestamp := int64(params[0])
		res = time.Unix(timestamp, 0).Format(format)
	} else {
		res = time.Now().Format(format)
	}

	return res;
}

func getDatetime(params ...int64) string {
	res := ""
	format := "2006-01-02 15:04:05"

	if len(params) > 0 {
		timestamp := int64(params[0])
		res = time.Unix(timestamp, 0).Format(format)
	} else {
		res = time.Now().Format(format)
	}

	return res;
}

func test2(){
	fmt.Println("-------test2-----------")

// 		str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
// fmt.Println(str_time)

	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("------------------")
	fmt.Println(time.Now().Weekday())
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(time.Sunday)
	fmt.Println(time.Saturday)
	// fmt.Println(time.Now().Month().String())

	_, month, day := time.Now().Date()
	fmt.Println(month)
	fmt.Println(day)
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}

	fmt.Println("------------------")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"

	t1, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t1)

	const shortForm = "2006-Jan-02"
	t1, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t1)
}