package main

import (
	"flag"
	"fmt"
)

/**
指针
*/
func GetPtrValue(){
	/**
	获取指针的值
	*/
	var house =  "Hello 1080, 9999"
	// 对字符串取地址 *ptr 类型为string
	ptr := &house
	// 打印ptr类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr指针地址
	fmt.Printf("ptr address: %p\n", &ptr)
	// 获取ptr指针的值
	value := *ptr
	// 打印取值后的类型
	fmt.Printf("value type:%T\n",value)
	// 打印取值后的值
	fmt.Printf("value is:%s\n", value)
}

func swap(a,b *int){
	// 使用指针进行数值交换
	// 取指针a的值赋值给t
	t := *a
	// 取指针b的值赋值给指针a所指的变量
	*a = *b
	// 将指针a的值赋值给指针b
	*b = t

}

func swap2(a, b *int) {
	// 交换的是地址
	// 交换是不成功的。上面代码中的 swap() 函数交换的是 a 和 b 的地址，在交换完毕后，a 和 b 的变量值确实被交换。
	// 但和 a、b 关联的两个变量并没有实际关联。这就像写有两座房子的卡片放在桌上一字摊开，交换两座房子的卡片后并不会对两座房子有任何影响。
	b, a = a, b
}

func TestSwap(){
	a, b := 1, 2
	x, y := 1, 2
	// 交换变量
	swap(&a,&b)
	swap(&x,&y)
	// 输出变量值
	fmt.Println(a, b)
	fmt.Println(x, y)
}

// 定义命令行参数
var mode = flag.String("mode","","process mode")

func FlagPtr()  {
	// go run main.go --mode=fast
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)
}

func NewPtr(){
	// new 指针
	str := new(string)
	*str = "goblinM"
	fmt.Println("str is:", *str)
}

func main() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)
	GetPtrValue()
	TestSwap()
	FlagPtr()
}
