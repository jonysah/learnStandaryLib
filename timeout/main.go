package main

import "time"

func main() {
	ch := make(chan string, 1)
	go func(c chan string) {
		time.Sleep(time.Second * 10)
		c <- "result 1 done"
	}(ch)
	println("进入select了")
	select {
	case r := <-ch:
		println(r)
	case <-time.After(time.Second * 5):
		println("请求超时")
	}
	println("完成select了")
}
