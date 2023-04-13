package main

import "fmt"

var (
	matrix           [10001][10001]int //邻接矩阵
	count, edgeCount int               //顶点数量和边数量
	a, b, c          int               //起始顶点 结束顶点 权值
)

func init() {
	fmt.Println("请输入顶点数量和边的数量")
	fmt.Scanln(&count, &edgeCount)
	//初始化邻接矩阵
	for i := 1; i <= count; i++ {
		for j := 1; j <= count; j++ {
			if i == j {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = 99999999
			}

		}
	}
}

// disorderMatrix 存储无序图
func disorderMatrix() {
	fmt.Println("请输入边（两顶点和权值以空格相隔）")
	for i := 0; i < edgeCount; i++ {
		fmt.Scanln(&a, &b, &c)
		matrix[b][a] = c
		matrix[a][b] = c
	}
	printMatrix()
}

// orderMatrix 存储有序图
func orderMatrix() {
	fmt.Println("请输入边（两顶点和权值以空格相隔）")
	for i := 0; i < edgeCount; i++ {
		fmt.Scanln(&a, &b, &c)
		matrix[a][b] = c
	}
	printMatrix()
}

// printMatrix 打印邻接矩阵
func printMatrix() {
	for i := 1; i <= count; i++ {
		for j := 1; j <= count; j++ {
			fmt.Printf("%10.f ", float64(matrix[i][j]))
		}
		fmt.Println()
	}
}
func main() {

	//disorderMatrix()
	orderMatrix()
}
