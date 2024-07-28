package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//a := os.Getenv("JAVA_HOME")
	//fmt.Println(a)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
