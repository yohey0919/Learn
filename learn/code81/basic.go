//Go语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("-----------基础类型")
	fmt.Println("1) 整形")
	//int8、int16、int32和int64有符号整数类型(四种大小不同)
	var i int8 = -1 // 范围-128 ~ 127
	fmt.Println("i = ", i)
	//uint8、uint16、uint32和uint64四种无符号整数类型(四种大小不同)
	var u uint8 = 1 // 范围0 ~ 255
	fmt.Println("u = ", u)
	var run rune = 1            // type rune = int32
	fmt.Printf("run = %d", run) // rune（翻译：符文）
	var apples int32 = 1
	var oranges int16 = 2
	//var compote = apples + oranges	// 类型不符合无法相加
	var compote = int(apples) + int(oranges) // 使用int()强转
	fmt.Printf(", compote = %#x\n", compote) // #副词,输出生成0、0x或0X前缀

	fmt.Println("2) 浮点数")
	// 两种精度的浮点数，float32和float64,通常应该优先使用float64类型
	var e float32 = 2.71828 //浮点数的字面值可以直接写小数部分, 此时的e为常量 无法强转
	fmt.Printf("e = %g", e)
	var Planck float32 = 6.62606957e-34 // 普朗克常数
	fmt.Printf("Planck = %g\n", Planck)

	fmt.Println("3) 复数")
	// 两种精度的复数类型：complex64和complex128
	var x complex64 = complex(1, 2)  // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * (complex64(y)))

	fmt.Println("4) 布尔类型")
	// 一个布尔类型的值只有两种：true和false
	t := true
	f := false
	fmt.Println(!f == t)

	fmt.Println("5) 字符串类型")
	s := "hello, world"
	// s[0] = 'L'	// 字符串是不可修改的,尝试修改字符串内部数据的操作是被禁止的
	fmt.Println(len(s))              // "12"
	fmt.Println(s[0], s[7])          // "104 119" ('h' and 'w')
	fmt.Println(s[0:5])              // "hello"
	fmt.Println(s == "hello, world") //字符串可以用==和<进行比较；比较通过逐个字节比较完成的
	str := "abc"
	b := []byte(str) // type byte = uint8
	fmt.Println("b = ", b)
	var buf bytes.Buffer //bytes包还提供了Buffer类型用于字节slice的缓存。
	buf.WriteByte('[')
	bt := 'a'
	for i := 0; i < 5; i++ {
		bt++
		buf.WriteString(string(bt))
	}
	buf.WriteByte(']')
	fmt.Println("buf = ", buf.String())
	fmt.Println("字符串<-数字的转换")
	x1 := 123
	is1 := fmt.Sprintf("%d", x1) //用这个还可以提供一些额外信息，比如在数字后面加单位，fmt.Sprintf("%d F", x1)
	is2 := strconv.Itoa(x1)
	// %q   带双引号的字符串"abc"或带单引号的字符'c'
	fmt.Printf("%q %q\n", is1, is2) // "123" "123"
	fmt.Println("字符串->数字的转换")
	si1, _ := strconv.Atoi("123")             // x is an int
	si2, _ := strconv.ParseInt("123", 10, 64) // base 10, 第三个参数是用于指定返回的整型数的大小（int64）
	fmt.Println(si1, si2)                     // 123 123

	fmt.Println("5) 常量") //常量的值是在编译期就确定的,常量的值不可修改

	type Weekday int // 声明类型

	const (
		Sunday Weekday = 2 * iota //在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0
		Monday                    //然后在每一个有常量声明的行加一。
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("Sunday = %d, Monday = %d, Tuesday = %d", Sunday, Monday, Tuesday)
	const hide = 'a' // const hide untype rune = 97 无类型常量

}
