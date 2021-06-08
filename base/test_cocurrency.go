package main

import (
	"fmt"
	"time"
)

/*goroutine*/
func say(s string)  {
	for i:=0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

/*channel*/
func sumData(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	// 看到输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sumData(s[:len(s)/2], c)
	go sumData(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)
}
