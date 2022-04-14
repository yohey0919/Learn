// Go语言源文件总是用UTF8编码
// 类似于其它语言里的库（libraries）或者模块（modules）
package main // main包比较特殊。它定义了一个独立可执行的程序，而不是一个库
// 包的导入要加双引号
import "fmt" //fmt包，就含有格式化输出、接收输入的函数

// 在main里的main 函数 也很特殊，它是整个程序执行时的入口
func main() { // 1)fun必须与{ 同一行  , 2)一个go.mod里面只能有一个main()方法
	fmt.Println("Hello, World! -- 1")
	fmt.Println("Hello, World! -- 2") // 格式相关： 1)换行符 <=> 分号（;） 2)代码保存会“自动格式化”为标准格式
}
