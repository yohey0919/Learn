package popcount

// pc[i] is the population count of i.
var pc [256]byte

// 数据初始化并不是一个简单的赋值过程，可以使用init函数
// init初始化函数除了不能被调用或引用外，其他行为和普通函数类似

// pc表格用于处理每个8bit宽度的数字含二进制的1bit的bit个数
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 用于返回一个数字中含二进制1bit的个数
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
