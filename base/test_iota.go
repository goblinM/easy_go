package main

import "fmt"

/**
iota 常量生成器以及枚举
周日将对应 0，周一为 1，以此类推。
*/
type Weekday int

const (
	Sunday Weekday = iota  // 生成枚举值,默认为0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	//每次将上一次的值左移一位（二进制位），以得出每一位的常量值。
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)
// 声明芯片类型
type ChipType int
const (
	None ChipType = iota
	CPU    // 中央处理器
	GPU    // 图形处理器
)
func (c ChipType) String() string {
	// 定义 ChipType 类型的方法 String()，返回值为字符串类型。
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
func main() {
	// 输出枚举值
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	// 使用枚举类型并赋初值
	var KK Weekday = Saturday
	fmt.Println(KK)

	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)

	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU)
}
