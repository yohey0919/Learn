package main

import (
	"fmt"
	"io/ioutil"
	"net/http" //这些包可以更简单地用网络收发信息，还可以建立更底层的网络连接，编写服务器程序
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// http.Get函数是创建HTTP请求的函数
		resp, err := http.Get(url) //如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1) //os.Exit函数来终止进程，并且返回一个status错误码
		} // HTTP请求出错，提示
		b, err := ioutil.ReadAll(resp.Body) // io流读取body数据
		resp.Body.Close()                   //关闭resp的Body流，防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1) //os.Exit函数来终止进程，并且返回一个status错误码
		} // 读io流数据出错，提示
		fmt.Printf("%s", b)
	}
}
