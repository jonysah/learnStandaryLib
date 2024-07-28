package main

import (
	"time"
)

func worker(id int, jobs chan int, result chan int) {
	for job := range jobs {
		println("worcker", id, "start working", job)
		time.Sleep(time.Second)
		println("worcker", id, "done work", job)
		result <- job
	}

}

func main() {
	jobs := make(chan int, 5)
	result := make(chan int)

	//开启三个协程
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, result)
	}

	//调用协程
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	println("所有工作发送完毕")
	close(jobs)

	for r := range result {
		println("结果", r)
	}

}
