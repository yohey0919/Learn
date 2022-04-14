package main

// 多个包的导入用 import(), 并且要带"",用空行隔开
import (
	"fmt"
	"os" // 提供了一些与操作系统交互的函数和变量
	"strings"
)

func main() {
	var s1 string        // 变量没有显式初始化，则被隐式地赋予其类型的"零值"，此时""
	var sep1 string = "" // 显示初始化

	fmt.Println(os.Args) //os.Args变量是一个string的切片,比如s[m:n]这个切片

	for i := 1; i < len(os.Args); i++ { // :=是短变量声明，只能用在函数内部，而不能用于包变量
		s1 += sep1 + os.Args[i] //字符串拼接, 推荐使用strings包的Join函数
		sep1 = " "
	}
	// for (Go语言只有for循环这一种循环语句)
	// 1) for initialization; condition; post {}  ，其中3个部分都是可以省略的
	// 2) for condition {}  等价于 其他语言的 while
	// 3) for _, arg := range os.Args[1:] {}  等价于 for( : )
	// 4) for{}  无限循环 ，可以用break，return退出

	s2, sep2 := "", ""
	// range的语法要求，要处理元素，必须处理索引（但是我这里不用index）
	for _, arg := range os.Args[1:] { // 这种情况的解决方法是用空标识符，即_（下划线）
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(strings.Join(os.Args[0:], " yahaha "))
}
