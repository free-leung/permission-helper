package core

import (
	"bytes"
	"io"
	"log"
)

// Decoder To decode bytes.
// Currently only JavaDecoder is available.
type Decoder interface {
	decode([]byte) interface{}
}

type JavaDecoderDelegator interface {
	Decoder
}

// JavaDecoder decode .java file.
// It's going to parse out a Model.
// This model require extends Model.struct.
type JavaDecoder struct {
	delegator JavaDecoderDelegator
}

func (j *JavaDecoder) decode(data []byte) interface{} {
	return j.delegator.decode(data)
}

type JavaApiDecoder struct {
}

// decode The function from JavaApiDecoder.
// It's not easy to read
func (apiDecoder *JavaApiDecoder) decode(data []byte) interface{} {
	// Read bytes.
	bufferReaders := bytes.NewBuffer(data)

	// Create byteBuffer by syn pool.
	bytesBufferPool := CreateByteBufferPool()
	bytesBuffer := bytesBufferPool.Get().(bytes.Buffer)
	for {
		content, err := bufferReaders.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Print("Read bytes completes!")
				break
			} else {
				log.Fatal("Read line content err: ", err)
			}
		}
		bytesBuffer.Write(content)
		log.Print(bytesBuffer.Bytes())
		log.Print(bytesBuffer.String())
		// '\n' cut a line
		bytesBuffer.Reset()
	}
	return nil
}
