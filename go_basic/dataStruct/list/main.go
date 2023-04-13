package main

import (
	"fmt"
	"go_basic/dataStruct/list/seqListMgr_slice"
)

func main() {
	seqlist := seqListMgr_slice.NewSeqListMgr_slice(5, false)
	seqlist.Add("潘丽萍")
	seqlist.Add("周小林")
	seqlist.Add("樊雪怡")
	seqlist.Add("于欢")
	seqlist.Add("牛敏")
	fmt.Println(seqlist.GetSize())
	seqlist.ToString()
	seqlist.Delete(0)
	seqlist.ToString()
	fmt.Println(seqlist.GetSize())
	seqlist.Modify(3, "joy")
	seqlist.ToString()
	//fmt.Println("----------------------------------------")
	//list := seqListMgr_map.NewSeqList_map(5)
	//list.Add("mylover", "潘丽萍")
	//list.Add("myprefix", "周小林")
	//list.ToString()
	//fmt.Println(list.GetSize())
	//list.Modify("myprefix", "于欢")
	//list.ToString()
	//list.Delete("myprefix")
	//fmt.Println(list.GetSize())
	//fmt.Println()
	//fmt.Println("==================================================")
	//linkedList := linkedListMgr.NewLinkedListMgr()
	//linkedList.Add("潘丽萍")
	//linkedList.Add("周小林")
	//linkedList.Add("于欢")
	//linkedList.Add("牛敏")
	//linkedList.ToArray()
	//fmt.Println(linkedList.GetSize())
	//linkedList.Delete(3)
	//linkedList.ToArray()
	//linkedList.Modify(3, "于欢")
	//fmt.Println(linkedList.GetSize())
	//linkedList.ToArray()

}
