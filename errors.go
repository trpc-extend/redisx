package redisx

import (
	redigo "github.com/gomodule/redigo/redis"

	"git.code.oa.com/trpc-go/trpc-go/errs"
)

// Error
var (
	// ErrNil redis nil
	ErrNil      = redigo.ErrNil
	ErrKeyEmpty = errs.New(-10001, "redigo: keys empty")
	ErrCASLock  = errs.New(-10002, "redigo: cas lock failed")

	// Pipeline
	ErrPipelineCmdsEmpty = errs.New(-20001, "redigo: pipeline cmds empty")
	ErrPipelineFlush     = errs.New(-20002, "redigo: pipeline")
	ErrPipelineRspEmpty  = errs.New(-20003, "redigo: pipeline rsp empty")
)
