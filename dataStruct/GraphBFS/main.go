package main

import "fmt"

/*
	图的广度优先搜索来遍历图
	数据 输入
	5 5
	1 2
	1 3
	1 5
	2 4
	3 5
	输出 12354


*/
func main() {
	var count, head, tail, ec int //顶点的数量、队列的头指针和尾指针
	var queue []int               //存储访问过的城市的队列
	var book [10001]int           //存储顶点状态（是否被访问过）
	var matrix [10001][10001]int  //存储图的邻接矩阵
	fmt.Println("请输入顶点的数量和边的数量（回车确认）")
	fmt.Scanf("%d %d", &count, &ec) //顶点的数量和边的数量
	//初始化存储图的邻接矩阵
	for i := 1; i < 10001; i++ {
		for j := 1; j < 10001; j++ {
			if i == j {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = 99999999
			}
		}
	}
	//把边存入邻接矩阵
	var a, b int //边的顶点序号
	fmt.Println("请输入边（回车确认）")
	for i := 0; i <= ec; i++ {
		fmt.Scanf("%d %d\n", &a, &b) //输入边的顶点
		matrix[a][b] = 1
		matrix[b][a] = 1 //无向图
	}
	//至此，图的存储邻接矩阵完成
	fmt.Println("请输入遍历起始顶点:（回车确认）")
	var start int
	fmt.Scanln(&start)
	head = 1
	tail = 2
	queue[head] = start //初始化队列
	book[start] = 1
	for head < tail {
		for i := 1; i <= count; i++ {
			if book[i] == 1 || queue[head] == i { //当当前节点已经访问过或者当前顶点是头顶点 跳过此次循环
				continue
			}
			if matrix[head][i] == 1 { //如果当前顶点i没有访问过  那么标记访问并入队
				book[i] = 1
				queue[tail] = i
				tail++
			}
		}
		head++
	}
	for i := 1; i <= count; i++ {
		fmt.Printf("%d ", queue[i])
	}
}
