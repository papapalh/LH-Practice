package main

import (
	"fmt"
	"time"
)

const (
	BASE = "2019-09-09"
)

/*
 * 1s  = 1000ms
 * 1ms = 1000us
 * 1us = 1000ns
 */
func main() {
	fmt.Println("Time 方法:")
	fmt.Println("    Weekday(a)  返回星期a的英文名 ", time.Weekday(0))
	fmt.Println("    Month(a)    返回月份a的英文名 ", time.Month(1))

	fmt.Println("    Date()      构建一个时间点:", time.Date(2020, 1,2,1,1,1,1,time.Local))

	fmt.Println("    Now()       返回当前本地时间:", time.Now())

	fmt.Print("    Parse()     解析一个格式化字符时间串:")
	e, _ := time.Parse(BASE, "2019-01-01")
	fmt.Println(e)

	fmt.Print("    Unix()      根据时间戳构建一个Time对象(秒+纳秒ns):")
	f := time.Unix(1578477708,0)
	fmt.Println(f)

	fmt.Println("Time 对象:")
	fmt.Println("    Local   返回地点和时区信息 ", f.Local())
	fmt.Print("    Zone    返回时区信息 ")
		fmt.Println(f.Zone())
	fmt.Println("    Equal   判断两个时间是否相同(考虑时区影响) ", f.Equal(time.Unix(1578477708,0)))
	fmt.Println("    Before  判断时间在传入时间之前，如果f代表的时间点在X之前,返回真,否则返回假(考虑时区影响) ", f.Before(time.Unix(1578477708,0)))
	fmt.Println("    After   判断时间在传入时间之后，如果f代表的时间点在X之后,返回真,否则返回假(考虑时区影响) ", f.After(time.Unix(1578477708,0)))
	fmt.Print("    Date    返回时间点f对应的年、月、日。 ", )
		fmt.Println(f.Date())
	fmt.Print("    Clock   返回时间点f对应的时、分、秒。", )
		fmt.Println(f.Clock())
	fmt.Println("    Year    返回时间点f对应的年", f.Year())
	fmt.Println("    Month   返回时间点f对应的月", f.Month())
	fmt.Println("    Day     返回时间点f对应的日", f.Day())
	fmt.Println("    YearDay 返回时间点f对应一年的第几天", f.YearDay())
	fmt.Println("    Weekday 返回时间点f对应一年的周几", f.Weekday())

	fmt.Println("    Hour    返回时间点f对应一天的第几小时", f.Hour())
	fmt.Println("    Minute  返回时间点f对应一年的第几分钟", f.Minute())
	fmt.Println("    Second  返回时间点f对应一年的第几秒", f.Second())
	fmt.Println("    Format  格式化时间", f.Format(BASE))
	fmt.Println("    String  字符串的格式化时间 ", f.String())
}

