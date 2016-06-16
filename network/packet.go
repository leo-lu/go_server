package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
)

const (
	HeaderSize  = 6 //msgID(uint32) + size(uint16)
	MaxPackSize = 1024 * 8
)

//
type Package struct {
	MsgID uint32
	Data  []byte
}

func (self Package) ContextID() uint32 {
	return self.MsgID
}

func (self Package) Parse() (msg interface{}, err error) {
	msg = Package{}
	return
}

type Stream interface {
	Read() (*Package, error)
	Write(pkt *Package) error
}

type ltvStream struct {
	conn net.Conn
}

var (
	packTooBig      = errors.New("ReadPacket: package too big.")
	packSizeInvalid = errors.New("ReadPacket: package invalid size.")
)

func (self *ltvStream) Read() (p *Package, err error) {
	hData := make([]byte, HeaderSize)
	//read head data from network
	if _, err = io.ReadFull(self.conn, hData); nil != err {
		return nil, err
	}

	log.Println(hData)

	p = &Package{}

	hBuf := bytes.NewReader(hData)
	//read msgID
	if err = binary.Read(hBuf, binary.LittleEndian, &p.MsgID); nil != err {
		return nil, err
	}
	//log.Println("msgID:", p.MsgID)
	//read pack size
	var pSize uint16
	if err = binary.Read(hBuf, binary.LittleEndian, &pSize); nil != err {
		return nil, err
	}
	//log.Println("msg Size:", pSize)
	//pack too big
	if pSize > MaxPackSize {
		return nil, packTooBig
	}

	/*dSize := pSize - HeaderSize*/
	if 0 > pSize {
		return nil, packSizeInvalid
	}

	p.Data = make([]byte, pSize)
	if _, err = io.ReadFull(self.conn, p.Data); nil != err {
		return nil, err
	}

	return
}

func (self *ltvStream) Write(p *Package) (err error) {
	hBuf := bytes.NewBuffer([]byte{})
	//write MsgID to buf
	if err = binary.Write(hBuf, binary.LittleEndian, p.MsgID); nil != err {
		return
	}
	//write pack size to buff
	if err = binary.Write(hBuf, binary.LittleEndian, uint16(len(p.Data))); nil != err {
		return
	}
	//send head to network
	if _, err = self.conn.Write(hBuf.Bytes()); nil != err {
		return
	}
	//send pack data to network
	if _, err = self.conn.Write(p.Data); nil != err {
		return
	}
	return
}

func NewStream(conn net.Conn) Stream {
	return &ltvStream{
		conn: conn,
	}
}
