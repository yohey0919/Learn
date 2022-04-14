package main

import (
	"flag" //使用命令行参数来设置对应变量的值
	"fmt"
	"strings"
)

//第一个是命令行标志参数的名字“n”，然后是该标志参数的默认值（这里是false），最后是该标志参数对应的描述信息
var n = flag.Bool("n", false, "忽略行尾的换行符") //输入-n为true,不输入false
var sep = flag.String("s", " ", "分隔符")

func main() {
	flag.Parse()                               //调用flag.Parse函数，用于更新每个标志参数对应变量的值
	fmt.Print(strings.Join(flag.Args(), *sep)) // 通过调用flag.Args()来访问,“-命令” 以外的其他参数
	if !*n {
		fmt.Println(" end") //末尾添加一个 end
	}
}
