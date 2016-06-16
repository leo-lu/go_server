package network

import (
	"log"
)

type SesEvent struct {
	*Package
	Ses Session
}

type PeerEvent struct {
	MsgID int
	P     Peer
}

func RegistSesMsg(eq EventQueue, msgId uint32, userHandler func(interface{}, Session)) {
	eq.RegisterCallback(msgId, func(data interface{}) {
		if ev, ok := data.(*SesEvent); ok {
			log.Println("regist ses msg call back.")
			userHandler(ev.Package, ev.Ses)
		} else {
			log.Println("regist ses msg not call back.")
		}

	})
}
