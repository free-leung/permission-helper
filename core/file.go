package core

import (
	"log"
	"os"
)

type FileContext interface {
	fileName() string
	filePath() string
	isDir() bool
	getFiles() []os.File
	size() int64
}

type fileCtx struct {
	FilePath string
	FileName string
	IsDir    bool
	Files    []os.File
	Size     int64
	osFile   *os.File
	decode   Decode
}

func (f *fileCtx) fileName() string {
	return f.FileName
}

func (f *fileCtx) filePath() string {
	return f.FilePath
}

func (f *fileCtx) isDir() bool {
	return f.IsDir
}

func (f *fileCtx) getFiles() []os.File {
	return f.Files
}

func (f *fileCtx) size() int64 {
	return f.Size
}

func NewFileCtx(completePath string, decode Decode) *fileCtx {
	file, err := os.Open(completePath)
	if err != nil {
		log.Fatal("Error to open file : ", completePath, "\n Error msg : ", err)
	}
	info, _ := file.Stat()
	ctx := &fileCtx{
		FilePath: file.Name(),
		FileName: info.Name(),
		Size:     info.Size(),
		IsDir:    info.IsDir(),
		osFile:   file,
		decode:   decode,
	}
	return ctx
}
