package core

import (
	"testing"
)

func TestFileContext(t *testing.T) {
	fileCtx := NewFileCtx("/Users/ljackie/workspaces/go/permission-helper/RedisTokenManager.java", &JavaDecode{})
	ctx := ReadCtx{
		MaxReadSize: 4 << 20, // 4M
	}
	ctx.read(fileCtx)
}
