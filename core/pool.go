package core

import "sync"

// CreatSynPoolBytes Create syn.Pool to manager bytes.
func CreatSynPoolBytes(maxSize int64) sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			return make([]byte, maxSize)
		},
	}
}
