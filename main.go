/*
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

// 矩阵变图
func CreatImage(outer [][]uint8) {
	// 创建图像文件
	file, err := os.Create("tag.jpeg")
	// 抛错
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 新建图片，并定义大小
	alpha := image.NewAlpha(image.Rect(0, 0, 400, 400))
	// 将切片中的信息填入图片中
	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			alpha.Set(x, y, color.Alpha{outer[x][y]})
		}
	}
	// 生成图片
	jpeg.Encode(file, alpha, nil)

}

// 将一行数据转化为一个类似于400*400的矩阵的切片
func Matrix(data []uint8) [][]uint8 {

	outer := make([][]uint8, 400)
	for o := range outer {

		inner := make([]uint8, 400)
		for i := range inner {

			// 分析图像的列信息，并装载进inner中
			if data[400*i+o] == 49 {
				inner[i] = 255
			} else {
				inner[i] = 0
			}

		}
		// 装载有一列图像信息的inner赋值给outer中的一个元素
		outer[o] = inner
	}
	// 返回的outer就是一个元素是切片类型的切片
	return outer
}

// 从文件中读取数据，有好几种文件打开读取方式。选择了最简便的。
func ReadFile(filepath string) []uint8 {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
		return nil
	}
	return data
}

func main() {
	s := ReadFile("./1.txt")
	CreatImage(Matrix(s))
}
*/

package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(s string) []uint8 {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func Matrix(data []uint8, x int, y int) [][]uint8 {
	outer := make([][]uint8, x)
	for o := range outer {
		inner := make([]uint8, y)
		for i := range inner {
			if data[x*i+o] == 49 {
				inner[i] = 255
			} else {
				inner[i] = 0
			}
		}
		outer[o] = inner
	}
	return outer
}

func CreatImage(matrix [][]uint8) {
	file, err := os.Create("tag.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	alpha := image.NewAlpha(image.Rect(0, 0, len(matrix), len(matrix[0])))
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			alpha.Set(x, y, color.Alpha{matrix[x][y]})
		}
	}
	jpeg.Encode(file, alpha, nil)
}

func main() {
	s := ReadFile("./1.txt")
	CreatImage(Matrix(s, 400, 400))
}
