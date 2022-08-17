package core

import (
	"log"
	"testing"
)

type Bytes []byte

var REQUEST_MAPPING = Bytes{64, 82, 101, 113, 117, 101, 115, 116, 77, 97, 112, 112, 105, 110, 103}

var EQUEAL_API = Bytes{64, 82, 101, 113, 117, 101, 115, 116, 77, 97, 112, 112, 105, 110, 103, 104}

func TestByteMapToAscii(t *testing.T) {
	requestMapping := []byte("@RequestMapping")
	postMapping := []byte("@PostMapping")
	getMapping := []byte("@GetMapping")
	mapping := []byte("Mapping")
	blank := []uint8("\n")
	log.Print(requestMapping)
	log.Print(postMapping)
	log.Print(blank)
	log.Print(getMapping)
	log.Print(mapping)
}
