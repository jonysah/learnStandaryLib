package main

import "os"

func main() {
	defer println("函数执行完成")

	os.Exit(0)
}
