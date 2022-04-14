package main

// 问题：is not in GOROOT (C:\Program Files\Go\src\learn\code7\tempconv);
// 解决：go env -w  GO111MODULE=off
import (
	"fmt"
	"learn/code6/popcount"
	"learn/code6/tempconv" //导入的包路径
)

func main() {
	fmt.Print(tempconv.BoilingC.String() + "=" + tempconv.Celsius(100).String() + "=")
	fmt.Println(tempconv.CToF(100).String())
	var un uint64
	un = 1000
	fmt.Println(popcount.PopCount(un)) //6
}
