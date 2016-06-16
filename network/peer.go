package network

type Peer interface {
	Start(addr string) Peer
	Stop()
	//SetName(string)
	//Name() string

	EventQueue
	SessionManager
}

type SessionManager interface {
	GetSession(int64) Session
	IterateSession(func(Session) bool)
	SessionCount() int
}
