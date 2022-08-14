package model

import "bytes"

// Bytes Save fix bytes.
type Bytes []uint8

var mappingTag = Bytes{77, 97, 112, 112, 105, 110, 103}

// apiTag The bytes data, The tag to utf-8 string is @RequestMapping
var apiTag = append(Bytes{64, 82, 101, 113, 117, 101, 115, 116}, mappingTag...)

// postTag That change string is '@PostMapping'
var postTag = append(Bytes{64, 80, 111, 115, 116}, mappingTag...)

// getTag 'GetMapping'
var getTag = append(Bytes{64, 71, 101, 116}, mappingTag...)

// blankLine "\n" & ""
var blankLine = Bytes{10}

type Model struct {
}

type ApiModel struct {
	Model
	BasePath       string
	ApiPath        string
	ApiDescription string
}

// BytesToApiModel Resource bytes to ApiModel
func BytesToApiModel(byte []byte) *ApiModel {

	// Determines whether the byte is empty
	if len(byte) > 0 && bytes.Equal(byte, blankLine) {

	}
	return nil
}
