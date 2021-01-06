package main

import (
	"container/ring"
	"fmt"
	//"example.com/slidingAlgorithm"
	"time"
	"github.com/shirou/gopsutil/cpu"
)


var (
	limitTime time.Duration = 100 // 100 毫秒
	limitTimeBucket time.Duration = 500 //
	limitBucket int = 5 // 窗口宽度是 500毫秒 5*100
	head *ring.Ring     // 环形队列（链表）
)

func main(){
	ci := make(chan int)
	fmt.Print("hello")
	getCpuInfo()
	// 每100毫秒采样一次cpu

	// 初始化滑动窗口
	head = ring.New(limitBucket)
	for i := 0; i < limitBucket; i++ {
		head.Value = 0.00
		head = head.Next()
	}

	// 启动执行采样数据
	// 启动执行器
	go func() {
		timer := time.NewTicker(time.Millisecond * limitTime )
		for range timer.C {
			preval :=head.Prev()
			preval.Value =  getCpuInfo()
			fmt.Println("获取cpu 利用率 preval.Value ", preval.Value)
		}
	}()



	// 启动执行器
	go func() {
		timer := time.NewTicker(time.Millisecond * limitTimeBucket )
		for range timer.C {
			// 定时每隔500毫秒计算一次： 滑动窗口数据相加取平均值
			var prePrecent float64 =0.00
			var newPrecent float64 = 0.00

			arr := [6]float64{}

			// 打印出 上次五个平均值，最新的 平均值。
			for i := 0; i <= limitBucket; i++ {
				if i <5 {
					prePrecent = prePrecent +  head.Value.(float64)
				}
				if i>0 {
					newPrecent = newPrecent +  head.Value.(float64)
				}
				arr[i] = head.Value.(float64)
				head = head.Next()
			}
			fmt.Println("move prePrecent,newPrecent,arr", prePrecent, newPrecent,arr)
			fmt.Println("move prePrecent/5,newPrecent/5", prePrecent/5, newPrecent/5,arr)
			head.Value = 0.00
			head = head.Next()
		}
	}()

	<-ci
}

func getCpuInfo() float64{
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率

		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	return percent[0]
}
