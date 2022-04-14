package tempconv

// 对于每一个类型T，都有一个对应的类型转换操作T(x)，比如：Celsius(0) , 将0(int)转化为0(Celsius)
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// 如果两个值有着不同的类型，则不能直接进行比较(就算它们的底层类型一样)  比如：BoilingC不能与CToF(BoilingC)比较
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
