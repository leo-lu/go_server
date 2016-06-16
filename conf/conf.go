package conf

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"os"
)

type ServerConf  struct {
	Tcp struct {
		On 		bool
		Host 	string
		Port 	int
	}
	Http struct {
		On 		bool
		Host 	string
		Port 	int
	}
}

var confInfo ServerConf

func (conf *ServerConf) Init() ServerConf {
	f, err :=  os.Open("./conf/server.conf")
	if nil != err {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	
	log.Println("conf json:: \r\n",  string(data))

	err = json.Unmarshal(data, &confInfo)
	if err != nil {
		log.Fatalln(err)
	}

	return confInfo
}

func (conf *ServerConf) GetConf() *ServerConf {
	return &confInfo
}