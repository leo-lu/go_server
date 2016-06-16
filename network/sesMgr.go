package network

import (
	"sync"
	"sync/atomic"
)

type sesMgr struct {
	sesMap      map[int64]Session
	sesIDAcc    int64
	sesMapGuard sync.RWMutex
}

func (self *sesMgr) Add(ses Session) {
	self.sesMapGuard.Lock()
	defer self.sesMapGuard.Unlock()

	id := atomic.AddInt64(&self.sesIDAcc, 1)
	ltvSes := ses.(*ltvSession)
	ltvSes.id = id
	self.sesMap[id] = ses
}

func (self *sesMgr) Remove(ses Session) {
	self.sesMapGuard.Lock()
	delete(self.sesMap, ses.ID())
	self.sesMapGuard.Unlock()
}

func (self *sesMgr) GetSession(id int64) Session {
	self.sesMapGuard.RLock()
	defer self.sesMapGuard.RUnlock()

	v, ok := self.sesMap[id]
	if ok {
		return v
	}
	return nil
}

func (self *sesMgr) IterateSession(callback func(Session) bool) {
	self.sesMapGuard.RLock()
	defer self.sesMapGuard.RUnlock()

	for _, ses := range self.sesMap {
		if !callback(ses) {
			break
		}
	}
}

func (self *sesMgr) SessionCount() int {
	self.sesMapGuard.Lock()
	defer self.sesMapGuard.Unlock()

	return len(self.sesMap)
}

func NewSesManager() *sesMgr {
	return &sesMgr{
		sesMap: make(map[int64]Session),
	}
}
