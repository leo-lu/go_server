package network

import (
//"time"
//"log"
)

type EventQueue interface {
	RegisterCallback(id uint32, f func(interface{}))
	InjectData(func(interface{}) bool)
	PostData(data interface{})
	CallData(data interface{})
}

type evQueuc struct {
	contextMap      map[uint32][]func(interface{})
	queue           chan interface{}
	inject          func(interface{}) bool
	concurrenceMode bool
}

func (self *evQueuc) RegisterCallback(id uint32, f func(interface{})) {
	//
	em, ok := self.contextMap[id]
	if !ok {
		em = make([]func(interface{}), 0)
	}
	em = append(em, f)
	self.contextMap[id] = em
}

func (self *evQueuc) InjectData(f func(interface{}) bool) {
	self.inject = f
}

func (self *evQueuc) Exists(id uint32) bool {
	_, ok := self.contextMap[id]
	return ok
}

func (self *evQueuc) PostData(data interface{}) {
	if self.concurrenceMode {
		self.CallData(data)
	} else {
		self.queue <- data
	}
}

type contentIndexer interface {
	ContextID() uint32
}

func (self evQueuc) CallData(data interface{}) {
	if nil != self.inject && !self.inject(data) {
		return
	}

	switch d := data.(type) {
	case contentIndexer:
		if cArr, ok := self.contextMap[d.ContextID()]; ok {
			for _, c := range cArr {
				c(data)
			}
		}
	case func():
		d()
	default:

	}
}

const queueLen = 10

func newEventEueue() *evQueuc {
	self := &evQueuc{
		contextMap: make(map[uint32][]func(interface{})),
		queue:      make(chan interface{}, queueLen),
	}
	self.concurrenceMode = true
	return self
}
