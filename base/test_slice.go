package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
func main() {
	// 当前长度为3，最多可存10个
	numbers := make([]int, 3, 10)
	printSlice(numbers)
	// 空指针
	var nil_numbers []int
	if nil_numbers == nil {
		fmt.Printf("slice is nil\n")
	}
	// append(): 添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)
	// 添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	// 拷贝元素
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers) * 2))
	// 拷贝numbers的内容到numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)
	// {},{}
	res := map[string]int{"count": 1, "age": 24}
	res2 := map[string]string{"name": "alice"}
	type result struct {
		count int
		age   int
		name  string
	}
	//data := make([]result, 0)
	data := append(res, res2)

}
