package main
/**
变量逃逸
**/
import "fmt"

// 本函数测试入口参数和返回值情况
func dummy(b int) int {

	// 声明一个变量c并赋值
	var c int
	c = b

	return c
}

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy2() *Data {
	// 实例化c为Data类型
	var c Data
	//返回函数局部变量地址
	return &c
}

// 空函数, 什么也不做
func void() {

}

func main() {
	// 运行使用命令：go run -gcflags "-m -l" test_escape.go
	// 声明a变量并打印
	var a int
	// 调用void()函数
	void()

	// 打印a变量的值和dummy()函数返回
	fmt.Println(a, dummy(0))
	fmt.Println(dummy2())
}
