package demo

import (
	"context"
	"io"
)

//go:generate gors -service NotDefine -path_to_lower

type NotDefine interface {
	String(context.Context, string) (string, error)
	Bytes(context.Context, []byte) ([]byte, error)
	Reader(context.Context, io.Reader) (io.Reader, error)
	NotDefine(context.Context, *NotDefineReq) (*NotDefineResp, error)
}

type NotDefineReq struct {
	ID int `uri:"id"`
}

type NotDefineResp struct {
	V int `json:"v,omitempty"`
}
