package main

import (
	"errors"
)

func check(err error) {
	if err != nil {
		if err != nil {
			errors.New("创建文件错误")
		}
	}
}

func main() {
	//err := os.Mkdir("directories/tem", 0775)
	//check(err)
	//defer os.RemoveAll("directories/tem")

	//createEmptyFile := func(name string) {
	//	d := []byte("")
	//	check(os.WriteFile(name, d, 0755))
	//}
	//createEmptyFile("directories/file1")
}
