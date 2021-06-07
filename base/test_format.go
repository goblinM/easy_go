package main

import (
	"fmt"
	"strconv"
)

/**
	format 系列: 格式化，使用strconv
	Sprintf 格式化字符串并赋值给新串
*/

func TestFormatString() {
	// Sprintf 格式化字符串并赋值给新串
	name := "Mary"
	age := 20
	word := "my name is %s, i'm %d years old!"
	var introduce = fmt.Sprintf(word, name, age)
	fmt.Println(introduce)
}

func TestFormatBool(){
	// bool 转字符串
	num := true
	str := strconv.FormatBool(num)
	fmt.Printf("type: %T, value: %v \n", str, str)
}

func TestFormatInt(){
	// int 转 字符串
	var num int64 = 100
	str := strconv.FormatInt(num, 10)
	fmt.Printf("type: %T, value: %v \n", str, str)

}

func TestFormatUint(){
	// uint 转 字符串
	var num uint64 = 110
	str := strconv.FormatUint(num, 10)
	fmt.Printf("type: %T, value: %v \n", str, str)

}

func TestFormatFloat(){
	// float 转 字符串
	var num float64 = 3.1415926
	str := strconv.FormatFloat(num, 'E', -1, 64)
	fmt.Printf("type:%T,value:%v\n ", str, str)
}

func TestAppend(){
	// append
	// 声明一个slice
	b10 := []byte("int (base 10):")
	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}

func main() {
	TestFormatString()
	TestFormatBool()
	TestFormatInt()
	TestFormatUint()
	TestFormatFloat()
	TestAppend()
}
