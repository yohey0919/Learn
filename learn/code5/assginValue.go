package main

import (
	"fmt"
	"strings"
)

// 最大公约数（GCD）
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 斐波纳契数列
func fib(n int) int {
	x, y := 0, 1 // 一些简写方式赋值
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {

	fmt.Println("-----变量赋值-----")
	fib5 := fib(5)
	fmt.Println("fib5 = ", fib5)
	medals := []string{"gold", "silver", "bronze"} //用{}赋值数组
	medals[2] = "steel"                            // medals =  [gold silver steel]
	fmt.Print("medals.size = ", len(medals))       //数组长度用 len()
	medals = append(medals, "bronze2")             //medals =  [gold silver steel bronze2]
	fmt.Println(", medals = ", medals)
	fmt.Println("last = ", medals[len(medals)-1])           //bronze2
	fmt.Println("stringjoin = ", strings.Join(medals, "-")) // gold-silver-steel-bronze2
}
