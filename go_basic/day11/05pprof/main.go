package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode(){
	var ch chan int
	select{
	case c := <-ch:
		fmt.Printf("receive from cha,value:%v\n",c)
	default:
		// time.Sleep(500 * time.Millisecond)
	}
}

func main(){
	var isCpuPprof bool
	var isMemoryPprof bool
	flag.BoolVar(&isCpuPprof,"cpup",false,"turn cpu pprof on")
	flag.BoolVar(&isMemoryPprof,"memoryp",false,"turn memory pprof on")

	//解析flag参数
	flag.Parse()
	if isCpuPprof{
		file,err := os.Create("./cpup.pprof")
		if err != nil {
			fmt.Println("create cpu profile file failed:", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	//创建10个线程运行logicCode()
	for i:=0; i<10; i++ {
		go logicCode()
	}
	time.Sleep(time.Second*20)
	if isMemoryPprof{
		file,err :=os.Create("./memoryp.pprof")
		if err !=nil{
			fmt.Println("create memoryp profile file failed! err:",err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}

}