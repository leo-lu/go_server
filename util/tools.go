package util

import (
	"log"
	"runtime"
)

type Tools struct {
}

func ShowStackInfo() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		log.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
}

func GetNetworkBigLittle() int {
	x := 0x1234
	p := unsafe.Pointer(&x)
	p2 := (*[N]byte)(p)
	if p2[0] == 0 {
		return 1
	} else {
		return 0
	}
}
