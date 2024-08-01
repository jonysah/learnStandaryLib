package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	name string
	age  int
}

func main() {
	GetPrivateFiled()
}

// 通过unsafe指针获取私有变量
func GetFiledByUnsafe() {
	p := Person{name: "jony", age: 18}
	ref := reflect.ValueOf(&p).Elem()
	nameFiled := ref.FieldByName("name")
	ageFiled := ref.FieldByName("age")
	name := reflect.NewAt(nameFiled.Type(), unsafe.Pointer(nameFiled.UnsafeAddr())).Elem().Interface().(string)
	age := reflect.NewAt(ageFiled.Type(), unsafe.Pointer(ageFiled.UnsafeAddr())).Elem().Interface().(int)
	fmt.Println(name)
	fmt.Println(age)
}

type myStruct struct {
	privateField int
	PublicField  int
}

// 会panic，因为反射不能直接读取到私有变量
func GetPrivateFiled() {
	ms := myStruct{10, 20}
	val := reflect.ValueOf(ms)

	// 遍历结构体的所有字段
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("Field %d: %v\n", i, field.Interface())
	}
}
