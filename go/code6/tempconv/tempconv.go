//自定义了一个tempconv包，这个包, 包含Celsius, Fahrenheit的计算
package tempconv

import "fmt"

// 一个类型声明语句,创建了一个新的类型名称，和现有类型具有相同的底层结构
// 格式：type 类型名字 底层类型
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

// 许多类型都会定义一个String方法(该方法类似与JAVA的toString())
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
