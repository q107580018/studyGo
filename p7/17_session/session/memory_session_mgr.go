package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}
func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}
func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	// 用uuid作为sessionID,转string
	sessionId := uuid.NewV4().String()
	// 创建session
	session = NewMemorySession(sessionId)
	return
}
