package timedemo

import (
	"fmt"
	"time"
)

func TimeTest() {
	time1 := time.Now()
	fmt.Println(time1, time1.Local().Year(), time1.Local().Month(), time1.Local().Day(), time1.Local().Weekday())
	formatTime := fmt.Sprintf("%d-%d-%d %d:%d:%d", time1.Year(), time1.Month(), time1.Day(), time1.Hour(), time1.Minute(), time1.Second())
	fmt.Println(formatTime)

	str1 := time1.Format("2006/01/02 15:04:05")
	str2 := time.Now().Local().Format("2006/01/02 15:04:05")
	// 固定字符串2006/01/02 15:04:05,格式化返回当前时间Now()
	fmt.Println(str1, str2)

}
