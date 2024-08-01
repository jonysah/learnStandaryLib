package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Hello, World!"
	b := unsafeStringToBytes(s)
	fmt.Println(b)
}

// 字符串转字节数组0内存拷贝
func unsafeStringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
