//Go语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型
package main

import (
	"fmt"
	"math"
)

// 你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的
// 列如： func Sin(x float64) float64  （math包中）

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func first(x int, _ int) int { //_符号,可以强调某个参数未被使用
	return x
}

func main() {
	fmt.Println(hypot(3, 4))
	fmt.Println(first(1, 2))
}
