package taillog
import(
	"github.com/hpcloud/tail"
	"fmt"
	"time"
)

var(
	tailObj *tail.Tail
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:true,
		Follow:true,
		Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll:true,
	}
	tailObj,err = tail.TailFile(fileName,config)
	if err != nil {
		fmt.Println("tail file failed! err: ", err)
		return
	}
	return
}
func ReadLog()<-chan *tail.Line{
	return tail.Lines
}