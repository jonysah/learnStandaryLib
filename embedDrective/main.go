package main

import (
	"embed"
	"errors"
	"fmt"
)

func main() {
	var floder embed.FS

	c, err := floder.ReadFile("file1")
	if err != nil {
		errors.New("floder读取错误")
	}
	fmt.Println(string(c))
}
