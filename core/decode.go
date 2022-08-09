package core

import "log"

type Decode interface {
	setByte([]byte) Decode
	decode()
}

type JavaDecode struct {
	data []byte
}

func (j *JavaDecode) setByte(bytes []byte) Decode {
	j.data = bytes
	return j
}

func (j *JavaDecode) decode() {
	log.Print(string(j.data))
}
