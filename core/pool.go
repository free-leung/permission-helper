package core

import (
	"bytes"
	"sync"
)

// CreatSynPoolBytes Create syn.Pool to manager bytes.
func CreatSynPoolBytes(maxSize int64) sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			return make([]byte, maxSize)
		},
	}
}

// CreateStringPool Create type of string.
func CreateStringPool() sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			row := ""
			return row
		},
	}
}

func CreateByteBufferPool() sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			return bytes.Buffer{}
		},
	}
}
