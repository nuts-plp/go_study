package main
import(
	"fmt"
	"time"
	"context"
	"go.etcd.io/etcd/clientv3"
)

func main(){
	cli,err :=clientv3.New(clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout:5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed: %v\n", err)
		return
	}
	fmt.Println("connect to etcd succeeded!")
	defer cli.Close()
 
	
	//watch 
	//返回一个只读通道
	ch :=cli.Watch(context.Background(),"pan")
	//从通道中取值
	for wresp := range ch{
		for _,e := range wresp.Events{
			fmt.Printf("TYPE:%v  key:%v value:%v\n",e.Type,string(e.Kv.Key),string(e.Kv.Value))
		}
	}
}