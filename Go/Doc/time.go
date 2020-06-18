package main

import (
	"fmt"
	"time"
)

const (
	BASE = "2006-01-02 15:04:05"
)

/*
 * 时间单位
 *		1s  = 1000ms
 *		1ms = 1000us
 *		1us = 1000ns
 * 时区
 * 		GMT 林威治标准时间
 *		UTC 世界协调时间(和GMT差不多)
 *		CST
 *			Central Standard Time (USA) UT-6:00        美国
 * 			Central Standard Time (Australia) UT+9:30  澳大利亚
 *			China Standard Time UT+8:00                中国
 *			Cuba Standard Time UT-4:00                 古巴
 */
func main() {

	/*
	 * Time 对象初始化
	 */
	//字符串 -> Time 对象(格式必须完全一样)
	fmt.Println(time.Parse(BASE, "2020-01-08 18:01:48"))

	//字符串 -> Time 对象(格式必须完全一样)(加入时区配置)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.ParseInLocation(BASE, "2020-01-08 18:01:48", loc))

	//时间戳 -> Time 对象(秒+纳秒ns)
	fmt.Println(time.Unix(1578477708, 0))

	//当前时间 -> Time 对象
	fmt.Println(time.Now())

	//构造时间对象 -> 2020-01-02 03:04:05
	fmt.Println(time.Date(2020, 1, 8, 18, 01, 48, 0, time.Local))

	/*
	 * Time 方法
	 */
	t := time.Now()

	// 获取时间详细信息
	// 年 Year
	// 月 Month
	// 日 Day
	// 时 Hour
	// 分 Minute
	// 秒 Second
	// 当前时间是星期几 t.Weekday()
	// 是一年中的第几天 t.YearDay()
	// 时间戳
	fmt.Println(fmt.Sprintf("当前时间 %d-%d-%d %d:%d:%d \n 当前时间是星期%d \n 是一年中的第%d天 \n 当前时间戳为%d",
		t.Year(),
		int(t.Month()),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		int(t.Weekday()),
		t.YearDay(),
		t.Unix(),
	))

	//获取时间和时区详细信息
	fmt.Println(t.Local())

	//获取时区
	fmt.Println(t.Zone())

	//返回时间 [年 月 日]
	fmt.Println(t.Date())

	//返回时间 [时 分 秒]
	fmt.Println(t.Clock())

	//格式化时间输出 [不常用]
	fmt.Println(t.String())

	//Format格式化输出
	fmt.Println(t.Format(BASE))

	/*
	 * 时间大小判断(考虑时区影响的比较)
	 * 不推荐使用，推荐使用unix(时间戳比较)
	 */
	t1, _ := time.Parse(BASE, "2020-01-08 18:00:00")
	t2, _ := time.Parse(BASE, "2020-01-08 17:00:00")

	//t1 == t2 ?
	fmt.Println(t1.Equal(t2))

	// t1 < t2 ?
	fmt.Println(t1.Before(t2))

	// t1 > t2 ?
	fmt.Println(t1.After(t2))
}
