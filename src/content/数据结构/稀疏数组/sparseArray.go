package main

import "fmt"

func main() {
	// 创建一个11 × 11 数组：
	var Map [11][11]int

	// 记录位置：
	Map[1][2] = 1 //使用1表示黑子位置
	Map[2][3] = 2 //使用2表示白子位置

	// 循环遍历数组：
	for _, v := range Map {
		for _, v2 := range v {
			fmt.Printf("%d  ", v2)
		}
		fmt.Println()
	}

	// 定义节点；
	type ValueNode struct {
		row int
		col int
		val int
	}

	/*
		转换稀疏数组思路：
		1. 遍历Map，如果发现一个数组的某一个元素不为0,就为这个元素创建一个节点
		2. 将其保存在这个节点中（切片）
	*/
	var sparseArr []ValueNode

	// 标准的稀疏数组还应该记录原始二维数组的shape
	// 创建一个节点，存放原始信息：
	valNode := ValueNode{
		row: 11,
		col: 11,
		val: 0,
	}

	// 把原始信息节点存放到切片中：
	sparseArr = append(sparseArr, valNode)

	// _ 这个下划线表示占位符
	for i, rows := range Map {
		for j, v := range rows {
			if v != 0 {
				// 创建元素值结点：
				valNode := ValueNode{
					row: i,
					col: j,
					val: v,
				}
				// 向切片中追加元素：
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组：
	fmt.Println("稀疏数组：")
	for i, valNode := range sparseArr {
		fmt.Println(i, "---", valNode)
	}
}
