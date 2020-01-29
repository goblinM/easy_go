package main
/**
	类型别名例子
*/
import (
	"fmt"
	"reflect"
)

// 定义商标结构
type Brand struct {
}

// 为商标结构添加show() 方法
func (t Brand) Show(){
}

// 给Brand定义一个别名 FakeBrand
type FakeBrand = Brand

// 定义车结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main() {
	// 变量a声明为Vehicle
	var a Vehicle
	// 调用FakeBrand 的show方法
	a.FakeBrand.Show()

	// 取a的类型反射对象
	ta := reflect.TypeOf(a)

	// 遍历a的所有成员
	for i:=0; i<ta.NumField() ;i++  {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v \n", f.Name, f.Type.Name())
		// FieldName: FakeBrand, FieldType: Brand
		// FieldName: Brand, FieldType: Brand

	}
}
