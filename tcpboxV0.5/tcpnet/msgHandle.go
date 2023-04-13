package tcpnet

import (
	"fmt"
	"strconv"
	"tcpbox/tcpiface"
)

type MsgHandler struct {
	//属性 根据消息ID调度路由
	Apis map[uint32]tcpiface.IRouter
}

//初始化创建MsgHandler的处理逻辑
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]tcpiface.IRouter),
	}
}

func (mh *MsgHandler) DoMsgHandle(request tcpiface.IRequest) {

	//从request中找到msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID=", request.GetMsgID(), "is not dound! need register")
	}

	//根据msgID调度对应的业务处理
	//handler.PreHandle(request)
	handler.Handle(request)
	//handler.PostHandle(request)
}

//为消息添加具体的处理逻辑
func (mh *MsgHandler) AddRouter(msgID uint32, router tcpiface.IRouter) {
	//判断当前msg绑定的api处理方法事都已经存在
	if _, ok := mh.Apis[msgID]; ok {
		//id 已经注册了
		panic("repeat api,msgID=" + strconv.Itoa(int(msgID)))
	}
	//添加msg与api的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID=", msgID, "successed!")
}
