package core

import (
	"testing"
)

func TestFileContext(t *testing.T) {
	javaDecoder := &JavaDecoder{
		delegator: &JavaApiDecoder{},
	}
	fileCtx := NewFileCtx("/Users/ljackie/workspaces/go/permission-helper/RedisTokenManager.java", javaDecoder)
	ctx := ReadCtx{
		MaxReadSize: 4 << 20, // 4M
	}
	ctx.read(fileCtx)
}
