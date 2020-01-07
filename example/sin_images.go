package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)
/**
实例一：输出正弦图像
*/
func main() {
	// 图片大小 const 常量变量
	const size  = 300
	// 创建灰度图 初始化好的灰度图默认的灰度值都是 0，对的是黑色
	pic := image.NewGray(image.Rect(0,0,size,size))
	// 遍历像素
	for x:= 0; x < size; x++{
		for y:=0; y < size; y++{
			// 填充为白色
			//pic.Set(x,y,color.White)
			pic.SetGray(x,y,color.Gray{255})
		}
	}

	// 从0到最大像素生成x坐标  正弦函数是一个周期函数，定义域是实数集，取值范围是 [-1, 1]。
	// 用编程的通俗易懂的话来说就是：math.Sin 函数的参数支持任意浮点数范围，函数返回值的范围总是在 -1～1 之间（两端包含）。
	//要将正弦函数放在图片上需要考虑以下一些因素：
	//math.Sin 的返回值在 -1～1 之间，需要考虑将正弦的输出幅度变大，可以将 math.Sin 的返回值乘以一个常量进行放大。
	//图片的坐标系原点在左上角，而 math.Sin 基于笛卡尔坐标系原点在左下角，需要对图像进行上下翻转和平移。

	for x:= 0; x < size; x++{
		// 让sin的值的范围在0~2Pi之间
		s := float64(x)*2*math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x,int(y),color.Gray{0})
	}

	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil{
		log.Fatal(err)
	}
	// 使用png 格式写入数据到文件中
	png.Encode(file, pic)
	// 关闭文件
	file.Close()
}
