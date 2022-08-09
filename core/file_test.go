package core

import (
	"testing"
)

func TestFileContext(t *testing.T) {
	fileCtx := NewFileCtx("/Users/ljackie/workspaces/go/permission-helper/14-Netty实战.pdf", &JavaDecode{})
	ctx := ReadCtx{
		MaxReadSize: 4 << 20, // 4M
	}
	ctx.read(fileCtx)
}
