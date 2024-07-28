package main

import (
	"fmt"
	"os/exec"
)

// 生成外部进程
func main() {
	dateCmd := exec.Command("npm", "-v")
	dateOut, err := dateCmd.Output()
	if err != nil {

	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
}
