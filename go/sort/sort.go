package main

import "fmt"

// 从小到大 - 冒泡排序
func bubbleSort(a []int32) []int32 {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ { // 后面的是肯定是最大的 所以就不循环了
			if a[j+1] < a[j] { //相邻元素两两对比
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
	return a
}

// 从小到大 -  选择排序
func selectSort(a []int32) []int32 {

	for i := 0; i < len(a)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(a); j++ { //每次都是当前数 与 后面所有数做对比
			if a[j] < a[minIndex] {
				minIndex = j //获取最小数的坐标
			}
		}
		a[minIndex], a[i] = a[i], a[minIndex] //当前最小数 与 当前数 位置对换
	}
	return a
}

func main() {
	a := []int32{1, 5, 2, 7, 4, 8, 6}
	fmt.Println(" ", a)
	fmt.Println("after bubbleSort\n", bubbleSort(a))
	fmt.Println("after selectSort\n", selectSort(a))
}
