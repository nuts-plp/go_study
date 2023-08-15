package session

//SessionMgr 定义管理者管理所有的session
type SessionMgr interface {
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
