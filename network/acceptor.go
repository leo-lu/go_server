package network

import (
	"log"
	"net"
)

type Acceptor struct {
	*sesMgr
	EventQueue

	listener net.Listener
	running  bool
}

func (self *Acceptor) Start(addr string) Peer {
	ln, err := net.Listen("tcp", addr)
	if nil != err {
		log.Fatalf("#listen failed %v", err.Error())
		return self
	}

	self.listener = ln

	log.Println("#listen ", addr)

	self.running = true

	go func() {
		for self.running {
			conn, err := self.listener.Accept()
			if nil != err {
				log.Fatalf("#accept failed %v", err.Error())
				break
			}

			ses := newSession(NewStream(conn), self.EventQueue, self)

			self.sesMgr.Add(ses)
			ses.OnClose = func() {
				self.sesMgr.Remove(ses)
				log.Println("current session count:", self.sesMgr.SessionCount())
			}
			log.Printf("#accept(%v) sid:%d   current session count:%d", conn.RemoteAddr(), ses.ID(), self.sesMgr.SessionCount())
		}
	}()

	return self
}

func (self *Acceptor) Stop() {
	if !self.running {
		return
	}

	self.running = false
	self.listener.Close()
}

func Stop() {

}

func NewAcceptor() Peer {
	self := &Acceptor{
		sesMgr:  NewSesManager(),
		running: false,
	}
	self.EventQueue = newEventEueue()

	return self
}
