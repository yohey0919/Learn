package main

import "fmt"

var g string = "global"

func main() {
	fmt.Println("for的作用域一")
	x := "helloworld!"
	// fmt.Printf("%c", 0) //该操作会导致后续的输出无法进行!!
	for i := 0; i < len(x); i++ {
		s := x[i]
		if s != '!' {
			s := s + 'A' - 'a'  // 小写转换为大写
			fmt.Printf("%c", s) // Printf 支持 %c,%g
		}
	}
	fmt.Println()
	fmt.Println("for的作用域二")
	y := "test!"
	for index, z := range y { //range 会返回两个值 index,value
		fmt.Printf("index = %d, z = %c\n", index, z)
	}
	fmt.Println("if的作用域")
	if x := 2; x == 1 { // if条件不带括号
		fmt.Println("x = ", x)
	} else if y := 3; y == x { // 第一个判断内的语句（x:=0）可以作用在第二个判断语句内(以此类推到最后的else)
		fmt.Println("y = ", y)
	} else {
		fmt.Print("x = ", x, ", ", "y = ", y) // x = 2, y = 3
	}
	//g:="ss"		// 虽然g在外部已经声明过，但是:=语句还是将g重新声明为新的局部变量

}
