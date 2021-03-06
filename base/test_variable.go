package main

import (
	"fmt"
	"math"
)

/*
变量类型的测试
*/

func simpleShortDecarationRight(){
	// 短变量声明并且初始化正确使用
	i, j := 0, 1
	fmt.Println("i=",i)
	fmt.Println("j=", j)
}

func SimpleShortDecarationWrong(){
	// 短变量声明并且初始化错误使用
	var testVal string
	//testVal := "hello world"
	fmt.Printf(testVal)
}

func otherVariableSet(){
	var (
		a int
		b string
		c [3] float32
		//d func() bool
		//e struct{
		//	x int
		//}
	)
	b = "hello world"
	fmt.Printf(b)
	a = 666
	fmt.Println(a)
	for i:=0; i<len(c); i++ {
		c[i] += 1
		fmt.Println(c[i])
	}
	fmt.Println(c)
	//
	//fmt.Println(d)
}

func SwitchVariables(){
	// 变量交换
	// 方法一
	a := 100
	b := 200
	var t int
	fmt.Println("交换前(a,b)的值是：",a,b)
	t = a
	a = b
	b = t
	fmt.Println("交换后(a,b)的值是：",a,b)
	// 方法二
	//a, b = b, a
	//fmt.Println("交换后(a,b)的值是：",a,b)
}

func SecretVariableSet() (int, int){
	// 设置返回的值
	return 100, 200
}

func SecretVariableGet(){
	// 匿名变量
	_, a := SecretVariableSet()
	b, _ := SecretVariableSet()
	fmt.Println(a,b)
}

func sum(a, b int) int {
	// 形式参数
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

func FloatVariable(){
	// 浮点类型 Printf 格式化“%f”来控制保留几位小数
	var n1 float64 = 1352424.67
	fmt.Println(n1)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
}

func ComplexVariable(){
	// 复数的运算
	var x complex128 = complex(1,2)
	var y complex128 = complex(2,1)
	c := real(x+y) // x+y = 3+ 3i  real 取实部
	d := imag(x-y) // x-y = -1+i imag 取虚部
	e := real(x*y) // x*y = (2-2)+(i+4i) = 0+5i
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func SwitchBoolVariable(b bool) int{
	// 布尔值转换为0,1
	if b{
		return 1
	}
	return 0
}

func StringVariable(){
	// 字符串
	a := "hello"
	b := "world"
	// + 号拼接
	c := a+b
	fmt.Println(c)
	// 多行字符串 使用反引号 `
	multipleWord := `
	I am James
	I come from China
	I like to study
	`
	fmt.Println(multipleWord)
}

func TypeTranslate(){
	// 类型转换
	a := 2.0
	b := int(a)
	fmt.Println(b)
}

func TestCommit(){
	fmt.Println("github commit")
}

func TestIota() {
	/**
	* iota ，特殊常量，可以认为是一个可以被编译器修改的常量。
	iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
	简单表述:
	i=1：左移 0 位,不变仍为 1;
	j=3：左移 1 位,变为二进制 110, 即 6;
	k=3：左移 2 位,变为二进制 1100, 即 12;
	l=3：左移 3 位,变为二进制 11000,即 24。
	注：<<n==*(2^n)。
	*/

	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)
	fmt.Println("i=", i)  // 1
	fmt.Println("j=", j)  // 6
	fmt.Println("k=", k)  // 12
	fmt.Println("l=", l)  // 24
}

func main() {
	//otherVariableSet()
	//simpleShortDecarationRight()
	//simple_short_decaration_wrong()
	//SwitchVariables()
	//SecretVariableGet()
	//sum(3,4)
	//FloatVariable()
	//ComplexVariable()
	//SwitchBoolVariable(true)
	//StringVariable()
	//TypeTranslate()
	TestIota()

}
