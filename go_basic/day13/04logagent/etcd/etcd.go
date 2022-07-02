package etcd
import(
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
	

)

func Init(addrs []string,timeout int)(err error){
	
	cli,err :=clientv3.New(clientv3.Config{
		Endpoints: addrs,
		DialTimeout:time.Duration(timeout)*time.Second ,
	})

	if err != nil {
		fmt.Printf("connect to etcd failed! err: %v\n", err)
		return
	}
	fmt.Println("connect to etcd succeeded!")
	defer cli.Close()
	//put
	ctx,cancel :=context.WithTimeout(context.Background(),time.Second)
	_,err =cli.Put(ctx,"zhou","love")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed! err: %v\n", err)
		return
	}
	//get
	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	resp,err:=cli.Get(ctx,"zhou")
	cancel()
	if err!= nil{
		fmt.Printf("get from etcd failed ! err:%v\n",err)
		return
	}
	for _,ev := range resp.Kvs{
		fmt.Printf("%s:%s\n",ev.Key,ev.Value)
	}
	return
}