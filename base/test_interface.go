package main

import "fmt"

/**接口*/
/*定义接口*/
type Animal interface {
	// 接口方法
	colorType()
	legs()
	animalType()
}
/*定义结构体*/
type Cat struct {
	name string
	leg int
	color string
}

type Dog struct {
	name string
	leg int
	color string
}
/**实现接口方法*/
func (cat Cat) animalType() {
	fmt.Println("name = ", cat.name)
}
func (cat Cat) legs() {
	fmt.Println("leg = ", cat.leg)
}
func (cat Cat) colorType() {
	fmt.Println("color = ", cat.color)
}

func (dog Dog) animalType() {
	fmt.Println("name = ", dog.name)
}
func (dog Dog) legs() {
	fmt.Println("leg = ", dog.leg)
}
func (dog Dog) colorType() {
	fmt.Println("color = ", dog.color)
}

func main() {
	var animal Animal
	//cat := Cat{name:"cuteM", leg:4, color:"white"}
	animal = Cat{name:"cuteM", leg:4, color:"white"}
	animal.animalType()
	animal.colorType()
	animal.legs()

	animal = Dog{name:"luckyDog", leg:4, color:"black"}
	animal.animalType()
	animal.colorType()
	animal.legs()
}
