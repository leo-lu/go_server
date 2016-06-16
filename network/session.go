package network

import (
	"log"
	//"sync"
)

type Session interface {
	ID() int64
	Send(interface{})
	RawSend(*Package)
	Close()
}

type ltvSession struct {
	//writeChan chan interface{}
	stream  Stream
	id      int64
	OnClose func()
}

func (self *ltvSession) ID() int64 {
	return self.id
}

func (self *ltvSession) recvThread(eq EventQueue) {
	var err error
	var pack *Package
	for {
		pack, err = self.stream.Read()
		if nil != err {
			//conn cl
			log.Println("recvThread err:", err.Error())

			if self.OnClose != nil {
				self.OnClose()
			}
			break
		}

		eq.PostData(pack)
	}
}

func (self *ltvSession) RawSend(pack *Package) {

}

func (self *ltvSession) Send(data interface{}) {

}

func (self *ltvSession) Close() {

}

func newSession(stream Stream, eq EventQueue, p Peer) *ltvSession {
	self := &ltvSession{
		stream: stream,
	}

	go self.recvThread(eq)

	return self
}
