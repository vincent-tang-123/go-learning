package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 时间对象
	fmt.Println(now) // 2022-05-28 19:40:21.784995 +0800 CST m=+0.000144195

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println(year, month, day, hour, minute, second) // 2022 May 28 19 42 33

	// 时间戳 从 1970.1.1 到现在的一个秒数
	timeStamp1 := now.Unix() // 秒
	timeStamp2 := now.UnixNano() // 纳秒
	fmt.Println(timeStamp1, timeStamp2) // 1653738375 1653738375850393000

	// 将一个时间戳转换成具体的时间格式
	t := time.Unix(1653738375, 0)
	fmt.Println(t) // 2022-05-28 19:46:15 +0800 CST

	// 时间间隔
	//time.Sleep(5)
	//time.Duration()
	//n := 5 // type int
	//time.Sleep(time.Duration(n) * time.Second)
	// 延迟5s
	//time.Sleep(time.Duration(n) * time.Second)
	// 延迟5 分钟
	//time.Sleep(time.Duration(n) * time.Minute)
	fmt.Println("over")


	// 1 小时后的时间
	// now + 1 hour
	fmt.Println(now) // 2022-05-28 23:08:02.118615 +0800 CST m=+0.000100241

	t2 := now.Add(time.Hour)
	fmt.Println(t2) // 2022-05-29 00:08:02.118615 +0800 CST m=+3600.000100241

	// sub 减去一个 Duration 类型对象的 差值
	fmt.Println(t2.Sub(now)) // 1h0m0s

	// 定时器
	// time.Second 一秒触发一次
	//for tmp := range time.Tick(time.Second * 2){ //  * 2 即 两秒一次
	//	fmt.Println(tmp)
	//}

	// 时间格式化:  Y   m   d   H    M   S   年月日时分秒
	//          2006 01  02  15   04  05

	ret1 := now.Format("2006-01-02")
	fmt.Println(ret1) // 2022-05-30
	ret2 := now.Format("2006-01-02 15:04:05")
	fmt.Println(ret2) // 2022-05-30 00:13:02
	ret3 := now.Format("2006-01-02 03:04:05 PM")
	fmt.Println(ret3) // 2022-05-30 12:22:18 AM
	ret4 := now.Format("2006-01-02 15:03:04.000") // 拿到毫秒
	fmt.Println(ret4)

	// 解析字符串类型的时间
	timeStr := "2019/08/07 15:00:00"
	// 1 拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(loc) // Asia/Shanghai

	// 2 根据时区解析一个字符串格式的时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil{
		fmt.Printf("parse timeStr failed , err %v\n", err)
		return
	}
	fmt.Println(timeObj) // 2019-08-07 15:00:00 +0000 UTC   utc 时间

	// 传入时区
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil{
		fmt.Printf("parse timeStr failed , err %v\n", err)
		return
	}
	fmt.Println(timeObj2) // 2019-08-07 15:00:00 +0800 CST cst 东8区时间









}
