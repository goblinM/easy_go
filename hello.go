package main

import (
	"fmt"
	"strconv"
)

func TestFor() {
	sum := 0
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)

	strings := []string{"google", "easygo"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}

func main() {
	na := 1
	var k = 1
	fmt.Print(k)
	fmt.Println(strconv.Itoa(na))
	fmt.Println("Hello World")
}
