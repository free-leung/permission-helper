package core

import (
	"bufio"
	"io"
	"log"
	"sync"
)

var waitGroup sync.WaitGroup

type Reader interface {
	read(fileCtx *fileCtx)
}

type ReadCtx struct {
	MaxReadSize int64
	data        []byte
}

func (rctx ReadCtx) read(fileCtx *fileCtx) {
	file := fileCtx.osFile
	// UPDATE: close after checking error
	defer file.Close() //Do not forget to close the file
	reader := bufio.NewReader(file)
	for {
		// Create max read byte. But that byte possible not enough, so we can chunk to read.
		synPool := CreatSynPoolBytes(fileCtx.Size)
		// Get clean byte from the pool
		readBytes := synPool.Get().([]byte)
		// Do read and that read data send in readBytes.
		len, err := reader.Read(readBytes)
		rctx.data = readBytes[:len]
		// If read len just 0, That means we got the tail or cause err.
		if len == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Err to read file: ", err)
				break
			}
		}
		waitGroup.Add(1)
		go func() {
			fileCtx.decode.setByte(rctx.data).decode()
			waitGroup.Done()
		}()
	}
	// Wait all files read complete
	waitGroup.Wait()
}
