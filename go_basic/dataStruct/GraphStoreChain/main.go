package main

import "fmt"

type node struct {
	Data []int
	Next *node
}

var (
	queue            []node //存储顶点
	a, b, w          int    //边的起始顶点和结束顶点及权重
	count, edgeCount int    //顶点个数和边数
)

func init() {
	fmt.Println("请输入顶点数和边数")
	fmt.Scanln(&count, &edgeCount)
	queue = make([]node, count+1, count+1)
}

func (n *node) next(new *node) { //递归遍历当前链至最后一个节点
	if n.Next == nil {
		n.Next = new
		return
	}
	n.Next.next(new)
}

// disorderChain 邻接表存储无序图
func disorderChain() {
	fmt.Println("请输入边（起始顶点、结束顶点和权值以空格相隔）")
	for i := 1; i <= edgeCount; i++ {
		fmt.Scanln(&a, &b, &w)
		if queue[a].Data == nil { //如果该边的起始顶点未在数组中，加入数组
			queue[a].Data = []int{a}
		}
		if queue[b].Data == nil { //如果该边的结束顶点未在数组中，加入数组
			queue[b].Data = []int{b}
		}
		var n = &node{
			Data: []int{a, b, w},
		}
		queue[a].next(n) //递归 直到找到当前链的最后一个节点
		queue[b].next(n) //无序图
	}
	printChain()
}

// orderChain 邻接表存储有序图
func orderChain() {
	fmt.Println("请输入边（起始顶点、结束顶点和权值以空格相隔）")
	for i := 1; i <= edgeCount; i++ {
		fmt.Scanln(&a, &b, &w)
		if queue[a].Data == nil { //如果该边的起始顶点未在数组中，加入数组
			queue[i].Data = []int{a}
		}
		if queue[b].Data == nil { //如果该边的结束顶点未在数组中，加入数组
			queue[i].Data = []int{b}
		}
		var n = &node{
			Data: []int{a, b, w},
		}
		queue[a].next(n) //递归 直到找到当前链的最后一个节点
		//有序图
	}
	printChain()
}

// printChain 打印函数
func printChain() {
	var pointer *node
	for i := 1; i <= count; i++ {
		pointer = &queue[i]
		for j := 1; j <= count; j++ {
			if pointer.Next == nil {
				break
			}
			fmt.Printf("%d----%d    w:%d       ", pointer.Next.Data[0], pointer.Next.Data[1], pointer.Next.Data[2])
			pointer = pointer.Next
		}
		fmt.Println()
	}
}
func main() {
	//disorderChain()
	orderChain()
}
