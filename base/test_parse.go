package main
/**
	parse系列：解析 使用strconv
*/
import (
	"fmt"
	"strconv"
)

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
	TestParseBool()
	TestParseInt()
	TestParseUint()
	TestParseFloat()
}
