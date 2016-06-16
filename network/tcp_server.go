package network

import (
	"fmt"
	"net"
	//"sync"
	//"time"
	"log"
	//"encoding/binary"
	//"bytes"
	//flatbuffers "github.com/google/flatbuffers/go"
	MyGame "../generated"
)

type TcpServer struct {
	Addr       string
	MaxConnNum int

	listener net.Listener
	running  bool
}

func (self *TcpServer) Start(addr string) {
	log.Println("#Tcp Server Start")

	ln, err := net.Listen("tcp", addr)
	if nil != err {
		log.Fatalf("#listen failed %v", err.Error())
		return
	}

	self.listener = ln

	log.Println("#listen ", addr)

	self.running = true

	defer func() {
		ln.Close()
		self.running = false
	}()

	//
	for self.running {
		log.Println("Accept......")

		conn, err := self.listener.Accept()
		if nil != err {
			log.Fatalf("#accept failed %v", err.Error())
			break
		}

		go handel(conn)
	}

	log.Println("Tcp Server Stop.")
	//time.Sleep(2 * time.Second)
}

func handel(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		//pack:= "hello world. 中文"
		//str := fmt.Sprintf("%d%s", len([]byte(pack)) , pack)
		//log.Println([]byte(str))
		//conn.Write([]byte(str))

		n, err := conn.Read(buf)
		if nil != err {
			log.Println("#read failed ", err.Error())
			conn.Close()
			return
		}
		//log.Println(string(buf[0:n]))

		msg := fmt.Sprintf("hello  %s \n", string(buf[0:n]))
		conn.Write([]byte(msg))

		obj := MyGame.GetRootAsMonster(buf, 0)

		//log.Println("buffer:", obj)
		log.Println("Name:", string(obj.Name()))
		log.Println("mana:", obj.Mana())
		log.Println("Hp:", obj.Hp())
		pos := *obj.Pos(nil)
		log.Println("Pos:", pos.X())
		log.Println("Pos:", pos.Y())
		log.Println("Pos:", pos.Z())
		log.Println("Color:", obj.Color())
	}
}
