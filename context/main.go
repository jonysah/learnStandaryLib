package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second): // 10s后超时
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done(): // 用户断开链接
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
func valueCtx() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	ctx = context.WithValue(ctx, "key2", "value2")
	ctx = context.WithValue(ctx, "key3", "value3")
	fmt.Println(ctx.Value("key"))
}
func cancelCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("cancel ctx")
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {

	//http.HandleFunc("/hello", hello)
	//http.ListenAndServe(":8090", nil)
	//valueCtx()
	cancelCtx()
}
