package main

import (
	"fmt"
	"log"
	"net/http" //web服务器变得异常地简单
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // 有发送到/路径下的请求和handler函数关联起来
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) //服务监听8000端口
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++ // 当前web服务访问次数加+1
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) //返回当前用户正在访问的URL
}

func counter(w http.ResponseWriter, r *http.Request) { // 查看当前web服务访问次数
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
