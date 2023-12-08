package main

/*
	图的深度优先搜索
*/
import "fmt"

var (
	book   [101]int
	sum, n int //sum记录访问过的顶点数 n表示顶点数量
	e      [101][101]int
)

func dfs(cur int) {
	var i int
	fmt.Printf("%d", cur)
	sum++ //每访问一个点，sum就加一
	if sum == n {
		return //所有节点都已访问过就退出
	}
	for i = 1; i <= n; i++ { //从1号到n号顶点依次尝试，看哪些点与当前顶点cur相连
		//判断当前顶点cur到顶点i是否有边，并判断顶点i是否已访问过
		if e[cur][i] == 1 && book[i] == 0 {
			book[i] = 1 // 标记顶点i已经访问过
			dfs(i)      //从顶点i再出发继续便利
		}
	}
	return
}
func main() {
	var i, j, m, a, b int
	c, err := fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		fmt.Println(err, c)
		return
	}

	//初始化二维矩阵
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if i == j {
				e[i][j] = 0
			} else {
				e[i][j] = 99999999 //这里假设99999999为正无穷
			}
		}
	}
	//读入顶点之间的边
	for i := 1; i <= m; i++ {
		c, err = fmt.Scanln(&a, &b)
		if err != nil {
			fmt.Println(err, c)
			return
		}
		e[a][b] = 1
		e[b][a] = 1 //这里是无向图，所以将e[b][a]=1
	}
	//从1号城市出发
	book[1] = 1 //标记一号顶点已访问
	dfs(1)      //从1号顶点开始遍历
	return
}
