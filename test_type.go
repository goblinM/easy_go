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

func TestParseBool()  {
	// ParseBool ：字符串转bool ， 只接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	str1 := "110"
	boo1, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Printf("str1: %v\n", err)
	} else {
		fmt.Println(boo1)
	}
	str2 := "t"
	boo2, err := strconv.ParseBool(str2)
	if err != nil {
		fmt.Printf("str2: %v\n", err)
	} else {
		fmt.Println(boo2)
	}
}

func TestParseInt(){
	// ParseInt : 返回字符串表示的整数值(包括正负号)
	str := "-11"
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}

func TestParseUint(){
	// ParseUint : 返回字符串表示的整数值(不包括正负号)
	str := "11"
	num, err := strconv.ParseUint(str, 10, 0)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(num)
	}
}

func TestParseFloat(){
	// ParseUint : 浮点数的字符串转换为 float 类型
	str := "3.1415926"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(num)
	}
}

func main() {
	GetType() //显示 a 的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型，a2 类型是 int，IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。
	TestItoa()
	TestAtoi()
	TestParseBool()
}
