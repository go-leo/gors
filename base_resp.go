package gors

// BaseResp 基本响应
type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
