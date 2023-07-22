package main

import (
	"fmt"
	"time"
)
func main(){
	now:=time.Now()
	fmt.Println(now)
	fmt.Printf("year:%v\n",now.Year())
	fmt.Printf("month:%v\n",now.Month())
	fmt.Printf("month:%v\n",int(now.Month()))
	fmt.Printf("day:%v\n",now.Day())
	fmt.Printf("hour:%v\n",now.Hour())
	fmt.Printf("minute:%v\n",now.Minute())
	fmt.Printf("second:%v\n",now.Second())
	fmt.Printf("date&time:%d-%d-%d %v:%v:%v\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	datestr:=fmt.Sprintf("date&time:%d-%d-%d %v:%v:%v\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(datestr)
	//特定格式,各参数数字必须是固定的，必须这样写。任意组合均可，但对应数字不可更改
	datestr2:=now.Format("2006/01/02 15:04:05")
	fmt.Println(datestr2)
}
