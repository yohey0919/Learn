package main

import (
	"fmt"
)

// 指针练习
func main() {

	// p存放的是地址,而*p是让程序去那个地址取出的数据
	var value1 int16 = 11
	p := &value1                        //此时的p ,是一个 int16类型的 指针变量(存储指针的变量)
	fmt.Println("p =", p)               //0xc000014098, p存放的是value的地址
	fmt.Println("*p =", *p)             //11, *p表示此指针p指向的内存地址中存放的内容
	fmt.Println("*&value1 =", *&value1) //11, 先获取地址，再指针指向地址中存放的内容

	var value2 int16 = 22
	var s *int16                     //声明一个 int16类型的 指针变量s
	fmt.Println("s = ", s)           //<nil>, 此时的s,存储的指针指向nil
	s = &value2                      //此时的s存放的是value2的地址
	fmt.Println("*s = ", *s)         //22
	*s = 33                          //*s的改变 影响 value2
	fmt.Println("value2 = ", value2) //33
	value2 = 44                      //value2的改变 影响 *s
	fmt.Println("*s = ", *s)         //44

	var value3 int16 = 55
	var q *int16 = &value3   // <=>  var q *int16 ; q =&value3
	fmt.Println("q = ", q)   //0xc0000140ba
	fmt.Println("*q = ", *q) //55

	var value4 int16 = 66
	var l *int16 = &value4
	mutil2(l)
	fmt.Println("*l = ", *l) //132
	mutil2(&value4)
	fmt.Println("value4 = ", value4) //264
}

func mutil2(p *int16) {
	*p = *p * 2
}
