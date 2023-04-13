package main

import "fmt"

/*
	无向图的深度优先搜索 带权边 邻接矩阵存储    邻接表存储的深度搜索在对应的邻接表存储中
*/
var (
	stack, book      []int //一个栈和一个数组  深度优先搜索时使用、
	queue            []int //队列 广度优先搜索时使用
	head             int   //队列的指针
	count, edgeCount int   //顶点个数 边数
	matrix           [1001][1001]int
	tail             int //队列的指针
	a, b, w          int //分别是边的俩、权重
	start            int //搜索的起点
	sum              int //访问过的顶点的数量
)

// init 借助ini函数的特性实现邻接矩阵的初始化
func init() {
	fmt.Println("请输入图的顶点数和边数（以空格分隔）")
	fmt.Scanln(&count, &edgeCount)
	book = make([]int, count+1, count+1)
	stack = make([]int, count+1, count+1)
	tail = 0
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
	matrixPrint()
}

// stackPrint 打印遍历序列
func stackPrint() {
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", stack[i])
	}
}

// matrixPrint 打印邻接矩阵
func matrixPrint() {
	for i := 1; i <= count; i++ {
		for j := 1; j <= count; j++ {
			fmt.Printf("%10.f", float64(matrix[i][j]))
		}
		fmt.Println()
	}
}

// initMatrix 把图存入邻接矩阵
func initMatrix() {

	fmt.Println("请输入边的顶点和权重（空格隔开）")
	for i := 1; i <= edgeCount; i++ {
		fmt.Scanln(&a, &b, &w)
		matrix[a][b] = w
		matrix[b][a] = w //无序   如果为有向图就把此语句删除
	}
}
func DFS() {
	fmt.Println("请输入深度优先搜索的起点")
	fmt.Scanln(&start)
	stack[tail] = start
	tail++
	book[start] = 1
	sum = 1
	dfs(start)
	fmt.Println("打印存储图的邻接矩阵")
	matrixPrint()
	fmt.Println("打印遍历序列")
	stackPrint()
}
func dfs(cur int) {
	var i int
	fmt.Printf("%d", cur)
	sum++ //每访问一个点，sum就加一
	if sum == count {
		return //所有节点都已访问过就退出
	}
	for i = 1; i <= count; i++ { //从1号到n号顶点依次尝试，看哪些点与当前顶点cur相连
		//判断当前顶点cur到顶点i是否有边，并判断顶点i是否已访问过
		if matrix[cur][i] == 1 && book[i] == 0 {
			book[i] = 1 // 标记顶点i已经访问过
			dfs(i)      //从顶点i再出发继续便利
		}
	}
	return
}

func main() {
	initMatrix()
	DFS()

}
