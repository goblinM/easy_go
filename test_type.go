package main

import (
	"fmt"
	"strconv"
)

/**
类型别名
*/
// 将NewInt 定义为int 类型
// 常见的定义类型的方法，通过 type 关键字的定义，NewInt 会形成一种新的类型，NewInt 本身依然具备 int 类型的特性
type NewInt int

// 将int 取别名为IntAlias
type IntAlias = int

func GetType()  {
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type:%T \n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
}

func TestItoa(){
	// Itoa : 整型转字符型
	num := 100
	strNum := strconv.Itoa(num)
	fmt.Printf("type:%T  value:%#v\n", strNum, strNum)
}

func TestAtoi(){
	// Atoi : 字符串转整形
	str1 := "110"
	str2 := "s100"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败！", str1)
	} else {
		fmt.Printf("type:%T value:%#v\n", num1, num1)
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败！", str2)
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)
	}
}

func main() {
	GetType() //显示 a 的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型，a2 类型是 int，IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。
	TestItoa()
	TestAtoi()
	TestParseBool()
	TestParseInt()
	TestParseUint()
	TestParseFloat()
}
