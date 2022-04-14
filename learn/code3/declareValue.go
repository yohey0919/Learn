package main

import "fmt"

// Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明

const boilingF = 212.0 // 常量声明，一级声明语句声明的名字可在整个包对应的每个源文件中访问

// 华氏度转换为摄氏度
func fToC(f float64) float64 { //函数声明
	return (f - 32) * 5 / 9
}

// 返回函数中局部变量的地址也是安全的
func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func main() {
	fmt.Println("-----常量变量声明-----")
	const freezingF = 32.0                                         // 局部常量仅在main()中起作用                                      // 常量声明，一级声明语句声明的名字可在整个包对应的每个源文件中访问
	fmt.Printf("局部常量： %g°F = %g°C \n", freezingF, fToC(freezingF)) // %g 浮点数
	fmt.Printf("全局常量： %g°F = %g°C \n", boilingF, fToC(boilingF))
	// var 变量声明，用于需要显式指定变量类型
	fmt.Println("-----var变量声明-----")
	var s1 string        // 变量没有显式初始化，则被隐式地赋予其类型的"零值"，此时""
	var s2 string = "s2" // 显示初始化
	var s3 []string
	// var t, f, s = true, 2.3, "four" // 直接声明, 这种方法一般不用
	fmt.Printf("s1=%s, s2=%s", s1, s2) //s1=, s2=s2
	fmt.Println(", s3=", s3)           //s3= []

	// 简短变量声明 ，用于声明和初始化局部变量，不需要显式指定变量类型
	fmt.Println("-----简短变量声明-----")
	k := 2.1
	i, j := 0, 1 // 可以使用 i, j = j, i 交换 i 和 j 的值
	fmt.Printf("i=%d, j=%d, k=%g \n", i, j, k)

	// 指针声明, 一个变量对应一个保存了变量对应类型值的内存空间
	fmt.Println("-----指针声明-----")
	x := 11
	p := &x                              // p类型为 *int, 指向x
	fmt.Println("p =", p)                // p的值为x的地址
	fmt.Println("*p =", *p)              // *p=11
	*p = 22                              // *p就是变量v的别名。指针特别有价值的地方在于我们可以不用名字而访问一个变量
	fmt.Println("after *p =22 , x =", x) // x=22
	fmt.Println("f()=f()", f() == f())   //false 每次的f()返回的局部变量地址，都不同
	v := 1
	fmt.Println("first time v =", incr(&v))  // v = 2
	fmt.Println("second time v =", incr(&v)) // v = 3

	//表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T
	fmt.Println("-----new函数声明-----") //new函数使用通常相对比较少
	p1 := new(int)                   // p1, *int 类型, 指向匿名的 int 变量
	fmt.Print("p1 = ", p1)           //值为地址
	fmt.Print(" ,*p1 = ", *p1)       //零值
}
