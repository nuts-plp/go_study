package tcpnet

type Message struct {
	//消息的ID
	MsgID uint32

	//消息的长度
	MsgLen uint32

	//消息的内容
	MsgData []byte
}

//获取消息的ID
func (m *Message) GetMsgID() uint32 {
	return m.MsgID
}

//获取消息的长度
func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}

//获取消息的内容
func (m *Message) GetMsgData() []byte {
	return m.MsgData
}

//设置消息的ID
func (m *Message) SetMsgID(id uint32) {
	m.MsgID = id
}

//设置消息的长度
func (m *Message) SetMsgLen(len uint32) {
	m.MsgLen = len
}

//设置消息的内容
func (m *Message) SetMsgData(data []byte) {
	m.MsgData = data
}
